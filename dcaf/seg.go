package main

import (
	"strings"
)

type LineJoint struct {
	segment int
	delim rune
}

func constructDataJoiner(pattern string) []LineJoint {
	var out []LineJoint
	/*
		Convert pattern string to a series of LineJoints 
		where all number-sequences becomes a LineJoints.segment
		and the symbol(s) following the number becomes the LineJoints.delim
	*/

	return out
}

func createMatchingDataJoiner(data_matcher []DataSlice, joint rune) []LineJoint {
	var out []LineJoint
	for i, _ := range data_matcher {
		if i == len(data_matcher)-1 {
			out = append(out, LineJoint{i, rune(0)})
		} else {
			out = append(out, LineJoint{i, joint})
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

func getLineJointDesegmenter(rules []LineJoint) (Desegmenter) {
	
	return func(s []string) string {
		if s == nil {
			return ""
		}
		out := ""
		for _, joint := range rules {
			
			if joint.segment >= 0 && joint.segment < len(s) {
				out += s[joint.segment] 
				if joint.delim != rune(0) {
					out += string(joint.delim)
				}
				
			}
			//fmt.Println(out)
		}
		if !strings.HasSuffix(out, "\n") {
			out += "\n"
		}
		return out
	}
}

func getDataSliceSegmenter(data_matcher []DataSlice) (Segmenter) {
	
	return func(s string) []string {
		var segments []string
		str := s
		for i, col_match := range data_matcher {

			/*
				'If sep does not appear in s, cut returns s, "", false. '
			*/
			segment, rest, found := strings.Cut(str, string(col_match.delim))

			/*
				Get all the wanted columns of data in file
				validate the data
				
				get the segments one at a time

				for each segment, validate the data
				if match fails, return no segments

				if last delimiter is not found, that is ok, just 
				use the remaning string as the column data

			*/
			if i == len(data_matcher)-1 || found {
				if col_match.match(segment) {
					segments = append(segments, segment)
				} else {
					return nil
				}
			} else if !found {
				return nil
			}

			str= rest
		}
		return segments
	}
}