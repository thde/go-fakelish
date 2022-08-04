package probs

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const DefaultPrefixLength = 4

type Words struct {
	PrefixLength int
	counter      map[string]map[string]int
}

func (w *Words) Add(words ...string) *Words {
	if w.PrefixLength == 0 {
		w.PrefixLength = DefaultPrefixLength
	}

	for _, word := range words {
		word := strings.ToLower(strings.TrimSpace(word))
		word = fmt.Sprintf("^%s$", word)

		for i := 2; i <= w.PrefixLength; i++ {
			for _, substr := range splitToLength(word, i) {
				w.count(substr)
			}
		}
	}

	return w
}

func (w *Words) Read(in io.Reader) *Words {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		l := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(l, "#") {
			continue
		}

		words := strings.Fields(scanner.Text())
		w.Add(words...)
	}

	return w
}

func (w *Words) Calculate() map[string]map[string]float32 {
	m := map[string]map[string]float32{}

	for prefix, suffix := range w.counter {
		results := map[string]float32{}
		var sum int

		for _, c := range suffix {
			sum += c
		}

		for s, c := range suffix {
			results[s] = float32(c) / float32(sum)
		}

		m[prefix] = results
	}

	return m
}

// count counts the amount of occurencies of a suffix
func (w *Words) count(substr string) {
	prefix := substr[:len(substr)-1]
	suffix := substr[len(substr)-1:]

	if w.counter == nil {
		m := map[string]map[string]int{}
		w.counter = m
	}

	_, ok := w.counter[prefix]
	if !ok {
		w.counter[prefix] = map[string]int{}
	}

	w.counter[prefix][suffix] += 1
}

// splitToLength splits a string to substrings of length
func splitToLength(s string, length int) []string {
	substrs := []string{}

	for i := 0; i <= len(s)-1; i++ {
		j := i + length
		if j > len(s) {
			continue
		}

		substrs = append(substrs, s[i:j])
	}

	return substrs
}
