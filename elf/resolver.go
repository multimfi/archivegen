package elf

import (
	"debug/elf"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const ldConfigDir = "/etc/ld.so.conf.d"

// overridden in testing
var open = eopen

type errorNotFound struct {
	e string
}

func (e errorNotFound) Error() string {
	return "resolver: not found: " + e.e
}

func split(a []string) []string {
	var r []string
	for _, v := range a {
		r = append(r, strings.Split(v, ":")...)
	}
	return r
}

type pathset map[string]int

func (p pathset) copy() pathset {
	r := make(pathset, len(p))
	for k, v := range p {
		r[k] = v
	}
	return r
}

func tokenExpander(origin string) *strings.Replacer {
	return strings.NewReplacer(
		"$ORIGIN", origin,
		"${ORIGIN}", origin,
		"$LIB", "lib64",
		"${LIB}", "lib64",
		"$PLATFORM", "x86_64",
		"${PLATFORM}", "x86_64",
	)
}

func (p pathset) add(origin string, s ...string) pathset {
	var (
		sr *strings.Replacer
		c  bool
		r  = p
	)

	for _, v := range s {
		if len(v) < 1 {
			continue
		}
		switch v[0] {
		case '/', '$':
			break
		default:
			continue
		}

		if sr == nil {
			sr = tokenExpander(origin)
		}
		v = sr.Replace(v)

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

func (p pathset) list() []string {
	i := len(p)
	r := make([]string, i)

	for k, v := range p {
		r[v] = k
	}

	return r
}

func rootprefix(file string, rootfs *string, abs bool) string {
	if abs {
		return file
	}
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

type context struct {
	err    set
	ldconf pathset
	class  elf.Class
	root   *string
}

func (c *context) search1(file string, ret set, from []string) (string, elfFile, error) {
	var r string

	for _, v := range from {
		if file[0] != '/' {
			// relative path
			r = path.Join(
				rootprefix(v, c.root, false),
				file,
			)
		} else {
			// absolute path
			r = file
		}

		_, exists := ret[r]
		if exists {
			return r, nil, nil
		}

		f, err := open(r)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			switch err.(type) {
			case *elf.FormatError:
				continue
			}
			return "", nil, err
		}

		if e, ok := f.(*elf.File); ok {
			if e.Class != c.class {
				if err := f.Close(); err != nil {
					return "", nil, err
				}
				continue
			}
		}

		return r, f, nil
	}

	return r, nil, errorNotFound{r}
}

func (c *context) search(file string, ret set, path ...[]string) (string, elfFile, error) {
	var (
		f   elfFile
		r   string
		err error
	)

	for _, v := range path {
		if r, f, err = c.search1(file, ret, v); err == nil {
			return r, f, nil
		}

		switch err.(type) {
		case errorNotFound:
			continue
		default:
			return r, nil, err
		}
	}

	if r, f, err = c.search1(file, ret, c.ldconf.list()); err != nil {
		switch err.(type) {
		case errorNotFound:
			return c.search1(file, ret, defaultLibs)
		default:
			return r, nil, err
		}
	}

	return r, f, err
}

func (c *context) resolv(file string, f elfFile, rpath pathset, runpath []string, ret set) error {
	if ret.add(file) {
		return nil
	}

	needed, err := f.DynString(elf.DT_NEEDED)
	if err != nil {
		return err
	}

	oldrunpath := runpath
	runpath, err = f.DynString(elf.DT_RUNPATH)
	if err != nil {
		return err
	}

	rpathE, err := f.DynString(elf.DT_RPATH)
	if err != nil {
		return err
	}

	// opened in resolve/search
	if err := f.Close(); err != nil {
		return err
	}

	var rd string
	if c.root != nil {
		rd = path.Dir(strings.TrimPrefix(file, *c.root))
	} else {
		rd = path.Dir(file)
	}
	rpath = rpath.add(
		rd,
		split(rpathE)...,
	)

	if len(runpath) > 0 {
		x := tokenExpander(rd)
		for k, _ := range runpath {
			runpath[k] = x.Replace(runpath[k])
		}
	}

	runpath = split(runpath)

	for _, v := range needed {
		// glibc libc.so is not an elf and
		// in musl it is the interpreter
		if v == "libc.so" {
			break
		}

		s, fd, err := c.search(
			v,
			ret,
			oldrunpath,
			runpath,
			rpath.list(),
		)
		if err != nil {
			switch err.(type) {
			case errorNotFound:
				c.err.add(v)
				continue
			default:
				return err
			}
		}

		if fd == nil {
			continue
		}

		delete(c.err, v)

		if err := c.resolv(s, fd, rpath, runpath, ret); err != nil {
			return err
		}

		ret.add(s)
	}

	return nil
}

func ldList(dir string) ([]string, error) {
	var r []string

	d, err := ioutil.ReadDir(dir)
	if err != nil {
		// non-fatal
		return r, nil
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
		r = append(r, strings.Split(data, "\n")...)
	}

	return r, nil
}

var defaultLibs = []string{
	"/lib64",
	"/usr/lib64",
	"/lib",
	"/usr/lib",
	// freebsd
	// "/usr/lib/compat",
	// "/usr/local/lib",
}

func resolve(file string, rootfs *string, abs bool) ([]string, error) {
	var ctx context
	ctx.err = make(set)

	dir := rootprefix(path.Dir(file), rootfs, abs)

	if config, err := ldList(
		rootprefix(ldConfigDir, rootfs, false),
	); err != nil {
		return nil, err
	} else {
		ctx.ldconf = ctx.ldconf.add(dir, config...)
	}

	f, err := open(rootprefix(file, rootfs, abs))
	if err != nil {
		return nil, err
	}

	if e, ok := f.(*elf.File); !ok {
		ctx.class = elf.ELFCLASS64
	} else {
		ctx.class = e.Class
	}
	ctx.root = rootfs

	ret := make(set)
	if i, err := readinterp(f); err == nil {
		ret.add(rootprefix(i, rootfs, false))
	} /* else {
		log.Println("resolve:", err)
	} */

	if err := ctx.resolv(
		rootprefix(file, rootfs, abs),
		f,
		make(pathset),
		nil,
		ret,
	); err != nil {
		return nil, err
	}

	if len(ctx.err) > 0 {
		return nil, errorNotFound{
			strings.Join(ctx.err.list(), ", "),
		}
	}

	return ret.list(), nil
}

// Resolve resolves libraries needed by an ELF file.
func Resolve(file string) ([]string, error) {
	return resolve(file, nil, true)
}

// ResolveRoot searches libraries from rootfs. If abs, file will not prefixed with rootfs.
func ResolveRoot(file, rootfs string, abs bool) ([]string, error) {
	return resolve(file, &rootfs, abs)
}
