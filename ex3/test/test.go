package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/clinaresl/goex/ex3/register"
)

// globals
// ----------------------------------------------------------------------------
const EXIT_SUCCESS = 0
const EXIT_FAILURE = 1

const version = "0.1"

var port int = -1
var want_version bool

// functions
// ----------------------------------------------------------------------------

// init module
//
// setup the flag environment for the on-line help
func init() {

	// first, create a command-line argument for parsing the number
	flag.IntVar(&port, "port", -1, "port used by the server to attend requests")

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

	// secondly, verify whether a port was given or not. Note the trick: the
	// flag package does not allow the definition of mandatory arguments, thus
	// the idea consists of giving mandatory arguments sensible default values
	// that could be verified in run time. This is not the problem this time as
	// -1 is not legal port
	if port < 0 {

		log.Fatalf(" Use -port to provide a port")
		os.Exit(EXIT_FAILURE)
	}

	// if a port has been given, just setup the services provided by our tiny
	// server and attend all requests
	register.Serve(port)
}
