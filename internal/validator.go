package internal

import (
	"strings"
	"unicode"
)

var sensitiveWords = map[string]struct{}{
	"password": {},
	"passwd": {},
	"pwd": {},
	"token": {},
	"secret": {},
	"apikey": {},
	"api_key": {},
	"access_token": {},
	"refresh_token": {},
	"authorization": {},
}

func normalizeWords(s string) []string {
	var b strings.Builder

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(unicode.ToLower(r))
		} else {
			b.WriteRune(' ')
		}
	}

	return strings.Fields(b.String())
}

func notContainsSensitiveWord(s string) bool {
	words := normalizeWords(s)

	for _, w := range words {
		if _, ok := sensitiveWords[w]; ok {
			return false
		}
	}

	return true
}


func isEnglish(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.In(r, unicode.Latin) {
			return false
		}
	}
	return true
}

func notHasSpecialSymbols(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == ' ' || r == '_' || r == '-'{
			continue
		}
		return false
	}
	return true
}

func isSupportedMethodSlog(method string) bool {
	switch method {
	case "Debug", "Info", "Warn", "Error", "DebugContext", "InfoContext", "WarnContext", "ErrorContext":
		return true
	default:
		return false
	}
}

func checkLowerCase(arg string) bool {
	arg = strings.TrimSpace(arg)
	r := []rune(arg)

	if len(r) == 0 {
		return true
	}
	return !(unicode.IsLetter(r[0]) && unicode.IsUpper(r[0]))
}