package screen

import "image/color"

type Object interface {
	GetColorAt(x, y int) color.RGBA
}
