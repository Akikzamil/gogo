package utils

import (
	"regexp"
	"strings"
	"unicode"
)

func ToPascelCase(input string) string {
	// Split the string into words
	words := strings.FieldsFunc(input, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})

	// Capitalize the first letter of each word
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}

	// Join the words to form UpperCamelCase
	return strings.Join(words, "")
}

func ToSnakeCase(input string) string {
	// Use a regex to add a space before each capital letter (except the first one)
	regex := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	input = regex.ReplaceAllString(input, `${1} ${2}`)

	// Split the string into words
	words := strings.FieldsFunc(input, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})

	// Convert all words to lowercase
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	// Join the words with underscores
	return strings.Join(words, "_")
}


