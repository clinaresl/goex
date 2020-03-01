package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

// globals
// ----------------------------------------------------------------------------
const EXIT_SUCCESS = 0
const EXIT_FAILURE = 1

const version = "0.1"

// if a line is found in the first file, in1 should be enabled; likewise, if a
// line is found in the second file, in2 should be enabled as well. Those lines
// appearing in both files should then get the sum of both values, 3 and
// appearing in both files is the only way to get this sum
const (
	in1 = 1 << iota
	in2
)

// the following globals are used to capture the arguments
var file1, file2 string
var suppress [4]bool
var want_version bool

// functions
// ----------------------------------------------------------------------------

// init module
//
// setup the flag environment for the on-line help
func init() {

	// first, create a command-line argument for parsing the filenames
	flag.StringVar(&file1, "file1", "", "first file used in the comparison")
	flag.StringVar(&file2, "file2", "", "second file used in the comparison")

	// and also other arguments for allowing the user to suppress some output
	flag.BoolVar(&suppress[1], "1", false, "suppress column 1")
	flag.BoolVar(&suppress[2], "2", false, "suppress column 2")
	flag.BoolVar(&suppress[3], "3", false, "suppress column 3")

	// also, create an additional flag for showing the version
	flag.BoolVar(&want_version, "version", false, "shows version info and exits")
}

// showVersion
//
// show the current version of this program and exits with the given signal
func showVersion(signal int) {

	fmt.Printf(" %v %v\n", os.Args[0], version)
	os.Exit(signal)
}

// processFile
//
// process the given 'filename' and updates the contents of the map 'presence'
// giving the value 'token' to each line appearing in the file
func processFile(filename string, presence map[string]int, token int) {

	// open the file
	stream, err := os.Open(filename)
	if err != nil {

		log.Printf("It was not possible to open '%v'\n", filename)
		os.Exit(EXIT_FAILURE)
	}

	// note now how resources are released in Go, instead of awaiting until the
	// last line, closing the file is deferred so that it is invoked in spite of
	// the reason to exit. Moreover, passing a lambda function enables error
	// checking
	defer func() {
		if err = stream.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// create a reader to extract the lines from the file
	reader := bufio.NewScanner(stream)
	for reader.Scan() {

		// count this line in the 'presence' map
		presence[reader.Text()] += token
	}
}

// showResults
//
// show the results. If any of the flags suppress take the value true the
// corresponding output of that column is suppressed
func showResults(presence map[string]int, suppress [4]bool) {

	// for all values in the map
	for key, value := range presence {

		// if value==1 then the line in 'key' appears only in the first file; if
		// value==2 then it appears only in the second file; finally, if
		// value==3, then it appears in both. Hence, data is shown on the
		// standard output: first, leaving a number of tabs equal to the value -
		// 1; second, only if it is not suppressed
		if !suppress[value] {
			for nbtabs := 1; nbtabs < value; nbtabs++ {
				fmt.Print("\t")
			}
			fmt.Println(key)
		}
	}
}

// main function
//
// UNIX comm utility
func main() {

	// first things first, parse the flags
	flag.Parse()

	// if the current version is requested, then show it on the standard output
	// and exit
	if want_version {
		showVersion(EXIT_SUCCESS)
	}

	// secondly, verify that both mandatory arguments were given. Note in the
	// following if statements two different forms to verify whether a string is
	// empty or not. Both are valid and idiomatic of Go
	if len(file1) == 0 {

		log.Fatalf(" Use -file1 to provide the location a file to compare")
		os.Exit(EXIT_FAILURE)
	}

	if file2 == "" {

		log.Fatalf(" Use -file2 to provide the location the second file to compare")
		os.Exit(EXIT_FAILURE)
	}

	// process both files using the same map
	presence := make(map[string]int)
	processFile(file1, presence, in1)
	processFile(file2, presence, in2)

	// and show the results
	showResults(presence, suppress)

	// it is idiomatic of golang to do not return any value upon normal
	// termination
}
