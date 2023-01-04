package screen

type option = func(*Screen) error

func WithDimensions(w, h int) option {
	return func(s *Screen) error {
		s.w, s.h = w, h
		return nil
	}
}

func withDefaultBrightness() option { return WithBrightness(1) }
func WithBrightness(brightness float64) option {
	return func(s *Screen) error {
		s.brightness = brightness
		return nil
	}
}

func WithObjects(objects ...Object) option {
	return func(s *Screen) error {
		s.objects = append(s.objects, objects...)
		return nil
	}
}
