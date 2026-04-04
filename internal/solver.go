package internal

type Solver interface {
	GetPath(start Cube, end Cube) []Move
}
