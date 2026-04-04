package internal

import "fmt"

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

func ParseAxis(str string) (Axis, error) {
	for i, name := range [...]string{"x", "y", "z"} {
		if str == name {
			return Axis(i), nil
		}
	}
	return 0, fmt.Errorf("ParseAxis %s", str)
}
