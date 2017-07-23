package elf

import (
	"debug/elf"
	"errors"
	"testing"
)

type ef [][]string

func get(s [][]string, n int) []string {
	var r []string
	if len(s)-1 >= n {
		r = s[n]
	}
	return r
}

func (e ef) DynString(tag elf.DynTag) ([]string, error) {
	switch tag {
	case elf.DT_NEEDED:
		return get(e, 0), nil
	case elf.DT_RUNPATH:
		return get(e, 1), nil
	case elf.DT_RPATH:
		return get(e, 2), nil
	}
	return nil, errors.New("no such tag: " + tag.String())
}

func (e ef) Close() error {
	return nil
}

func mapAccess(d map[string]ef) func(f string) error {
	return func(f string) error {
		if _, exists := d[f]; !exists {
			return errorNotFound{f}
		}
		return nil
	}
}

func mapOpen(d map[string]ef) func(f string) (elfFile, error) {
	return func(f string) (elfFile, error) {
		r, exists := d[f]
		if !exists {
			return nil, errorNotFound{f}
		}
		return r, nil
	}
}

func testResolve(t *testing.T, f string, re map[string]string, data map[string]ef) {
	eOpen = mapOpen(data)
	fAccess = mapAccess(data)

	r, err := Resolve(f)
	if err != nil {
		t.Fatal(err)
	}

	if w, r := len(re), len(r); w != r {
		t.Fatalf("len w(%d) != r(%d)", w, r)
	}

	for k, v := range re {
		var (
			rV     string
			exists bool
		)
		if rV, exists = r[k]; !exists {
			t.Fatalf("key %q does not exist", k)
		}
		if rV != v {
			t.Fatalf("value does not equal %q != %q", rV, v)
		}
	}
}

func TestResolveMPD(t *testing.T) {
	testResolve(t, "/usr/bin/mpd", mpdresolved, mpdfiles)
}

func TestResolveQemu(t *testing.T) {
	testResolve(t, "/usr/bin/qemu-system-x86_64", qemuresolved, qemufiles)
}
