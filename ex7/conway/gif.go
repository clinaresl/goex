package conway

import (
	"image"
	"image/color"
	"image/gif"
)

// Phase
// ----------------------------------------------------------------------------

// methods

// returns a palette with precisely two colors to represent dead and alive cells
// respectively, i.e., position 0 is the color used for dead cells and color 1
// is used for living cells
func GeneratePalette() []color.Color {

	// first, insert the white color for dead cells and then the black for
	// living cells
	return []color.Color{color.RGBA{0, 0, 0, 255}, color.RGBA{0, 255, 0, 255}}
}

// returns a paletted image representing the contents of a phase. The palette
// should consist of only two colors representing dead and alive in positions 0
// and 1 respectively
func (p *phase) GetImage(palette []color.Color) *image.Paletted {

	// create an imagen circumscribed to the dimensions of the phase with the
	// specified palette of colors
	rect := image.Rect(0, 0, p.width, p.height)
	img := image.NewPaletted(rect, palette)

	// now, set all pixels of the image
	for y := 0; y < p.height; y++ {
		for x := 0; x < p.width; x++ {

			// set the contents of this pixel. If the cell (x, y) is alive then
			// take the color in position 1, otherwise, take the color in
			// position 0
			var c uint8
			if p.contents[p.offset(x, (p.height-y-1))] {
				c = 1
			}
			img.SetColorIndex(x, y, c)
		}
	}

	// and return now the image
	return img
}

// Conway
// ----------------------------------------------------------------------------

// methods

// returns a gif animation of the Conway's game with the given delay in 100th of
// a second between frames
func (g *Conway) GetGIF(delay int) gif.GIF {

	// create an array of images and delays between successive images
	var delays []int = make([]int, g.generations)
	var images []*image.Paletted = make([]*image.Paletted, g.generations)

	// first, create a simple palete with two colors for drawing dead and living
	// cells
	palette := GeneratePalette()

	// transform each phase of the game into a paletted image
	for index, phase := range g.phases {

		// add the next image after 1 100th of a second
		delays[index] = delay
		images[index] = phase.GetImage(palette)
	}

	// and now return the GIF image
	return gif.GIF{Delay: delays, Image: images}
}
