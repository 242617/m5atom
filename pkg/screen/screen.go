package screen

import (
	"image/color"

	"github.com/pkg/errors"

	"github.com/242617/m5atom/pkg/colors"
)

func New(options ...option) (*Screen, error) {
	var s Screen

	for _, option := range append([]option{withDefaultBrightness()}, options...) {
		if err := option(&s); err != nil {
			return nil, errors.Wrap(err, "apply option")
		}
	}

	if s.w == 0 {
		return nil, errors.New("empty width")
	}
	if s.h == 0 {
		return nil, errors.New("empty height")
	}

	s.pixels = make([][]color.RGBA, s.h)
	for i := range s.pixels {
		s.pixels[i] = make([]color.RGBA, s.w)
	}
	s.colors = make([]color.RGBA, s.w*s.h)

	return &s, nil
}

type Screen struct {
	pixels     [][]color.RGBA
	colors     []color.RGBA
	objects    []Object
	brightness float64
	w, h       int
}

func (s *Screen) Render() []color.RGBA {
	for i := range s.colors {
		s.colors[i] = color.RGBA{}
	}

	for _, object := range s.objects {

		var n int
		for y := 0; y < s.h; y++ {
			for x := 0; x < s.h; x++ {
				c := object.GetColorAt(x, y)
				c = setBrightness(c, s.brightness)
				s.colors[n] = colors.Mix(s.colors[n], c)
				n++
			}
		}

	}

	return s.colors
}

func (s *Screen) Add(object Object) {
	s.objects = append(s.objects, object)
}
func (s *Screen) Remove(object Object) {
	for i := 0; i < len(s.objects); i++ {
		if s.objects[i] == object {
			s.objects = append(s.objects[:i], s.objects[i+1:]...)
		}
	}
}
func (s *Screen) Objects() []Object { return s.objects }

func setBrightness(c color.RGBA, brightness float64) color.RGBA {
	return color.RGBA{
		R: uint8(float64(c.R) * brightness),
		G: uint8(float64(c.G) * brightness),
		B: uint8(float64(c.B) * brightness),
		A: uint8(float64(c.A) * brightness),
	}
}
