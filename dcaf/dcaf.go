package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Filter func(string) string

/*
	Write filtered lines 
*/
func process(r *bufio.Scanner, w *bufio.Writer, f Filter) {
	for r.Scan() {
		line := f(r.Text())
		if line != "" {
			w.WriteString(line)
		}
	}
}

func reconstructionFilter(seg Segmenter, des Desegmenter) (Filter) {
	return func(s string) string {return des(seg(s))}
}


func main() {
	/*
		Parse input
	*/
	input_file_ptr := flag.String("i", "", "file to read")
	output_file_ptr := flag.String("o", "", "file to write")

	data_match_string_ptr := flag.String("c", "", "format of data in input")
	data_join_string_ptr := flag.String("r", "", "reconstruction format of line")

	default_joint_ptr := flag.String("j", ",", "default char to join data")
	
	flag.Parse() 

	/*
		Create data flow
	*/
	data_matcher := constructDataMatcher(*data_match_string_ptr)
	

	data_join_string := *data_join_string_ptr
	data_joiner := constructDataJoiner(data_join_string)

	if data_join_string == "" {
		data_joiner = createMatchingDataJoiner(data_matcher, ([]rune(*default_joint_ptr))[0])
	}

	seg := getDataSliceSegmenter(data_matcher)
	des := getLineJointDesegmenter(data_joiner)


	/*
		Open files
	*/
	dir, _ := os.Getwd()

	if len(*input_file_ptr) == 0 {
		fmt.Println("No input file provided")
		return
	}
	if len(*output_file_ptr) == 0 {
		fmt.Println("No output file provided")
		return
	}

	input_file_path := dir + "/" + (*input_file_ptr)
	input_file, err := os.Open(input_file_path)
	if err != nil {
		fmt.Println(err)
		return 
	}
	input_file_reader := bufio.NewScanner(input_file)
	input_file_reader.Split(bufio.ScanLines)

	output_file_path := dir + "/" + (*output_file_ptr)
	output_file, err := os.Create(output_file_path)
	if err != nil {
		fmt.Println(err)
		input_file.Close()
		return
	}
	output_file_writer := bufio.NewWriter(output_file)	
	
	/*
		process data line by line
	*/
	process(input_file_reader, output_file_writer, reconstructionFilter(seg, des))

	/*
		clean up
	*/
	input_file.Close()
	output_file_writer.Flush()
	output_file.Close()
}
