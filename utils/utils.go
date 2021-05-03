package utils

import (
	"bytes"
	"regexp"
	"strings"
	"unicode"
)

// ToLowerFirstCamelCase returns the given string in camelcase formatted string
// but with the first letter being lowercase.
func ToLowerFirstCamelCase(s string) string {
	if s == "" {
		return s
	}
	if len(s) == 1 {
		return strings.ToLower(string(s[0]))
	}
	return strings.ToLower(string(s[0])) + ToCamelCase(s)[1:]
}

// ToUpperFirst returns the given string with the first letter being uppercase.
func ToUpperFirst(s string) string {
	if s == "" {
		return s
	}
	if len(s) == 1 {
		return strings.ToLower(string(s[0]))
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

// ToLowerSnakeCase the given string in snake-case format.
func ToLowerSnakeCase(s string) string {
	return strings.ToLower(ToSnakeCase(s))
}

var camelCaseRex = regexp.MustCompile(`[\p{L}\p{N}]+`)

// ToCamelCase converts from underscore separated form to camel case form.
func ToCamelCase(s string) string {
	byteSrc := []byte(s)
	chunks := camelCaseRex.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		chunks[idx] = bytes.Title(val)
	}
	return string(bytes.Join(chunks, nil))
}

// ToSnakeCase converts from camel case form to underscore separated form.
func ToSnakeCase(s string) string {
	s = ToCamelCase(s)
	runes := []rune(s)
	length := len(runes)
	var out []rune
	for i := 0; i < length; i++ {
		out = append(out, unicode.ToLower(runes[i]))
		if i+1 < length && (unicode.IsUpper(runes[i+1]) && unicode.IsLower(runes[i])) {
			out = append(out, '_')
		}
	}

	return string(out)
}
