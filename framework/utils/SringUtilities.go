package utils

import (
	"math/rand"
	"strings"
	"time"
)

type Strings struct{}

type PlaceholderReplacer struct {
	Find      string
	ReplaceTo string
}

func (_ Strings) Replace(haystack string, needles []PlaceholderReplacer) string {
	for _, placeholder := range needles {
		haystack = strings.ReplaceAll(haystack, placeholder.Find, placeholder.ReplaceTo)
	}

	return haystack
}

func (_ Strings) Random(number int) string {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a slice of runes
	runes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// Create a byte slice to hold the random bytes
	b := make([]byte, number)

	// Generate random bytes
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}

	// Create a string from the random bytes
	s := ""
	for _, v := range b {
		s += string(runes[int(v)%len(runes)])
	}

	return s
}
