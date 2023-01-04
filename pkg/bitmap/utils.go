package bitmap

import (
	"image/color"

	"github.com/pkg/errors"
)

func MustNew5x5(colors ...color.RGBA) *Bitmap { return MustNewFromColors(5, 5, colors...) }
func MustNewFromColors(w, h int, colors ...color.RGBA) *Bitmap {
	b, err := NewFromColors(w, h, colors...)
	if err != nil {
		panic(errors.Wrap(err, "must new from colors"))
	}
	return b
}

func NewFromColors(w, h int, colors ...color.RGBA) (*Bitmap, error) {
	if w*h != len(colors) {
		return nil, errors.New("unexpected colors length")
	}

	pixels := make([][]color.RGBA, h)
	for i := range pixels {
		pixels[i] = make([]color.RGBA, w)
	}

	var n int
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			pixels[j][i] = colors[n]
			n++
		}
	}

	return New(pixels)
}

func MustComposeHorizontal(bitmaps ...*Bitmap) *Bitmap {
	b, err := ComposeHorizontal(bitmaps...)
	if err != nil {
		panic(errors.Wrap(err, "must compose horizontal"))
	}
	return b
}
func ComposeHorizontal(bitmaps ...*Bitmap) (*Bitmap, error) {
	if len(bitmaps) == 0 {
		return nil, errors.New("empty bitmaps")
	}

	var w int
	for i := range bitmaps {
		w += int(bitmaps[i].w)
	}
	h := 5

	pixels := make([][]color.RGBA, h)
	for i := range pixels {
		pixels[i] = make([]color.RGBA, w)
	}

	var shift int
	for _, bitmap := range bitmaps {
		for j := 0; j < len(bitmap.pixels); j++ {
			for i := 0; i < len(bitmap.pixels[j]); i++ {
				pixels[j][shift+i] = bitmap.pixels[j][i]
			}
		}
		shift += int(bitmap.w)
	}

	return New(pixels)
}
