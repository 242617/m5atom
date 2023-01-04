package colors

import "image/color"

var (
	White  = color.RGBA{0xff, 0xff, 0xff, 0xff}
	Black  = color.RGBA{0x00, 0x00, 0x00, 0xff}
	Red    = color.RGBA{0xff, 0x00, 0x00, 0xff}
	Green  = color.RGBA{0x00, 0xff, 0x00, 0xff}
	Blue   = color.RGBA{0x00, 0x00, 0xff, 0xff}
	Yellow = color.RGBA{0xff, 0xff, 0x00, 0xff}
)

func Mix(c1, c2 color.RGBA) color.RGBA {
	if c2.R > c1.R {
		c1.R = c2.R
	}
	if c2.G > c1.G {
		c1.G = c2.G
	}
	if c2.B > c1.B {
		c1.B = c2.B
	}
	if c2.A > c1.A {
		c1.A = c2.A
	}
	return c1
}
