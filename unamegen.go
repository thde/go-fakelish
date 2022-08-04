package unamegen

import (
	"math/rand"
	"strings"
	"time"
)

const (
	MaxSequences = 2
	Prefix = "^"
	Suffix = "$"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

type (
	WordProbability map[string]map[string]float32
)

func (w WordProbability) GenerateFakeWordWithUnexpectedLength() string {
	character := Prefix
	word := ""
	characters := []string{}
	for character != Suffix {
		characters = append(characters, character)
		if len(characters) > MaxSequences {
			characters = characters[1:]
		}
		var nextAccumedProbs map[string]float32
		n := 0
		for {
			str := strings.Join(characters[n:], "")
			nextAccumedProbs = w[str]
			n += 1
			if (nextAccumedProbs != nil || n >= len(characters)) {
				break
			}
		}
		nextCharacter := ""
		r := random.Float32()
		probability := float32(0)
		for ch, prob := range nextAccumedProbs {
			nextCharacterCandidate := ch
			probability += prob
			if r <= probability {
				nextCharacter = nextCharacterCandidate
				break
			}
		}
		if nextCharacter != Suffix {
			word += nextCharacter
		}
		character = nextCharacter
	}
	return word
}

func (w WordProbability) GenerateFakeWordByLength(length int) string {
	fakeWord := ""
	for len(fakeWord) != length {
		fakeWord = w.GenerateFakeWordWithUnexpectedLength()
	}
	return fakeWord
}

func (w WordProbability) GenerateFakeWord(minLength int, maxLength int) string {
	fakeWord := ""
	for (minLength > len(fakeWord) || len(fakeWord) > maxLength) {
		fakeWord = w.GenerateFakeWordWithUnexpectedLength()
	}
	return fakeWord
}
