package main

import (
	"compress/gzip"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"os"

	flag "github.com/spf13/pflag"
	"thde.io/fakeword"
)

var (
	version string
	commit  string
	date    string

	jsonEncoding    = flag.Bool("json", false, "Use JSON encoding instead of GOB")
	gzipCompression = flag.Bool("gzip", false, "Use GZIP compression")
	help            = flag.BoolP("help", "h", false, "Print help message")
)

type encoder interface {
	Encode(v any) error
}

func run(out io.Writer, errOut io.Writer) error {
	flag.Parse()
	flag.CommandLine.SetOutput(errOut)

	if *help {
		fmt.Fprintf(errOut, "Usage of %s (%s, %s %s):\n", os.Args[0], version, commit, date)
		flag.PrintDefaults()
		return nil
	}

	if *gzipCompression {
		out = gzip.NewWriter(out)
	}

	var enc encoder
	if *jsonEncoding {
		jenc := json.NewEncoder(out)
		jenc.SetIndent("", "  ")
		enc = jenc
	} else {
		enc = gob.NewEncoder(out)
	}

	w := fakeword.Dictionary{}
	return enc.Encode(w.Read(os.Stdin).Generator())
}

func main() {
	if err := run(os.Stdout, os.Stderr); err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}
}
