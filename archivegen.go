package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"text/tabwriter"

	"github.com/multimfi/archivegen/archive"
	"github.com/multimfi/archivegen/archive/cpio"
	"github.com/multimfi/archivegen/archive/tar"
	"github.com/multimfi/archivegen/config"
	"github.com/multimfi/archivegen/tree"
)

var (
	buildversion = "v0"
	rootprefix   string
)

func open(file *string) *os.File {
	if file == nil {
		return os.Stdout
	}
	r, err := os.OpenFile(
		*file,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	return r
}

func getTree(rootfs string, vars, files []string, stdin bool) *tree.Node {
	var (
		r   *config.Map
		err error
	)

	if stdin {
		r = config.FromReaderRoot(rootfs, vars, os.Stdin)
	} else {
		r, err = config.FromFilesRoot(rootfs, vars, files...)
	}
	if err != nil {
		log.Fatal(err)
	}
	if r == nil {
		log.Fatal("nil map")
	}

	return tree.Render(r)

}

func printTree(t *tree.Node) {
	tw := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)
	t.Print(rootprefix, tw)
	tw.Flush()
}

func getArchive(fmt string, dst io.Writer, timestamp bool) archive.Writer {
	switch fmt {
	case "cpio":
		return cpio.NewWriter(dst, timestamp)
	case "tar":
		return tar.NewWriter(dst, timestamp)
	}

	log.Fatal("unknown archive format", fmt)
	return nil
}

func stdinPipe() bool {
	f, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}
	return (f.Mode() & os.ModeCharDevice) == 0
}

type varValue []string

func (v *varValue) String() string {
	return ""
}

var errInvalidFlag = errors.New("invalid flag")

func (v *varValue) Set(val string) error {
	x := strings.SplitN(val, "=", 2)
	if len(x) < 2 {
		return errInvalidFlag
	}
	*v = append(*v, x...)
	return nil
}

var (
	flagOut           = flag.String("out", "out.archive", "output file")
	flagFormat        = flag.String("fmt", "tar", "file format, cpio/tar")
	flagRootfs        = flag.String("rootfs", "", "ELF rootfs")
	flagPrint         = flag.Bool("print", false, "print resolved tree in archivegen format")
	flagTimestamp     = flag.Bool("timestamp", false, "preserve file timestamps")
	flagStdout        = flag.Bool("stdout", false, "output to stdout")
	flagVersion       = flag.Bool("version", false, "version")
	flagArchiveFormat = flag.Bool("format", false, "print archive format")
)

const defaultBufSize = 1 << 24

func main() {
	log.SetFlags(log.Lshortfile)

	var varX varValue
	flag.Var(&varX, "X", "variable\n"+
		"e.g. '-X foo=bar -X a=b'",
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s %s\n", "archivegen", "[OPTIONS...] [FILES...]")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *flagArchiveFormat {
		fmt.Fprintln(os.Stderr, helpFormat)
		return
	}

	if *flagStdout {
		flagOut = nil
	}
	if *flagVersion {
		fmt.Printf("build: %s\nruntime: %s\n", buildversion, runtime.Version())
		return
	}

	p := stdinPipe()
	if flag.NArg() < 1 && !p {
		log.Fatal("not enough arguments")
	}

	root := getTree(*flagRootfs, []string(varX), flag.Args(), p && flag.NArg() < 1)

	if *flagPrint {
		printTree(root)
		os.Exit(0)
	}

	out := open(flagOut)
	buf := bufio.NewWriterSize(out, defaultBufSize)
	in := getArchive(*flagFormat, buf, *flagTimestamp)

	if err := root.Write(rootprefix, in); err != nil {
		log.Fatal("write:", err)
	}
	if err := in.Close(); err != nil {
		log.Fatal("archive: close:", err)
	}
	if err := buf.Flush(); err != nil {
		log.Fatal("buffer: flush:", err)
	}
	if err := out.Close(); err != nil {
		log.Fatal("output: close:", err)
	}
}
