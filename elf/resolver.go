package elf

import (
	"debug/elf"
	"path"
	"syscall"
)

// for testing
var (
	fAccess = access
	eOpen   = elfOpen
)

var libpaths = []string{
	"/usr/lib",
	"/usr/lib64",
	"/lib",
	"/lib64",
}

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
	return ret, errorNotFound{f}
}

func resolv(file string, ret map[string]string) error {
	_, ff := path.Split(file)
	if _, exists := ret[ff]; exists {
		return nil
	}
	ret[ff] = file

	var (
		n  []string
		r1 []string
		r2 []string
		lp []string
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

	lp = append(lp, r1...)
	lp = append(lp, r2...)

	for _, v := range n {
		s, err := search(v, append(lp, libpaths...))
		if err != nil {
			return err
		}
		if err := resolv(s, ret); err != nil {
			return err
		}
		ret[v] = s
	}
	return nil
}

// Resolve resolves libraries needed by an ELF file.
func Resolve(file string) (map[string]string, error) {
	ret := make(map[string]string)
	if err := resolv(file, ret); err != nil {
		return nil, err
	}
	return ret, nil
}
