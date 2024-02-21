package stringUtils

import (
	"math/rand"
	"regexp"
	"strings"
)

func Pointer(s string) *string {
	return &s
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func RandomString(lenght int, charset ...string) string {
	if len(charset) == 0 {
		charset = []string{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"}
	}
	var letters = []byte(charset[0])

	s := make([]byte, lenght)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func IsDigit(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func SubstringBetween(txt, start, end string) string {
	s := strings.Index(txt, start)
	if s == -1 {
		s = 0
	} else {
		s += len(start)
	}

	// Trova l'indice della parola finale
	e := strings.Index(txt[s:], end)
	if e == -1 {
		e = len(txt)
	} else {
		e += s
	}
	return txt[s:e]
}

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func RemoveNonAlphanumeric(s string) string {
	return nonAlphanumericRegex.ReplaceAllString(s, "")
}
