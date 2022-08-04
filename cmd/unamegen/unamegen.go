package main

import (
	"bytes"
	_ "embed"
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	flag "github.com/spf13/pflag"

	"github.com/thde/unamegen"
	"github.com/thde/unamegen/probs"
)

//go:generate sh -c "go run ../calculate/calculate.go ../../dictionaries/en.txt > en.gob"

//go:embed en.gob
var en []byte

var (
	version string
	commit  string
	date    string

	min        = flag.Int("min", 6, "min length of fake word")
	max        = flag.Int("max", 12, "max length of fake word")
	in         = flag.StringP("input", "i", "", "Input file")
	noColumns  = flag.BoolP("no-columns", "1", false, "Don't print the generated usernames in columns")

	help = flag.BoolP("help", "h", false, "print help message")
)

func run() error {
	flag.Parse()

	if *help {
		fmt.Fprintf(os.Stderr, "Usage of %s [ amount ]:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "(%s, %s %s)\n", version, commit, date)
		return nil
	}

	amount := 28
	if flag.Arg(0) != "" {
		a, err := strconv.Atoi(flag.Arg(0))
		if err != nil {
			return fmt.Errorf("amount error: %w", err)
		}
		amount = a
	}

	if *min > *max {
		return fmt.Errorf("min larger than max")
	}

	var w unamegen.WordProbability
	if *in == "" {
		dec := gob.NewDecoder(bytes.NewBuffer(en))
		dec.Decode(&w)
	} else {
		file, err := os.Open(*in)
		if err != nil {
			return fmt.Errorf("file %s: %w", *in, err)
		}
		defer file.Close()

		p := probs.Words{}
		w = unamegen.WordProbability(p.Read(file).Calculate())
	}

	// if amount is negative, it will repeat forever
	words := []string{}
	for i := 0; amount < 0 || i < amount; i++ {
		words = append(words, w.GenerateFakeWord(*min, *max))
	}

	if *noColumns {
		fmt.Println(strings.Join(words, "\n"))
	} else {
		table(os.Stdout, words, 4)
	}

	return nil
}

func table(w io.Writer, input []string, cols int) {
	maxWidth := 0
	for _, s := range input {
		if len(s) > maxWidth {
			maxWidth = len(s)
		}
	}
	format := fmt.Sprintf("%%-%ds%%s ", maxWidth)

	rows := (len(input) + cols - 1) / cols
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			i := col*rows + row
			if i >= len(input) {
				break // This means the last column is not "full"
			}
			padding := ""
			if i < 7 {
				padding = " "
			}
			fmt.Fprintf(w, format, input[i], padding)
		}
		fmt.Fprintln(w)
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}
}
