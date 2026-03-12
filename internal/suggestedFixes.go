package internal

import (
	"strings"
	"unicode"
)

func toLowerCase(s string) (string) {
	if len(s) == 0 {
		return ""
	}
	return strings.Replace(s, string(s[0]), strings.ToLower(string(s[0])), 1)
}

func toStandardSymbols(s string) string {
	var b strings.Builder

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == ' ' || r == '_' || r == '-' {
			b.WriteRune(r)
		} else {
			b.WriteRune(' ')
		}
	}

	return strings.TrimSpace(b.String())
}