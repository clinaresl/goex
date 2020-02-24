// decide whether a number is divisible by 7
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
)

// To determine whether a number n is divisible by 7 follow this
// procedure:
//
// 1. Split the number in two parts: n1 with all digits but the units; n2 which
// contains only the units
// 2. If (n1-n2) is know to be divisible by 7, stop: n is divisible betwen 7.
// 4. Otherwise, n is divisible by 7 if and only if (n1-n2) is divisible by 7

// globals
// ----------------------------------------------------------------------------
const EXIT_SUCCESS = 0
const EXIT_FAILURE = 1

const version = "0.1"

var number int
var want_version bool

// functions
// ----------------------------------------------------------------------------

// init module
//
// setup the flag environment for the on-line help
func init() {

	// first, create a command-line argument for parsing the number
	flag.IntVar(&number, "number", math.MaxInt64, "number given to decide whether it is divisible by 7")

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

// divisible7
//
// return true if and only if the given number is divisible by 7 and false
// otherwse
func divisible7(n int) (result bool) {

	// make sure to consider only numbers in positive form
	if n < 0 {
		n *= -1
	}

	// if n equals 0 or 7, we are done! it is divisible by 7
	if n == 0 || n == 7 {
		result = true
	} else if n < 10 {
		result = false
	} else {
		// otherwise, apply the rule for deciding the divisibility by 7
		result = divisible7(int(n/10) - 2*(n%10))
	}

	// and return the value computed so far
	return
}

// main function
//
// given a number decide whether it is divisible by 7 or not
func main() {

	// first things first, parse the flags
	flag.Parse()

	// if the current version is requested, then show it on the standard output
	// and exit
	if want_version {
		showVersion(EXIT_SUCCESS)
	}

	// secondly, verify whether a number was given or not. Note the trick: the
	// flag package does not allow the definition of mandatory arguments, thus
	// the idea consists of giving mandatory arguments sensible default values
	// that could be verified in run time. Certainly, this disallows the user to
	// compute whether math.MaxInt64 is divisible by 7
	if number == math.MaxInt64 {

		log.Fatalf(" Use -number to provide a number")
		os.Exit(EXIT_FAILURE)
	}

	// Tell the user whether the given number is divisible by 7 or not
	if divisible7(number) {
		fmt.Printf(" The number %v is divisible by 7!\n", number)

		// verify if you want that your assessment is correct
		if number%7 != 0 {
			log.Fatal(" Oh, oh, something went deeply wrong")
		}
	} else {
		fmt.Printf(" The number %v is not divisible by 7 and the remainder, indeed is %v\n", number, number%7)
	}
}
