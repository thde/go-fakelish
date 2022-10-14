package main

import (
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
	help            = flag.BoolP("help", "h", false, "Print help message")
)

type Encoder interface {
	Encode(v any) error
}

func run(out io.Writer, errOut io.Writer) error {
	flag.Parse()
	flag.CommandLine.SetOutput(errOut)

	if *help {
		fmt.Fprintf(errOut, "Usage of %s (%s, %s %s) [ file ]:\n", os.Args[0], version, commit, date)
		flag.PrintDefaults()
		return nil
	}

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		return err
	}
	defer file.Close()

	var enc Encoder
	if *jsonEncoding {
		jenc := json.NewEncoder(out)
		jenc.SetIndent("", "  ")
		enc = jenc
	} else {
		enc = gob.NewEncoder(out)
	}

	w := fakeword.Dictionary{}
	return enc.Encode(w.Read(file).Generator())
}

func main() {
	if err := run(os.Stdout, os.Stderr); err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}
}
