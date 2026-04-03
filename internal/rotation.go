// Copyright 2020 Daniel S. Fava. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package internal

type Direction bool

const (
	Clock        = false
	Counterclock = true
)

func (dir Direction) String() string {
	if dir {
		return "counterclockwise"
	}
	return "clockwise"
}

// Returns a 90 degree rotation matrix about an axis,
// either counter-clockwise or clockwise
func GetRotationMatrix(a Axis, counter Direction) matrix {
	var ret matrix
	switch a {
	case Xax:
		if counter {
			ret[0][0] = 1
			ret[1][2] = 1
			ret[2][1] = -1
		} else {
			ret[0][0] = 1
			ret[1][2] = -1
			ret[2][1] = 1
		}
	case Yax:
		if counter {
			ret[0][2] = -1
			ret[1][1] = 1
			ret[2][0] = 1
		} else {
			ret[0][2] = 1
			ret[1][1] = 1
			ret[2][0] = -1
		}
	case Zax:
		if counter {
			ret[0][1] = 1
			ret[1][0] = -1
			ret[2][2] = 1
		} else {
			ret[0][1] = -1
			ret[1][0] = 1
			ret[2][2] = 1
		}
	}
	return ret
}
