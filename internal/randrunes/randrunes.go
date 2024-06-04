package randrunes

import (
	"flag"
	"math/rand"
	// "unicode/utf8"
)

var randBytes bool
var randRunes bool

func init () {
	flag.BoolVar(&randBytes, "rand-bytes", false, "RandBytes().")
	flag.BoolVar(&randRunes, "rand-runes", false, "RandRunes().")
}


func RandBytes(outputLength int, inputText string) (tokens string) {
	if !randBytes {
		return ""
	}
	for i := 0; i < outputLength; i++ {
		tokens += string(inputText[rand.Intn(outputLength)])
	}
	return
}


func RandRunes(outputLength int, inputText string) (tokens string) {
	if !randRunes {
		return ""
	}
	// runesLength := utf8.RuneCountInString(inputText)
	runes := []rune(inputText)
	runesLength := len(runes)
	for i := 0; i < outputLength; i++ {
		tokens += string(runes[rand.Intn(runesLength)])
	}
	return
}
