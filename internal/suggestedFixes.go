package internal

import (
	"strings"
	"unicode"
)

func toLowerCase(s string) string {
	if s == "" {
		return ""
	}
	if len(s) >=2 && unicode.IsUpper(rune(s[0])) && unicode.IsUpper(rune(s[1])) {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	
	return string(r)
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

	return strings.Join(strings.Fields(b.String()), " ")
}