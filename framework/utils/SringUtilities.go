package utils

import (
	"math/rand"
	"strings"
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
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, number)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
