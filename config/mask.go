package config

import (
	"regexp"
	"strconv"
)

const (
	maskMode      = "mm"
	maskClear     = "mc"
	maskIgnore    = "mi"
	maskIgnoreNeg = "mI"
	maskReplace   = "mr"
)

const (
	idxMaskID     = 1
	idxMaskRegexp = 2
	idxMaskDst    = 3
	idxMaskMode   = 3
	idxMaskUid    = 4
	idxMaskGid    = 5
)

type maskFunc func(*Entry) bool

type maskMap map[int]maskFunc

func (m maskMap) apply(e *Entry) bool {
	for i := 0; i < len(m); i++ {
		f := m[i]
		if f == nil {
			continue
		}
		if f(e) {
			return true
		}
	}
	return false
}

func (m maskMap) set(e entry) error {
	i, err := maskID(e)
	if err != nil {
		return err
	}
	m[i], err = maskFromEntry(e)
	if err != nil {
		return err
	}
	return nil
}

func (m maskMap) clear() {
	for k, _ := range m {
		delete(m, k)
	}
}

func (m maskMap) del(e entry) error {
	if len(e) < idxMaskID+1 {
		m.clear()
		return nil
	}

	i, err := maskID(e)
	if err != nil {
		return err
	}
	m[i] = nil

	return nil
}

func maskID(e entry) (int, error) {
	return strconv.Atoi(e[idxMaskID])
}

func maskFromEntry(e entry) (maskFunc, error) {
	if len(e) < 2 {
		return nil, errInvalidEntry
	}
	switch e[idxType] {
	case maskReplace:
		return regexReplaceMask(e)
	case maskMode:
		return regexModeMask(e)
	case maskIgnore:
		return regexIgnoreMask(e, false)
	case maskIgnoreNeg:
		return regexIgnoreMask(e, true)
	}
	return nil, errInvalidEntry
}

func regexReplaceMask(e entry) (maskFunc, error) {
	if len(e) < idxMaskDst {
		return nil, errInvalidEntry
	}

	r, err := regexp.Compile(e[idxMaskRegexp])
	if err != nil {
		return nil, err
	}

	return func(E *Entry) bool {
		switch E.Type {
		case
			TypeLinked,
			TypeRecursive:
			return false
		}
		E.Dst = r.ReplaceAllString(E.Dst, e[idxMaskDst])
		return false
	}, nil
}

func regexIgnoreMask(e entry, neg bool) (maskFunc, error) {
	if len(e) < idxMaskDst {
		return nil, errInvalidEntry
	}

	r, err := regexp.Compile(e[idxMaskRegexp])
	if err != nil {
		return nil, err
	}

	return func(E *Entry) bool {
		switch E.Type {
		case
			TypeLinked,
			TypeRecursive:
			return false
		}
		if neg {
			return !r.MatchString(E.Dst)
		}
		return r.MatchString(E.Dst)
	}, nil
}

func regexModeMask(e entry) (maskFunc, error) {
	if len(e) < idxMaskDst {
		return nil, errInvalidEntry
	}

	r, err := regexp.Compile(e[idxMaskRegexp])
	if err != nil {
		return nil, err
	}

	var (
		mode *int
		uid  *int
		gid  *int
	)

	if mode, err = e.pMode(); err != nil {
		return nil, err
	}

	if uid, err = e.pUser(); err != nil {
		return nil, err
	}

	if gid, err = e.pGroup(); err != nil {
		return nil, err
	}

	return func(E *Entry) bool {
		switch E.Type {
		case
			TypeLinked,
			TypeRecursive:
			return false
		}
		if !r.MatchString(E.Dst) {
			return false
		}
		if mode != nil {
			E.Mode = *mode
		}
		if gid != nil {
			E.Group = *gid
		}
		if uid != nil {
			E.User = *uid
		}
		return false
	}, nil
}
