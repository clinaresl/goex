// Conway's Game
package main

import (
	"flag"
	"fmt"
	"image/gif"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/clinaresl/goex/ex7/conway"
)

// globals
// ----------------------------------------------------------------------------
const EXIT_SUCCESS = 0
const EXIT_FAILURE = 1

const version = "0.1"

// flag parameters
var (
	filename      string
	width, height int
	population    int
	nbgenerations int
	want_version  bool
)

// functions
// ----------------------------------------------------------------------------

// init module
//
// setup the flag environment for the on-line help
func init() {

	// command line arguments for parsing the name of the gif file
	flag.StringVar(&filename, "filename", "conway.gif", "name of the GIF file")

	// command line arguments for parsing the dimensions of the grid
	flag.IntVar(&width, "width", 100, "Width of the grid")
	flag.IntVar(&height, "height", 100, "Height of the grid")

	// command line argument to determine the initial number of alive cells
	flag.IntVar(&population, "population", 100, "initial population")

	// command line argument for getting the desired number of generations
	flag.IntVar(&nbgenerations, "generations", 100, "number of generations")

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

	// initialize the first generation randomly
	contents := make([]bool, width*height)
	for i := 0; i < population; i++ {
		contents[i] = true
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(width*height, func(i, j int) {
		contents[i], contents[j] = contents[j], contents[i]
	})

	// and set it
	phase := conway.NewPhase(width, height)
	phase.Set(contents)

	// Create a Conway's Game with this phase
	game := conway.NewConway(width, height, nbgenerations, phase)

	// and run the Conway's Game over this initial generation
	game.Run()

	// get the image of the entire Conway's game
	anim := game.GetGIF()

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	gif.EncodeAll(f, &anim)
}
