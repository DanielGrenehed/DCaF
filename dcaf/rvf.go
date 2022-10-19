package main

import "unicode"

func isNumeric(char rune) bool {
	return unicode.IsDigit(char)
}

func isAlphabetic(char rune) bool {
	return unicode.IsLetter(char)
}

func isAlphaNumeric(char rune) bool {
	if isAlphabetic(char) {
		return true
	}
	return isNumeric(char)
}