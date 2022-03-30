package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	flag "github.com/spf13/pflag"

	"github.com/thde/unamegen"
)

var (
	version string
	commit string
	date string

	min        = flag.Int("min", 6, "min length of fake word")
	max        = flag.Int("max", 12, "max length of fake word")
	capitalize = flag.BoolP("capitalize", "c", false, "capitalize the first letter")

	help = flag.BoolP("help", "h", false, "print help message")
)

func run() error {
	flag.Parse()

	if *help {
		fmt.Fprintf(os.Stderr, "Usage of %s (%s, %s %s) [ amount ]:\n", os.Args[0], version, commit, date)
		flag.PrintDefaults()
		return nil
	}

	amount := 10
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

	w, err := unamegen.New()
	if err != nil {
		return fmt.Errorf("character initialization error: %w", err)
	}

	// If nWords is negative, repeat forever
	// NOTE: i can be overflow, but it doesn't cause runtime error
	for i := 0; amount < 0 || i < amount; i++ {
		fakeWord := w.GenerateFakeWord(*min, *max)
		if *capitalize {
			fakeWord = strings.Title(fakeWord) //nolint
		}
		fmt.Println(fakeWord)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}
}
