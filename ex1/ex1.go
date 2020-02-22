// If we list all the natural numbers below 10 that are multiples of 3 or 5, we
// get 3, 5, 6 and 9.
//
// The sum of these multiples is 23.Find the sum of all the multiples of 3 or 5
// below 1000.
//
// (no intention to bother the great people behind projecteuler, this is just an
// incentive for others to play there!)

package main

import (
	"flag"
	"fmt"
	"os"
)

// globals
// ----------------------------------------------------------------------------

const EXIT_SUCCESS int = 0 // exit with success
const EXIT_FAILURE int = 1 // exit with failure

const version string = "0.1"

var bound int
var version_wanted bool

// functions
// ----------------------------------------------------------------------------

// initializes the command-line parser
func init() {

	// Flag to store the pgn file to parse
	flag.IntVar(&bound, "bound", 1000, "upper bound")

	// other optional parameters are verbose and version
	flag.BoolVar(&version_wanted, "version", false, "shows version info and exists")
}

// shows version information and exit
func showVersion(signal int) {

	fmt.Printf(" %s %s\n", os.Args[0], version)
	os.Exit(signal)
}

func main() {

	// parse the flags
	flag.Parse()

	// if version information was requested show it now and exit
	if version_wanted {
		showVersion(EXIT_SUCCESS)
	}

	// the solution is the sum of all terms of the arithmetic progression with a
	// difference equal to 3 plus the sum of terms of another arithmetic
	// progression with a difference equal to 5. Because multiples of 15 are
	// counted twice, these have to be substracted one considering also the sum
	// of the terms of another arithmetic series

	// compute the upper bounds of all these three series and the number of
	// items in each one
	limit3, limit5, limit15 := 3*((bound-1)/3), 5*((bound-1)/5), 15*((bound-1)/15)
	items3, items5, items15 := 1+limit3/3, 1+limit5/5, 1+limit15/15

	sol3, sol5, sol15 := limit3*items3/2, limit5*items5/2, limit15*items15/2
	fmt.Printf(" The sum of all the multiples of 3 or 5 below %v is %v\n", bound, sol3+sol5-sol15)
}
