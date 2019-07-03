package snaker

// Package snaker provides methods to convert CamelCase names to snake_case and back.
// It considers the list of allowed initialsms used by github.com/golang/lint/golint (e.g. ID or HTTP)

import (
	"strings"
	"unicode"
)

// CamelToSnake converts a given string to snake case
func CamelToSnake(s string) string {
	var result string
	var words []string
	var lastPos int
	rs := []rune(s)

	for i := 0; i < len(rs); i++ {
		if i > 0 && unicode.IsUpper(rs[i]) {
			if initialism := startsWithInitialism(s[lastPos:]); initialism != "" {
				words = append(words, initialism)

				i += len(initialism) - 1
				lastPos = i
				continue
			}

			words = append(words, s[lastPos:i])
			lastPos = i
		}
	}

	// append the last word
	if s[lastPos:] != "" {
		words = append(words, s[lastPos:])
	}

	for k, word := range words {
		if k > 0 {
			result += "_"
		}

		result += strings.ToLower(word)
	}

	return result
}

func snakeToCamel(s string, upperCase bool) string {
	if len(s) == 0 {
		return s
	}
	var result string

	words := strings.Split(s, "_")

	//// if there is no underscore, first try commons and then just return
	//if len(words) == 1 {
	//	if exception := snakeToCamelExceptions[words[0]]; len(exception) > 0 {
	//		return exception
	//	}
	//
	//	if upperCase {
	//		if upper := strings.ToUpper(words[0]); commonInitialisms[upper] {
	//			return upper
	//		}
	//	}
	//
	//	w := []rune(s)
	//	if upperCase {
	//		w[0] = unicode.ToUpper(w[0])
	//	} else {
	//		w[0] = unicode.ToLower(w[0])
	//	}
	//
	//	return string(w)
	//}

	for i, word := range words {
		if exception := snakeToCamelExceptions[word]; len(exception) > 0 {
			result += exception
			continue
		}

		if upperCase || i > 0 {
			if upper := strings.ToUpper(word); commonInitialisms[upper] {
				result += upper
				continue
			}
		}

		if upperCase || i > 0 {
			result += camelizeWord(word, len(words) > 1)
		} else {
			result += word
		}
	}

	return result
}

func camelizeWord(word string, force bool) string {
	runes := []rune(word)

	for i, r := range runes {
		if i == 0 {
			runes[i] = unicode.ToUpper(r)
		} else {
			if !force && unicode.IsLower(r) { // already camelCase
				return string(runes)
			}

			runes[i] = unicode.ToLower(r)
		}
	}

	return string(runes)
}

// SnakeToCamel returns a string converted from snake case to uppercase
func SnakeToCamel(s string) string {
	return snakeToCamel(s, true)
}

// SnakeToCamelLower returns a string converted from snake case to lowercase
func SnakeToCamelLower(s string) string {
	return snakeToCamel(s, false)
}

// startsWithInitialism returns the initialism if the given string begins with it
func startsWithInitialism(s string) string {
	var initialism string
	// the longest initialism is 5 char, the shortest 2
	for i := 1; i <= 5; i++ {
		if len(s) > i-1 && commonInitialisms[s[:i]] {
			initialism = s[:i]
		}
	}
	return initialism
}

// commonInitialisms, taken from
// https://github.com/golang/lint/blob/206c0f020eba0f7fbcfbc467a5eb808037df2ed6/lint.go#L731
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"ETA":   true,
	"GPU":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"OS":    true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
	"OAuth": true,
}

// add exceptions here for things that are not automatically convertable
var snakeToCamelExceptions = map[string]string{
	"oauth": "OAuth",
}