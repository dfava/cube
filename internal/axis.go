package internal

type Axis uint

const (
	Xax Axis = iota
	Yax
	Zax
)

func (ax Axis) String() string {
	names := [...]string{"x", "y", "z"}
	return names[ax]
}
