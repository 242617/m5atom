package bitmap

import . "github.com/242617/m5atom/pkg/colors"

var (
	Space = MustNewFromColors(1, 5,
		Black,
		Black,
		Black,
		Black,
		Black,
	)
	A = MustNew5x5(
		Black, White, White, White, Black,
		White, Black, Black, Black, White,
		White, Black, Black, Black, White,
		White, White, White, White, White,
		White, Black, Black, Black, White,
	)
	B = MustNew5x5(
		White, White, White, White, Black,
		White, Black, Black, Black, White,
		White, White, White, White, White,
		White, Black, Black, Black, White,
		White, White, White, White, Black,
	)
	C = MustNew5x5(
		Black, White, White, White, Black,
		White, Black, Black, Black, Black,
		White, Black, Black, Black, Black,
		White, Black, Black, Black, Black,
		Black, White, White, White, Black,
	)
	D = MustNew5x5(
		White, White, White, White, Black,
		White, Black, Black, Black, White,
		White, Black, Black, Black, White,
		White, Black, Black, Black, White,
		White, White, White, White, Black,
	)
	E = MustNew5x5(
		White, White, White, White, White,
		White, Black, Black, Black, Black,
		White, White, White, White, Black,
		White, Black, Black, Black, Black,
		White, White, White, White, White,
	)
	F = MustNew5x5(
		White, White, White, White, White,
		White, Black, Black, Black, Black,
		White, White, White, White, Black,
		White, Black, Black, Black, Black,
		White, Black, Black, Black, Black,
	)
	G = MustNew5x5(
		Black, White, White, White, Black,
		White, Black, Black, Black, Black,
		White, Black, White, White, White,
		White, Black, Black, Black, White,
		Black, White, White, White, Black,
	)
	H = MustNew5x5(
		White, Black, Black, Black, White,
		White, Black, Black, Black, White,
		White, White, White, White, White,
		White, Black, Black, Black, White,
		White, Black, Black, Black, White,
	)
	I = MustNew5x5(
		Black, White, White, White, Black,
		Black, Black, White, Black, Black,
		Black, Black, White, Black, Black,
		Black, Black, White, Black, Black,
		Black, White, White, White, Black,
	)
	J = MustNew5x5(
		Black, Black, White, White, White,
		Black, Black, Black, Black, White,
		Black, White, Black, Black, White,
		Black, White, Black, Black, White,
		Black, Black, White, White, Black,
	)
	K = MustNew5x5(
		White, Black, Black, White, Black,
		White, Black, White, Black, Black,
		White, White, Black, Black, Black,
		White, Black, White, Black, Black,
		White, Black, Black, White, Black,
	)
	L = MustNew5x5(
		White, Black, Black, Black, Black,
		White, Black, Black, Black, Black,
		White, Black, Black, Black, Black,
		White, Black, Black, Black, Black,
		White, White, White, White, White,
	)
	M = MustNew5x5(
		White, Black, Black, Black, White,
		White, White, Black, White, White,
		White, Black, White, Black, White,
		White, Black, Black, Black, White,
		White, Black, Black, Black, White,
	)
	N = MustNew5x5(
		White, Black, Black, Black, White,
		White, White, Black, Black, White,
		White, Black, White, Black, White,
		White, Black, Black, White, White,
		White, Black, Black, Black, White,
	)
	O = MustNew5x5(
		Black, White, White, White, Black,
		White, Black, Black, Black, White,
		White, Black, Black, Black, White,
		White, Black, Black, Black, White,
		Black, White, White, White, Black,
	)
	P = MustNew5x5(
		White, White, White, White, Black,
		White, Black, Black, Black, White,
		White, White, White, White, Black,
		White, Black, Black, Black, Black,
		White, Black, Black, Black, Black,
	)
	Q = MustNew5x5(
		Black, White, White, White, Black,
		White, Black, Black, Black, White,
		White, Black, White, Black, White,
		White, Black, Black, White, Black,
		Black, White, White, Black, White,
	)
	R = MustNew5x5(
		White, White, White, White, Black,
		White, Black, Black, Black, White,
		White, White, White, White, Black,
		White, Black, Black, White, Black,
		White, Black, Black, Black, White,
	)
	S = MustNew5x5(
		Black, White, White, White, White,
		White, Black, Black, Black, Black,
		Black, White, White, White, Black,
		Black, Black, Black, Black, White,
		White, White, White, White, Black,
	)
	T = MustNew5x5(
		White, White, White, White, White,
		Black, Black, White, Black, Black,
		Black, Black, White, Black, Black,
		Black, Black, White, Black, Black,
		Black, Black, White, Black, Black,
	)
	U = MustNew5x5(
		White, Black, Black, Black, White,
		White, Black, Black, Black, White,
		White, Black, Black, Black, White,
		White, Black, Black, Black, White,
		Black, White, White, White, Black,
	)
	V = MustNew5x5(
		White, Black, Black, Black, White,
		White, Black, Black, Black, White,
		Black, White, Black, White, Black,
		Black, White, Black, White, Black,
		Black, Black, White, Black, Black,
	)
	W = MustNew5x5(
		White, Black, Black, Black, White,
		White, Black, Black, Black, White,
		White, Black, White, Black, White,
		White, Black, White, Black, White,
		Black, White, Black, White, Black,
	)
	X = MustNew5x5(
		White, Black, Black, Black, White,
		Black, White, Black, White, Black,
		Black, Black, White, Black, Black,
		Black, White, Black, White, Black,
		White, Black, Black, Black, White,
	)
	Y = MustNew5x5(
		White, Black, Black, Black, White,
		White, Black, Black, Black, White,
		Black, White, Black, White, Black,
		Black, Black, White, Black, Black,
		Black, Black, White, Black, Black,
	)
	Z = MustNew5x5(
		White, White, White, White, White,
		Black, Black, Black, White, Black,
		Black, Black, White, Black, Black,
		Black, White, Black, Black, Black,
		White, White, White, White, White,
	)
	Empty = MustNew5x5(
		Black, Black, Black, Black, Black,
		Black, Black, Black, Black, Black,
		Black, Black, Black, Black, Black,
		Black, Black, Black, Black, Black,
		Black, Black, Black, Black, Black,
	)
)
