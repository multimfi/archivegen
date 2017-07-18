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

var buildversion string

var (
	flagOut     = flag.String("out", "out.archive", "Output file, empty for stdout.")
	flagFormat  = flag.String("fmt", "cpio", "Output file format, cpio/tar.")
	flagPrint   = flag.Bool("print", false, "Do not write outfile, print resolved tree.")
	flagVersion = flag.Bool("v", false, "Print version.")
)

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile)

	if *flagVersion {
		fmt.Printf("build: %s\nruntime: %s\n", buildversion, runtime.Version())
		return
	}

	if flag.NArg() < 1 {
		log.Fatal("not enough arguments")
	}

	var (
		outw io.WriteCloser
		w    archive.Writer
	)

	if *flagOut == "" {
		outw = os.Stdout
	} else {
		f, err := os.OpenFile(*flagOut, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatal(err)
		}
		outw = f
	}
	defer outw.Close()

	switch *flagFormat {
	case "cpio":
		w = cpio.NewWriter(outw)
	case "tar":
		w = tar.NewWriter(outw)
	}
	defer w.Close()

	m, err := config.FromFiles(flag.Args()...)
	if err != nil {
		log.Fatal(err)
	}

	t := tree.Render(m)
	if *flagPrint {
		w := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)
		t.Print("", w)
		w.Flush()
		os.Exit(0)
	}
	if err := t.Write("", w); err != nil {
		log.Fatal(err)
	}
}
