package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"syscall"

	"bitbucket.org/multimfi/archivegen/elf"
)

const (
	TypeDirectory = "d"
	TypeRecursive = "R"
	TypeRegular   = "f"
	TypeSymlink   = "l"
	TypeCreate    = "c"
	TypeLinked    = "L"
	TypeLinkedAbs = "LA"
	TypeOmit      = "-"
)

type Map struct {
	// overlapping entries will be
	// replaced by subsequent entries.
	A []Entry

	// current set of masks.
	mm maskMap

	// lookup existance/index of entries.
	m map[string]int
}

func newMap() *Map {
	return &Map{
		m:  make(map[string]int),
		mm: make(maskMap),
		A:  make([]Entry, 0),
	}
}

func (m *Map) add(e entry, rootfs *string) error {
	switch e.Type() {
	case
		maskReplace,
		maskIgnore,
		maskIgnoreNeg,
		maskMode:
		return m.mm.set(e)
	case maskClear:
		return m.mm.del(e)
	}

	E, err := e.Entry()
	if err != nil {
		return err
	}

	if m.mm.apply(&E) {
		// ignored by mask
		return nil
	}

	// entry rootfs takes priority
	if r := e.Root(); r != nil {
		rootfs = r
	}

	switch E.Type {
	case
		TypeLinkedAbs,
		TypeLinked:
		return m.addElf(E, rootfs)
	case TypeRecursive:
		return m.addRecursive(
			E,
			e.isSet(idxUser),
			e.isSet(idxGroup),
		)
	}

	if i, exists := m.m[E.Dst]; exists {
		m.A[i] = E
		return nil
	}

	m.A = append(m.A, E)
	m.m[E.Dst] = len(m.A) - 1
	return nil
}

func (m *Map) Add(e Entry) {
	if m.mm.apply(&e) {
		return
	}

	if i, exists := m.m[e.Dst]; exists {
		m.A[i] = e
		return
	}

	m.A = append(m.A, e)
	m.m[e.Dst] = len(m.A) - 1
}

func (m *Map) addElf(e Entry, rootfs *string) error {
	var (
		r   []string
		err error
	)
	if rootfs != nil {
		r, err = elf.ResolveRoot(e.Src, *rootfs, e.Type == TypeLinkedAbs)
	} else {
		r, err = elf.Resolve(e.Src)
	}

	if err != nil {
		return err
	}

	m.Add(Entry{
		e.Src,
		e.Dst,
		e.User,
		e.Group,
		0755,
		TypeRegular,
		nil,
	})

	for _, v := range r {
		// '/usr/lib/lib.so'

		if v == e.Src {
			continue
		}

		var err error

		v, err = m.expand(v, rootfs)
		if err != nil {
			return err
		}

		dst := v
		if rootfs != nil {
			dst = strings.TrimPrefix(dst, *rootfs)
		}
		dst = strings.TrimPrefix(dst, "/")

		m.Add(Entry{
			v,
			dst,
			e.User,
			e.Group,
			0755,
			TypeRegular,
			nil,
		})
	}

	return nil
}

func (m *Map) Merge(t *Map) error {
	for _, v := range t.A {
		m.Add(v)
	}
	return nil
}

func (m *Map) addRecursive(e Entry, user, group bool) error {
	var uid, gid *int
	if user {
		uid = &e.User
	}
	if group {
		gid = &e.Group
	}
	return filepath.Walk(e.Src, mapW{m, e, uid, gid}.walkFunc)
}

type mapW struct {
	m   *Map
	e   Entry
	uid *int
	gid *int
}

func intPtr(i *int, d uint32) int {
	if i != nil {
		return *i
	}
	return int(d)
}

func (m mapW) walkFunc(file string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	// archive filepath
	af := strings.TrimPrefix(file, m.e.Src)
	if af == "" {
		return nil
	}

	var rf string
	if m.e.Dst != TypeOmit {
		rf = path.Join(m.e.Dst, af)
	} else {
		rf = path.Clean(af)
	}

	rf = strings.TrimPrefix(rf, "/")

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return fmt.Errorf("config: recursive: fileinfo not *Stat_t, %v", info.Sys())
	}

	if info.IsDir() {
		// TODO: mode masks
		m.m.Add(Entry{
			rf,
			rf,
			intPtr(m.uid, stat.Uid),
			intPtr(m.gid, stat.Gid),
			mode(info),
			TypeDirectory,
			nil,
		})
		return nil
	}

	if info.Mode().IsRegular() {
		m.m.Add(Entry{
			file,
			rf,
			intPtr(m.uid, stat.Uid),
			intPtr(m.gid, stat.Gid),
			mode(info),
			TypeRegular,
			nil,
		})
		return nil
	}

	if info.Mode()&os.ModeSymlink != 0 {
		f, err := os.Readlink(file)
		if err != nil {
			return err
		}

		m.m.Add(Entry{
			f,
			rf,
			intPtr(m.uid, stat.Uid),
			intPtr(m.gid, stat.Gid),
			0777,
			TypeSymlink,
			nil,
		})
		return nil
	}

	return fmt.Errorf("config: recursive: unknown file: %s", file)
}
