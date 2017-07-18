package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
)

const (
	modeSticky = 1 << (iota + 9)
	modeSetgid
	modeSetuid
)

const (
	idxType = iota
	idxSrc
	idxDst
	idxMode
	idxUser
	idxGroup
	idxData
)

type entry []string

// TODO: uint32 uid, gid <-> syscall.Stat_t
type Entry struct {
	Src, Dst    string
	User, Group int
	Mode        int
	Type        string
	Data        []byte
}

func (e entry) Type() string {
	return e[idxType]
}

func (e entry) Src() (string, error) {
	switch e.Type() {
	case TypeRegular:
		if len(e) < 2 {
			break
		}
		return e[1], nil

	case TypeSymlink:
		if len(e) < 3 {
			break
		}
		return e[1], nil

	case
		TypeDirectory,
		TypeRecursive,
		TypeCreate,
		TypeLinked:
		if len(e) < 2 {
			break
		}
		return e[1], nil
	}

	log.Printf("error: %v", e)
	return "", errInvalidEntry
}

func clean(file string) string {
	p := path.Clean(file)
	if p[0] != '/' {
		return p
	}
	return p[1:]
}

func (e entry) Dst() (string, error) {
	switch e.Type() {
	case TypeDirectory:
		// invalid entry
		if len(e) < 2 {
			break
		}
		// directory dst is the src
		return clean(e[1]), nil

	case TypeRegular:
		// invalid entry
		if len(e) < 2 {
			break
		}

		// omitted dst
		if len(e) < 3 {
			return clean(e[1]), nil
		}

		// dst set
		if e[2] != TypeOmit {
			return clean(e[2]), nil
		}

		// explicitly omitted dst
		return clean(e[1]), nil

	case TypeLinked:
		if len(e) > 2 {
			return clean(e[2]), nil
		}
		return clean(e[1]), nil
	case
		TypeSymlink,
		TypeRecursive:
		if len(e) < 3 {
			break
		}
		return clean(e[2]), nil

	case TypeCreate:
		if len(e) < 2 {
			break
		}
		return clean(e[1]), nil

	}

	log.Printf("error: %v", e)
	return "", errInvalidEntry
}

func (e entry) Mode() (int, error) {
	var (
		i int    = idxMode
		t string = e.Type()
	)

	// src field omitted
	if t == TypeDirectory || t == TypeCreate {
		i -= 1
	}

	if len(e) <= i || e[i] == TypeOmit {
		switch t {
		case TypeDirectory:
			return 0755, nil
		case TypeSymlink:
			return 0777, nil
		default:
			return 0644, nil
		}
	}

	m, err := strconv.ParseInt(e[i], 8, 0)
	if err != nil {
		return 0, err
	}

	if m&modeSticky != 0 {
		m |= int64(os.ModeSticky)
	}
	if m&modeSetgid != 0 {
		m |= int64(os.ModeSetgid)
	}
	if m&modeSetuid != 0 {
		m |= int64(os.ModeSetuid)
	}

	return int(m), nil
}

func (e entry) parseIndex(i int) (int, error) {
	t := e.Type()
	if t == TypeDirectory || t == TypeCreate {
		i -= 1
	}
	if len(e) <= i || e[i] == TypeOmit {
		return 0, nil
	}
	r, err := strconv.ParseInt(e[i], 10, 0)
	return int(r), err
}

func (e entry) User() (int, error) {
	r, err := e.parseIndex(idxUser)
	return int(r), err
}

func (e entry) Group() (int, error) {
	r, err := e.parseIndex(idxGroup)
	return int(r), err
}

func (e entry) Data() []byte {
	if e.Type() != TypeCreate {
		return nil
	}
	return []byte(e[idxData-1])
}

func (e entry) Entry() (Entry, error) {
	var (
		r   Entry
		err error
	)

	// TODO: error handling
	r.Dst, err = e.Dst()
	if err != nil {
		return r, err
	}

	r.Src, err = e.Src()
	if err != nil {
		return r, err
	}

	r.Mode, err = e.Mode()
	if err != nil {
		return r, err
	}

	r.User, err = e.User()
	if err != nil {
		return r, err
	}

	r.Group, err = e.Group()
	if err != nil {
		return r, err
	}

	r.Type = e.Type()
	r.Data = e.Data()

	return r, nil
}

func (e Entry) Format() string {
	switch e.Type {
	case TypeDirectory:
		return fmt.Sprintf("%s\t%s\t\t%04o\t%d\t%d",
			e.Type,
			e.Dst,
			e.Mode,
			e.User,
			e.Group,
		)

	case TypeCreate:
		return fmt.Sprintf("%s\t%s\t\t%04o\t%d\t%d\t%s",
			e.Type,
			e.Dst,
			e.Mode,
			e.User,
			e.Group,
			e.Data,
		)
	}

	return fmt.Sprintf("%s\t%s\t%s\t%04o\t%d\t%d",
		e.Type,
		e.Src,
		e.Dst,
		e.Mode,
		e.User,
		e.Group,
	)
}
