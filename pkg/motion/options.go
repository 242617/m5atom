package motion

import "time"

type option = func(*Motion)

func withDefaultEase() option   { return WithEase(EaseLinear) }
func WithEase(ease ease) option { return func(m *Motion) { m.ease = ease } }

func withDefaultFPS() option { return WithFPS(24) }
func WithFPS(fps int) option {
	return func(m *Motion) {
		m.fps = fps
		m.delay = time.Second / time.Duration(m.fps)
	}
}
