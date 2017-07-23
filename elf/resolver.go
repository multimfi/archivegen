package elf

import (
	"debug/elf"
	"path"
	"strings"
	"syscall"
)

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

func search(f string, d []string) (string, error) {
	var ret string
	for _, v := range d {
		ret = path.Join(v, f)
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

type order map[string]int

func (o order) copy() order {
	r := make(order, len(o))
	for k, v := range o {
		r[k] = v
	}
	return r
}

func (o order) add(s ...string) order {
	var (
		c bool
		r order = o
	)

	for _, v := range s {
		if _, exists := r[v]; exists {
			continue
		}
		if !c {
			r = r.copy()
			c = true
		}
		r[v] = len(o)
	}

	return r
}

func (o order) list() []string {
	r := make([]string, len(o))
	for k, v := range o {
		r[v] = k
	}
	return r
}

func resolv(file string, ret map[string]string, lp order) error {
	_, ff := path.Split(file)
	if _, exists := ret[ff]; exists {
		return nil
	}
	ret[ff] = file

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
		s, err := search(v, lp.list())
		if err != nil {
			return err
		}
		if err := resolv(s, ret, lp); err != nil {
			return err
		}
		ret[v] = s
	}
	return nil
}

var libpaths = order{
	"/usr/lib":   0,
	"/usr/lib64": 1,
	"/lib":       2,
	"/lib64":     3,
}

// Resolve resolves libraries needed by an ELF file.
func Resolve(file string) (map[string]string, error) {
	ret := make(map[string]string)
	if err := resolv(file, ret, libpaths); err != nil {
		return nil, err
	}
	return ret, nil
}
