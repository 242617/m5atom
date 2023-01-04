package motion

import "time"

func New(object Object, options ...option) *Motion {
	m := Motion{object: object}
	for _, option := range append([]option{withDefaultEase(), withDefaultFPS()}, options...) {
		option(&m)
	}
	return &m
}

type Motion struct {
	object Object
	fps    int
	delay  time.Duration
	ease   ease
}

func (m *Motion) Wait(duration time.Duration) *Motion {
	time.Sleep(duration)
	return m
}

func (m *Motion) To(x, y int, duration time.Duration) *Motion {
	startX, startY := m.object.Position()
	deltaX := x - startX
	deltaY := y - startY
	steps := int(float64(m.fps) * duration.Seconds())

	for i := 0; i < steps; i++ {
		progress := float64(i) / float64(steps)
		ratio := m.ease(progress)
		x := startX + int(ratio*float64(deltaX))
		y := startY + int(ratio*float64(deltaY))
		m.object.MoveTo(x, y)
		time.Sleep(m.delay)
	}
	m.object.MoveTo(x, y)

	return m
}
