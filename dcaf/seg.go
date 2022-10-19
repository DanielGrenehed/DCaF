package main

import (
	"fmt"
	"strconv"
	"strings"
)

type JoinRule struct {
	segment int
	delim string
}



func findNextDigit(start int, str string) int {
	for pos, char := range str[start:] {
		if isNumeric(char) {
			return start + pos
		}
	}
	return -1
}

func findNextNonDigit(start int, str string) int {
	for pos, char := range str[start:] {
		if !isNumeric(char) {
			return start + pos
		}
	}
	return -1
}

func constructJoinRules(pattern string) []JoinRule {
	var out []JoinRule
	/*
		Convert pattern string to a series of LineJoints 
		where all number-sequences becomes a LineJoints.segment
		and the symbol(s) following the number becomes the LineJoints.delim
	*/
	if len(pattern) == 0 {
		return out
	}

	start := findNextDigit(0, pattern)
	for start >= 0 {
		var rule JoinRule
		n_end := findNextNonDigit(start, pattern)
		if n_end == -1 {
			rule.segment , _ = strconv.Atoi(pattern[start:len(pattern)])
			out = append(out, rule)
			break
		}

		rule.segment , _ = strconv.Atoi(pattern[start:n_end])

		d_start := findNextDigit(n_end, pattern)
		if d_start == -1 {
			rule.delim = pattern[n_end:len(pattern)]			
		} else {
			rule.delim = pattern[n_end:d_start]
		}
		
		out = append(out, rule)
		start = d_start
	}
	return out
}

func createMatchingJoinRules(slice_rules []SliceRule, joint string) []JoinRule {
	var out []JoinRule
	for i, _ := range slice_rules {
		if i == len(slice_rules)-1 {
			out = append(out, JoinRule{i, ""})
		} else {
			out = append(out, JoinRule{i, joint})
		}
		//fmt.Println(out[i])
	}
	return out
}

/*
	A function that takes a string and converts it
	into a series of substrings
*/
type Segmenter func(string) []string

/*
	A function that takes a series of strings and 
	constructs a single string from them
*/
type Desegmenter func([]string) string 

func getJoinRuleDesegmenter(rules []JoinRule) (Desegmenter) {
	
	return func(s []string) string {
		if s == nil {
			return ""
		}

		out := ""
		for _, joint := range rules {
			
			if joint.segment >= 0 && joint.segment < len(s) {
				out += s[joint.segment] + joint.delim
			} else {
				fmt.Println("Invalid segment ", joint.segment)
			}
		}
		
		if len(s) == 1 && len(rules) == 0 {
			out += s[0]
		}

		if !strings.HasSuffix(out, "\n") {
			out += "\n"
		}
		return out
	}
}

func getSliceRuleSegmenter(slice_rules []SliceRule) (Segmenter) {
	
	return func(s string) []string {
		var segments []string
		
		if len(slice_rules) == 0 {
			segments = append(segments, s)
			return segments
		}

		str := s
		for i, rule := range slice_rules {

			/*
				'If sep does not appear in s, cut returns s, "", false. '
			*/
			segment, rest, found := strings.Cut(str, string(rule.delim))

			/*
				Get all the wanted columns of data in file
				validate the data
				
				get the segments one at a time

				for each segment, validate the data
				if match fails, return no segments

				if last delimiter is not found, that is ok, just 
				use the remaning string as the column data

			*/
			if i == len(slice_rules)-1 || found {
				if rule.match(segment) {
					segments = append(segments, segment)
				} else {
					return nil
				}
			} else if !found {
				return nil
			}

			str = rest
		}
		return segments
	}
}