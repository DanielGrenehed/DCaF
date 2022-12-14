package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

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

type Function func() ()

type DcafModel struct {
	close_input Function
	close_output Function
	scanner *bufio.Scanner
	writer *bufio.Writer

	filter Filter
	fail bool
	message string
} 

func (model DcafModel) cleanup() {
	if model.close_input != nil {
		model.close_input()
	}
	if model.writer != nil {
		model.writer.Flush()
	}
	if model.close_output != nil {
		model.close_output()
	}
}

func (model DcafModel) setInputFile(dir string, file string) DcafModel {
	if model.fail {
		return model
	}
	if len(file) == 0 {
		model.message = "No input file provided"
		model.fail = true
		return model
	}
	input_file_path := dir + "/" + (file)
	input_file, err := os.Open(input_file_path)
	if err != nil {
		model.fail = true
		model.message = err.Error()
		return model
	} else {
		model.close_input = func() {input_file.Close()}
	}
	model.scanner = bufio.NewScanner(input_file)
	model.scanner.Split(bufio.ScanLines)
	return model
}

func (model DcafModel) setOutputFile(dir string, file string, append bool) DcafModel {
	if model.fail {
		return model
	}
	if len(file) == 0 {
		model.message = "No output file provided"
		model.fail = true
		return model
	}
	output_file_path := dir + "/" + (file)

	var output_file *os.File
	var err error
	if append {
		output_file, err = os.OpenFile(output_file_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	} else {
		output_file, err = os.Create(output_file_path)
	}
	
	if err != nil {
		model.fail = true
		model.message = err.Error()
		return model
	} else {
		model.close_output = func() {output_file.Close()}
	}
	model.writer = bufio.NewWriter(output_file)	
	return model
}

func constructDcafModel() DcafModel {
	var model DcafModel

	/*
		Parse input
	*/
	input_file_ptr := flag.String("i", "", "file to read")
	output_file_ptr := flag.String("o", "", "file to write")

	data_match_string_ptr := flag.String("c", "", "format of data in input")
	data_join_string_ptr := flag.String("r", "", "reconstruction format of line")
	default_joint_ptr := flag.String("j", ",", "default char to join data")
	append_flag_ptr := flag.Bool("a", false, "use flag to append to file")

	flag.Parse() 
	//fmt.Println(rest)
	model.filter = constructFilter(*data_match_string_ptr, *data_join_string_ptr, *default_joint_ptr, flag.Args())

	/*
		Open files
	*/
	dir, _ := os.Getwd()

	model = model.setInputFile(dir, *input_file_ptr)
	model = model.setOutputFile(dir, *output_file_ptr, *append_flag_ptr)
	return model
}



func main() {
	model := constructDcafModel()
	
	if model.fail {
		fmt.Println(model.message)
		model.cleanup()
		return
	}

	/*
		process data line by line
	*/
	process(model.scanner, model.writer, model.filter)

	model.cleanup()
}
