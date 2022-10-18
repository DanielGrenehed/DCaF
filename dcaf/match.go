package main

import (
	"fmt"
	"regexp"
)


type Match func(string) bool
/*
	Match column to regex and delimiter
*/
type DataSlice struct {
	match Match
	delim rune
}

func MatchNumber() (Match) {
	re := regexp.MustCompile("^[+-]?([0-9]+([.][0-9]*)?|[.][0-9]+)$")
	return func(s string) bool {return re.MatchString(s)}
}

func MatchDate() (Match) {
	re := regexp.MustCompile("^(19|20)[0-9]{2}[- /.](0[1-9]|1[012])[- /.](0[1-9]|[12][0-9]|3[01])$")
	return func(s string) bool {return re.MatchString(s)}
}

func MatchTime() (Match) {
	re := regexp.MustCompile("^(2[0-3]|[01]?[0-9]):([0-5]?[0-9]):([0-5]?[0-9])$")
	return func(s string) bool {return re.MatchString(s)}
}

func MatchInteger() (Match) {
	re := regexp.MustCompile("^([+-]?[1-9][0-9]*|0)$")
	return func(s string) bool {return re.MatchString(s)}
}

func MatchAll() (Match) {
	return func(s string) bool {return true}
}

func getMatchingFunction(_type rune) Match {
	switch _type {
	case 'D':
		return MatchDate()
	case 'T':
		return MatchTime()
	case 'N':
		return MatchNumber()
	case 'I':
		return MatchInteger()
	case 'A':
		return MatchAll()
	}
	fmt.Println("No match for type", string(_type))
	return func(s string) bool {return true}
}

func isAlphaNumeric(char rune) bool {
	if char >= 'a' && char <= 'z' {
		return true
	}
	if char >= 'A' && char <= 'Z' {
		return true
	}
	if char >= '0' && char <= '9' {
		return true
	}
	return false
}

func constructDataMatcher(patterns string) []DataSlice {
	var out []DataSlice
	/*
		Create array of regex strings to match data in file
	*/

	var vc DataSlice
	for _, char := range patterns {
		if isAlphaNumeric(char) {
			vc.match = getMatchingFunction(char)
		} else {
			vc.delim = char
			if vc.match == nil {
				vc.match = MatchAll()
			}
			out = append(out, vc)
			vc = DataSlice{}
		}
	}

	if vc.match != nil {
		out = append(out, vc)
	}

	return out
}