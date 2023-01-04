package motion

import "math"

type ease = func(float64) float64

func EaseLinear(x float64) float64    { return x }
func EaseInOutSine(x float64) float64 { return -(math.Cos(math.Pi*x) - 1) / 2 }
func EaseInOutExpo(x float64) float64 {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	if x < 0.5 {
		return math.Pow(2, 20*x-10) / 2
	}
	return (2 - math.Pow(2, -20*x+10)) / 2
}

func EaseInOutBack(x float64) float64 {
	const c1 = 1.70158
	const c2 = c1 * 1.525

	if x < 0.5 {
		return (math.Pow(2*x, 2) * ((c2+1)*2*x - c2)) / 2
	}
	return (math.Pow(2*x-2, 2)*((c2+1)*(x*2-2)+c2) + 2) / 2
}
