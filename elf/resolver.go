package elf

import (
	"debug/elf"
	"io/ioutil"
	"path"
	"strings"
	"syscall"
)

const ldConfigDir = "/etc/ld.so.conf.d"

// for testing
var (
	fAccess = access
	eOpen   = elfOpen
)

type errorNotFound struct {
	e string
}

func (e errorNotFound) Error() string {
	return "resolver: not found: " + e.e
}

func access(f string) error {
	return syscall.Access(f, syscall.F_OK)
}

func search(f string, d []string, rootfs *string) (string, error) {
	var ret string
	for _, v := range d {
		ret = path.Join(rPath(v, rootfs), f)
		if err := fAccess(ret); err != nil {
			continue
		}
		return ret, nil
	}
	return ret, errorNotFound{ret}
}

func split(a []string) []string {
	var r []string
	for _, v := range a {
		r = append(r, strings.Split(v, ":")...)
	}
	return r
}

type libOrder map[string]int

func (o libOrder) copy() libOrder {
	r := make(libOrder, len(o))
	for k, v := range o {
		r[k] = v
	}
	return r
}

func (o libOrder) add(s ...string) libOrder {
	var (
		c bool
		r = o
	)

	for _, v := range s {
		if len(v) < 1 {
			continue
		}
		if v[0] != '/' {
			continue
		}
		if _, exists := r[v]; exists {
			continue
		}
		if !c {
			r = r.copy()
			c = true
		}
		r[v] = len(r)
	}

	return r
}

func (o libOrder) list() []string {
	r := make([]string, len(o))
	for k, v := range o {
		r[v] = k
	}
	return r
}

func rPath(file string, rootfs *string) string {
	if rootfs == nil {
		return file
	}
	if *rootfs == "" {
		return file
	}
	return path.Join(*rootfs, file)
}

type set map[string]struct{}

func (s set) add(key string) bool {
	if _, exists := s[key]; exists {
		return exists
	}

	s[key] = struct{}{}

	return false
}

func (s set) list() []string {
	var (
		r = make([]string, len(s))
		i int
	)
	for k, _ := range s {
		r[i] = k
		i++
	}
	return r
}

func resolv(file string, lp libOrder, rootfs *string, ret set) error {
	if ret.add(file) {
		return nil
	}

	var (
		n  []string
		r1 []string
		r2 []string
	)

	f, err := eOpen(file)
	if err != nil {
		return err
	}

	n, err = f.DynString(elf.DT_NEEDED)
	if err != nil {
		return err
	}
	r1, err = f.DynString(elf.DT_RUNPATH)
	if err != nil {
		return err
	}
	r2, err = f.DynString(elf.DT_RPATH)
	if err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	lp = lp.add(split(r1)...)
	lp = lp.add(split(r2)...)

	for _, v := range n {
		s, err := search(v, lp.list(), rootfs)
		if err != nil {
			return err
		}
		if err := resolv(s, lp, rootfs, ret); err != nil {
			return err
		}
		ret.add(s)
	}
	return nil
}

var defaultLibs = libOrder{
	"/usr/lib":   0,
	"/usr/lib64": 1,
	"/lib":       2,
	"/lib64":     3,
}

func ldList(dir string) ([]string, error) {
	var ret []string

	d, err := ioutil.ReadDir(dir)
	if err != nil {
		// non-fatal
		return ret, nil
	}

	for _, v := range d {
		if v.IsDir() {
			continue
		}
		f, err := ioutil.ReadFile(path.Join(dir, v.Name()))
		if err != nil {
			return nil, err
		}
		data := strings.TrimRight(string(f), "\n")
		ret = append(ret, strings.Split(data, "\n")...)
	}

	return ret, nil
}

func resolve(file string, rootfs *string) ([]string, error) {
	ld, err := ldList(rPath(ldConfigDir, rootfs))
	if err != nil {
		return nil, err
	}

	ret := make(set)
	if err := resolv(
		rPath(file, rootfs),
		defaultLibs.add(ld...),
		rootfs,
		ret,
	); err != nil {
		return nil, err
	}

	return ret.list(), nil
}

// Resolve resolves libraries needed by an ELF file.
func Resolve(file string) ([]string, error) {
	return resolve(file, nil)
}

// ResolveWithRoot searches libraries from rootfs.
func ResolveWithRoot(file, rootfs string) ([]string, error) {
	return resolve(file, &rootfs)
}
