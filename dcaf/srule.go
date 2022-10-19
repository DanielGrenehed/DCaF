package main

import (
	"fmt"
	"regexp"
)


type Match func(string) bool
/*
	Match column to regex and delimiter
*/
type SliceRule struct {
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

func MatchRegex(regex string) (Match) {
	re := regexp.MustCompile(regex)
	return func(s string) bool {return re.MatchString(s)}
}

type CustomType struct {
	regex string
	symbol rune
}

func getMatchingFunction(symbol rune, types []CustomType) Match {
	switch symbol {
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

	for _, c_type := range types {
		if symbol == c_type.symbol {
			return MatchRegex(c_type.regex)
		}
	}

	fmt.Println("No match for type", string(symbol))
	return MatchAll()
}

func constructSliceRules(patterns string, c_types []CustomType) []SliceRule {
	var out []SliceRule
	/*
		Create array of regex strings to match data in file
	*/

	var vc SliceRule
	for _, char := range patterns {
		if isAlphaNumeric(char) {
			vc.match = getMatchingFunction(char, c_types)
		} else {
			vc.delim = char
			if vc.match == nil {
				vc.match = MatchAll()
			}
			out = append(out, vc)
			vc = SliceRule{}
		}
	}

	if vc.match != nil {
		out = append(out, vc)
	}

	return out
}

func removeQuotation(str string) string {
	re := regexp.MustCompile("^['\"]*['\"]$")
	if re.MatchString(str) {
		return str[1:len(str)-1]
	}
	return str
}

func createCustomTypes(input []string) []CustomType {
	var types []CustomType
	re := regexp.MustCompile("^[a-zA-Z]:")
	for _, str := range input {
		if re.MatchString(str) {
		
			var ct CustomType
			ct.symbol = []rune(str)[0]
			ct.regex = removeQuotation(str[2:])
			types = append(types, ct)
		}
	}
	return types
}