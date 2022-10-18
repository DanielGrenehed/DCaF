package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)



type Filter func(string) string

func process(r *bufio.Scanner, w *bufio.Writer, f Filter) {
	for r.Scan() {
		line := f(r.Text())
		if line != "" {
			w.WriteString(line)
		}
	}
}




func parseLine(line string, sep Segmenter, con Desegmenter) string {
	segments := sep(line)
	//fmt.Println(segments)
	str := con(segments)
	return str
}


func main() {
	in_file_ptr := flag.String("i", "", "file to read")
	out_file_ptr := flag.String("o", "", "file to write")
	dms_ptr := flag.String("c", "", "format of data in input")
	djs_ptr := flag.String("r", "", "reconstruction format of line")
	djoint_ptr := flag.String("j", ",", "default char to join data")
	explicit_ptr := flag.Bool("e", false, "Print omitted lines")
	flag.Parse()

	if *explicit_ptr {
		fmt.Println("Explicit output enabled")
	}

	/*
		Create data flow
	*/
	data_matcher := constructDataMatcher(*dms_ptr)
	seg := getDataSliceSegmenter(data_matcher)

	data_join_string := *djs_ptr
	data_joiner := constructDataJoiner(data_join_string)

	default_joint := []rune(*djoint_ptr)
	if data_join_string == "" {
		data_joiner = createMatchingDataJoiner(data_matcher, default_joint[0])
	}
	des := getLineJointDesegmenter(data_joiner)
	var f = func(s string) string {return parseLine(s, seg, des)}


	/*
		Open files
	*/
	dir, _ := os.Getwd()

	if len(*in_file_ptr) == 0 {
		fmt.Println("No input file provided")
		return
	}
	if len(*out_file_ptr) == 0 {
		fmt.Println("No output file provided")
		return
	}

	input_path := dir + "/" + (*in_file_ptr)
	rf, err := os.Open(input_path)
	if err != nil {
		fmt.Println(err)
		return 
	}
	r := bufio.NewScanner(rf)
	r.Split(bufio.ScanLines)

	output_path := dir + "/" + (*out_file_ptr)
	fout, err := os.Create(output_path)
	if err != nil {
		fmt.Println(err)
		rf.Close()
		return
	}
	w := bufio.NewWriter(fout)	
	
	/*
		process data line by line
	*/
	process(r, w, f)

	/*
		clean up
	*/
	rf.Close()
	w.Flush()
	fout.Close()
}
