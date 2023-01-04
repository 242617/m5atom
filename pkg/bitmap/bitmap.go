package bitmap

import (
	"errors"
	"image/color"
)

func New(pixels [][]color.RGBA) (*Bitmap, error) {
	if len(pixels) == 0 {
		return nil, errors.New("empty pixels")
	}

	var h = len(pixels)
	var w = len(pixels[0])

	for i := range pixels {
		if len(pixels[i]) == 0 || len(pixels[i]) != w {
			return nil, errors.New("unexpected pixels length")
		}
	}

	return &Bitmap{
		pixels: pixels,
		w:      uint(w),
		h:      uint(h),
	}, nil
}

type Bitmap struct {
	pixels [][]color.RGBA
	x, y   int
	w, h   uint
}

func (b *Bitmap) GetColorAt(x, y int) color.RGBA {
	dx, dy := x-b.x, y-b.y
	if x < b.x || dy < 0 {
		return color.RGBA{}
	}
	if len(b.pixels)-1 < dy || len(b.pixels[dy])-1 < dx {
		return color.RGBA{}
	}
	return b.pixels[dy][dx]
}

func (b *Bitmap) MoveTo(x, y int)         { b.x, b.y = x, y }
func (b *Bitmap) Position() (x, y int)    { return b.x, b.y }
func (b *Bitmap) Dimensions() (w, h uint) { return b.w, b.h }
