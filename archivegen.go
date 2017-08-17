package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"text/tabwriter"

	"bitbucket.org/multimfi/archivegen/archive"
	"bitbucket.org/multimfi/archivegen/archive/cpio"
	"bitbucket.org/multimfi/archivegen/archive/tar"
	"bitbucket.org/multimfi/archivegen/config"
	"bitbucket.org/multimfi/archivegen/tree"
)

var (
	buildversion string
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

func getTree(rootfs string, files []string, stdin bool) *tree.Node {
	var (
		r   *config.Map
		err error
	)

	if stdin {
		r = config.FromReaderRoot(rootfs, os.Stdin)
	} else {
		r, err = config.FromFilesRoot(rootfs, files...)
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

func getArchive(fmt string, dst io.Writer) archive.Writer {
	switch fmt {
	case "cpio":
		return cpio.NewWriter(dst)
	case "tar":
		return tar.NewWriter(dst)
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

var (
	flagOut     = flag.String("out", "out.archive", "output file")
	flagFormat  = flag.String("fmt", "tar", "file format, cpio/tar")
	flagRootfs  = flag.String("rootfs", "", "ELF rootfs")
	flagPrint   = flag.Bool("print", false, "print resolved tree in archivegen format")
	flagStdout  = flag.Bool("stdout", false, "output to stdout")
	flagVersion = flag.Bool("version", false, "version")
)

func main() {
	log.SetFlags(log.Lshortfile)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s %s\n\n", "archivegen", "[OPTIONS...] [FILES...]")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, helpFormat)
	}
	flag.Parse()

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

	root := getTree(*flagRootfs, flag.Args(), p)

	if *flagPrint {
		printTree(root)
		os.Exit(0)
	}

	out := open(flagOut)
	in := getArchive(*flagFormat, out)

	if err := root.Write(rootprefix, in); err != nil {
		log.Fatal("error: write:", err)
	}
	if err := in.Close(); err != nil {
		log.Fatal("error: archive: close:", err)
	}
	if err := out.Close(); err != nil {
		log.Fatal("error: output: close:", err)
	}
}
