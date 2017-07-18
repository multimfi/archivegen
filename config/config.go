package config

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

var (
	errInvalidEntry = errors.New("config: invalid entry")
	errIgnoredEntry = errors.New("config: ignored entry")
	errNoArguments  = errors.New("config: no arguments")
)

type ctx struct {
	l int
	r rune
	n int
	f int
}

func split(r rune) bool {
	return r == ' ' || r == '\t'
}

func (c *ctx) split(r rune) bool {
	if c.n == 0 {
		c.r = r
	}
	if c.r != 'c' {
		return split(r)
	}
	if c.n >= c.l {
		c.n = 0
		c.f = 0
	}

	c.n++

	if c.f >= idxData-1 {
		return false
	}

	if split(r) {
		c.f++
		return true
	}

	return false
}

func nc(r string) func(rune) bool {
	return (&ctx{
		l: len(r),
	}).split
}

// TODO: error handling
func FromReader(r io.Reader) *Map {
	s := bufio.NewScanner(r)
	m := newMap()

	var n int
	for s.Scan() {
		n++
		if err := s.Err(); err != nil {
			log.Printf("error: %q, line %d", err, n)
			return nil
		}
		d := s.Text()
		if len(d) < 1 {
			continue
		}
		if d[0] == '\n' || d[0] == '#' {
			continue
		}

		f := strings.FieldsFunc(d, nc(d))
		if len(f) < 2 {
			log.Printf("error: %s, line %d", errNoArguments, n)
			continue
		}
		if err := m.add(f); err != nil {
			log.Printf("error: %s, line %d", err, n)
			return nil
		}
	}

	return m
}

func FromFiles(files ...string) (*Map, error) {
	cfg := newMap()
	for _, v := range files {
		f, err := os.Open(path.Clean(v))
		if err != nil {
			return nil, err
		}

		c := FromReader(f)
		if c == nil {
			return nil, fmt.Errorf("error")
		}

		if err := cfg.Merge(c); err != nil {
			return nil, err
		}

		if err := f.Close(); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}
