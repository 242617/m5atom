package motion

type Object interface {
	Position() (int, int)
	MoveTo(x, y int)
}
