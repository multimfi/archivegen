package elf

import "debug/elf"

type elfFile interface {
	DynString(tag elf.DynTag) ([]string, error)
	Close() error
}

func elfOpen(f string) (r elfFile, err error) {
	r, err = elf.Open(f)
	return
}
