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

func getTree(files []string) *tree.Node {
	r, err := config.FromFiles(files...)
	if err != nil {
		log.Fatal(err)
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

var (
	flagOut     = flag.String("out", "out.archive", "output file")
	flagFormat  = flag.String("fmt", "cpio", "file format, cpio/tar")
	flagPrint   = flag.Bool("print", false, "print resolved tree")
	flagStdout  = flag.Bool("stdout", false, "output to stdout")
	flagVersion = flag.Bool("version", false, "version")
)

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile)

	if *flagStdout {
		flagOut = nil
	}
	if *flagVersion {
		fmt.Printf("build: %s\nruntime: %s\n", buildversion, runtime.Version())
		return
	}
	if flag.NArg() < 1 {
		log.Fatal("not enough arguments")
	}

	root := getTree(flag.Args())

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
