package string_utils

import (
	"math/rand/v2"
	"regexp"
	"strings"
	"unicode"

	"github.com/pasa33/goutils/slice_utils"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
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
		s[i] = letters[rand.IntN(len(letters))]
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

func ShuffleString(s string) string {
	runes := []rune(s)
	slice_utils.ShuffleSlice(runes)
	return string(runes)
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

func RemoveAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, _ := transform.String(t, s)
	return output
}

func SplitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

func SplitTakeLast(s string, seps string) string {
	return s[strings.LastIndex(s, seps)+len(seps):]
}

func ContainsAnyStr(s string, substrings ...string) bool {
	for _, substr := range substrings {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}

func ContainsAnyStrInsensitive(s string, substrings ...string) bool {
	s = strings.ToLower(s)
	for _, substr := range substrings {
		if strings.Contains(s, strings.ToLower(substr)) {
			return true
		}
	}
	return false
}

func RandomizeTextCase(input string) string {
	var sb strings.Builder

	for _, char := range input {
		if rand.N(2) == 0 {
			sb.WriteRune(char)
		} else {
			if char >= 'a' && char <= 'z' {
				sb.WriteRune(char - 'a' + 'A')
			} else if char >= 'A' && char <= 'Z' {
				sb.WriteRune(char - 'A' + 'a')
			} else {
				sb.WriteRune(char)
			}
		}
	}

	return sb.String()
}
