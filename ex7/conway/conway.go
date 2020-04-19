// This package provides all means for creating the Conway's Game for arbitrary
// sizes and number of generations
package conway

import (
	"errors"
)

// Phase
// ----------------------------------------------------------------------------

// type

// A phase, consists of a specification of those cells that are alive and those
// that are dead. Admittedly, no more is needed but the following structure
// stores for every phase the width and height only for the convenience of the
// methods that handle it
type phase struct {
	width, height int
	contents      []bool

	// note that each phase also keeps the number of alive cells around each
	// cell. This is done to avoid re-computing information and thus, to make
	// the Conway's Games faster ---Conway deserves it!
	alive []int
}

// methods

// return a new phase which is initially empty
func NewPhase(width, height int) phase {

	// note that in creation, room is allocated for both the contents and the
	// number of alive cells around each cell
	return phase{
		width:    width,
		height:   height,
		contents: make([]bool, width*height),
		alive:    make([]int, width*height)}
}

// return the index of a position (x, y)
func (p *phase) offset(x, y int) int {

	return y*p.width + x
}

// return the number of cells alive around the given position
func (p *phase) nbalive(x, y int) (result int) {

	// if (x, y) is not at the top row
	if y < p.height-1 {
		if p.contents[p.offset(x, y+1)] {
			result += 1
		}

		// if this is not the leftmost column
		if x > 0 {
			if p.contents[p.offset(x-1, y+1)] {
				result += 1
			}
		}

		// if this is not the rightmost column
		if x < p.width-1 {
			if p.contents[p.offset(x+1, y+1)] {
				result += 1
			}
		}
	}

	// if (x, y) is not at the bottom row
	if y > 0 {
		if p.contents[p.offset(x, y-1)] {
			result += 1
		}

		// if this is not the leftmost column
		if x > 0 {
			if p.contents[p.offset(x-1, y-1)] {
				result += 1
			}
		}

		// if this is not the rightmost column
		if x < p.width-1 {
			if p.contents[p.offset(x+1, y-1)] {
				result += 1
			}
		}
	}

	// if (x,y) is not at the leftmost column
	if x > 0 {
		if p.contents[p.offset(x-1, y)] {
			result += 1
		}
	}

	// if (x,y) is not at the rightmost column
	if x < p.width-1 {
		if p.contents[p.offset(x+1, y)] {
			result += 1
		}
	}

	// and return (implicitly) the number of alive cells around (x, y)
	return
}

// Update the value of a cell. In case the cell (x,y) falls out of bounds an
// error is returned and the phase is not modified
func (p *phase) update(x, y int, value bool) error {

	if x < 0 || x >= p.width || y < 0 || y >= p.height {
		return errors.New("Set out of bounds")
	}

	// otherwise, just modify the location
	p.contents[p.offset(x, y)] = value

	// at this point no error happened
	return nil
}

// Return the phase next to this one, i.e., apply the rules of the Conway's Game
func (p *phase) Next() phase {

	// create a new phase with the same dimensions than this one
	next := NewPhase(p.width, p.height)

	// for all cells in this phase
	for x := 0; x < p.width; x++ {
		for y := 0; y < p.height; y++ {

			// get the number of cells alive around cell (x, y)
			alive := p.alive[p.offset(x, y)]

			// by default, the next population is empty, i.e., all of them are
			// dead and thus, the only rules considered are those that make some
			// cells take birth or survive

			// -- survival: Any live cell with two or three live neighbors
			// survives
			if p.contents[p.offset(x, y)] && (alive == 2 || alive == 3) {
				next.update(x, y, true)
			}

			// -- birth: Any dead cell with three live neighbors becomes a live
			// cell
			if !p.contents[p.offset(x, y)] && alive == 3 {
				next.update(x, y, true)
			}
		}
	}

	// before leaving update the number of cells alive around each cell
	// count the number of alive cells around each cell
	for x := 0; x < p.width; x++ {
		for y := 0; y < p.height; y++ {
			next.alive[next.offset(x, y)] = next.nbalive(x, y)
		}
	}

	// and return the next phase
	return next
}

// Set the contents of a phase to those given in contents. In case the given
// slice and the length of the phase contents do not match an error is returned
func (p *phase) Set(contents []bool) error {

	if len(contents) != p.width*p.height {
		return errors.New("Mismatched contents")
	}

	// otherwise, just set the contents of the phase to those given in the slice
	p.contents = contents

	// count the number of alive cells around each cell
	for x := 0; x < p.width; x++ {
		for y := 0; y < p.height; y++ {
			p.alive[p.offset(x, y)] = p.nbalive(x, y)
		}
	}

	// and return no error
	return nil
}

// Phases are stringers so that they can be shown on an output stream
func (p phase) String() (output string) {

	for irow := 0; irow < p.height; irow++ {

		for icol := 0; icol < p.width; icol++ {

			// note the transformation, row #0 is at the bottom so that the last
			// row is shown first. In the following UTF-8 characters are used to
			// somehow beautify the output ---Conway deserves it!
			if p.contents[p.offset(icol, p.height-irow-1)] {
				output += "█"
			} else {
				output += "░"
			}
		}

		// and move to the next line
		output += "\n"
	}

	return output
}

// Conway
// ----------------------------------------------------------------------------

// type

// The Conway's Game consists of a slice with a number of phases each with a
// given width and height
type Conway struct {
	width, height int
	generations   int
	phases        []phase
}

// methods

// Return a new Conway's Game. Note that it is necessary to specify the first
// generation as a phase
func NewConway(width, height, generations int, contents phase) Conway {

	// when creating a new instance of the Conway's Game, note that space is
	// allocated for as many phases as generations, but these are not
	// initialized as a matter of fact
	conway := Conway{
		width:       width,
		height:      height,
		generations: generations,
		phases:      make([]phase, generations)}

	// set the initial contents
	conway.phases[0] = contents

	// and return the new instance
	return conway
}

// Return the phase in the ith generation of this Conway's Game
func (g *Conway) Get(ith int) phase {
	return g.phases[ith]
}

// Run the entire game and generate all generations from the initial population
// in the given instance of the Conway's Game
func (g *Conway) Run() {

	// for all generations but the first one
	for igeneration := 1; igeneration < g.generations; igeneration++ {

		// compute the phase next to the previous one
		g.phases[igeneration] = g.phases[igeneration-1].Next()
	}
}
