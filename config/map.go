package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/multimfi/archivegen/elf"
)

const (
	TypeOmit         = "-"
	TypeDirectory    = "d"
	TypeRecursive    = "R"
	TypeRecursiveRel = "Rr"
	TypeRegular      = "f"
	TypeRegularRel   = "fr"
	TypeGlob         = "g"
	TypeGlobRel      = "gr"
	TypeSymlink      = "l"
	TypeCreate       = "c"
	TypeLinked       = "L"
	TypeLinkedAbs    = "LA"
	TypeVariable     = "$"
)

type variable struct {
	value string
	flag  bool
}

type variableMap struct {
	m map[string]variable
	r *strings.Replacer
}

func newVariableMap(vars []string) *variableMap {
	if len(vars)&0x1 != 0 {
		panic("invalid vars")
	}

	ret := &variableMap{
		m: make(map[string]variable),
		r: strings.NewReplacer(),
	}

	if vars == nil {
		return ret
	}

	r := make([]string, 0)
	for x := 0; x < len(vars); x += 2 {
		ret.m[vars[x]] = variable{
			flag:  true,
			value: vars[x+1],
		}

		r = append(r, TypeVariable+vars[x], vars[x+1])
	}

	ret.r = strings.NewReplacer(r...)

	return ret
}

func (m *variableMap) add(e entry) error {
	if len(e) < 2 {
		return errInvalidEntry
	}

	// variables with flag variable are not mutable
	if v, ok := m.m[e[idxSrc]]; v.flag && ok {
		return nil
	}

	if len(e) > 2 {
		m.m[e[idxSrc]] = variable{value: e[idxDst]}
	} else {
		// without dst assume empty string
		// should this be TypeOmit?
		m.m[e[idxSrc]] = variable{}
	}

	r := make([]string, 0, len(m.m))
	for k, v := range m.m {
		r = append(r, TypeVariable+k, v.value)
	}
	m.r = strings.NewReplacer(r...)

	return nil
}

type Map struct {
	// overlapping entries will be
	// replaced by subsequent entries.
	A []Entry

	// current set of masks.
	mm maskMap

	// lookup existance/index of entries.
	m map[string]int

	// variable map
	v *variableMap
}

func newMap(vars []string) *Map {
	return &Map{
		m:  make(map[string]int),
		mm: make(maskMap),
		A:  make([]Entry, 0),
		v:  newVariableMap(vars),
	}
}

func (m *Map) add(e entry, rootfs *string) error {
	for k, _ := range e {
		e[k] = m.v.r.Replace(e[k])
	}

	switch e.Type() {
	case
		maskReplace,
		maskIgnore,
		maskIgnoreNeg,
		maskMode:
		return m.mm.set(e)
	case maskClear:
		return m.mm.del(e)
	case TypeVariable:
		return m.v.add(e)
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
	case TypeRegularRel:
		E.Type = TypeRegular
		fallthrough
	case
		TypeRecursiveRel,
		TypeGlobRel:
		E.Src = rootprefix(E.Src, rootfs)
	}

	switch E.Type {
	case
		TypeLinkedAbs,
		TypeLinked:
		return m.addElf(E, rootfs)
	case
		TypeRecursiveRel,
		TypeRecursive:
		return m.addRecursive(
			E,
			e.isSet(idxUser),
			e.isSet(idxGroup),
			rootfs,
		)
	case
		TypeGlob,
		TypeGlobRel:
		return m.addGlob(
			E,
			e.isSet(idxUser),
			e.isSet(idxGroup),
			rootfs,
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

func rootprefix(file string, rootfs *string) string {
	if rootfs == nil {
		return file
	}
	if *rootfs == "" {
		return file
	}
	return path.Join(*rootfs, file)
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

	var src string

	if e.Type != TypeLinkedAbs {
		src = rootprefix(e.Src, rootfs)
	} else {
		src = e.Src
	}

	m.Add(Entry{
		src,
		e.Dst,
		e.User,
		e.Group,
		0755,
		TypeRegular,
		nil, "",
	})

	for _, v := range r {
		// '/usr/lib/lib.so'

		if v == src {
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
			nil, "",
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

func (m *Map) addRecursive(e Entry, user, group bool, rootfs *string) error {
	var uid, gid *int
	if user {
		uid = &e.User
	}
	if group {
		gid = &e.Group
	}
	return filepath.Walk(e.Src, mapW{m, e, uid, gid, rootfs}.walkFunc)
}

type mapW struct {
	m      *Map
	e      Entry
	uid    *int
	gid    *int
	rootfs *string
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

	if m.rootfs != nil {
		rf = strings.TrimPrefix(rf, *m.rootfs)
	}
	rf = strings.TrimPrefix(rf, "/")

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return fmt.Errorf("config: recursive: fileinfo not *Stat_t, %#v)", info.Sys())
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
			nil, "",
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
			nil, "",
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
			nil, "",
		})
		return nil
	}

	return fmt.Errorf("config: recursive: unknown file: %s", file)
}

func (m *Map) addGlob(e Entry, user, group bool, rootfs *string) error {
	r, err := filepath.Glob(e.Src)
	if err != nil {
		return err
	}

	x := mapW{m: m, e: e, rootfs: rootfs}
	if user {
		x.uid = &e.User
	}
	if group {
		x.gid = &e.Group
	}

	for _, v := range r {
		s, err := os.Lstat(v)
		if err := x.walkFunc(v, s, err); err != nil {
			return err
		}
	}

	return nil
}
