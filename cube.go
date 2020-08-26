// Copyright 2020 Daniel S. Fava. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cube

import (
	"fmt"
	"math"
	"math/rand"
)

type Color int

var printInColors bool

const (
	zero   Color = iota // the "zero color"
	green               // green
	white               // white
	orange              // orange
	red                 // red
	yellow              // yellow
	blue                // blue
)

func (c Color) Abs() Color {
	if c < 0 {
		return -c
	}
	return c
}

func Sign(x int) int {
  if x < 0 {
    return -1
  }
  return 1
}

// using ANSI escape codes for colors
func (c Color) String() string {
	var names [7]string
	if printInColors {
		names = [...]string{" ", "\033[32mg\033[0m", "\033[37mw\033[0m", "\033[35mo\033[0m", "\033[31mr\033[0m", "\033[33my\033[0m", "\033[34mb\033[0m"} // no orange, using magenta instead
	} else {
		names = [...]string{" ", "g", "w", "o", "r", "y", "b"}
	}
	var str string
	if c < 0 {
		str = "-"
		c = -c
	}
	return str + names[c]
}

// Flat is a data-structure that makes it easy to represent a cube
// as a string (for example, when printing it to the screen).
//
// Flat is a representation in which the cube is "opened up" and
// flattened into two dimensions.
//
// A cube of size n^3 is flattened into a 2D matrix of size (3*n)x(4*n).
//
// For example, the Rubik's cube, which is of size n=3^3, is flattened
// into a 9x12 matrix.  Below we show a picture of the flattened Rubik's
// cube.  The key for reading the picture below is:
//
//   u for `up`
//   l for `left`
//   c for `center`
//   r for `right`
//   b for `back`
//   d for `down`
//   . is unused (it does not represent any part of the 3D cube)
//
// . . . u u u . . . . . .
// . . . u u u . . . . . .
// . . . u u u . . . . . .
// l l l c c c r r r b b b
// l l l c c c r r r b b b
// l l l c c c r r r b b b
// . . . d d d . . . . . .
// . . . d d d . . . . . .
// . . . d d d . . . . . .

type Flat [][]string

func (fl Flat) String() string {
	n := len(fl) / 3
	str := ""
	for r := 0; r < n*3; r++ {
		c := 0
		for ; c < n*4; c++ {
			if c != 0 && c != n*4 && c%n == 0 {
				if r > (n-1) && r < 2*n {
					str += fmt.Sprintf("| ")
				} else {
					str += fmt.Sprintf("  ")
				}
			}
			if fl[r][c] == "" {
				str += fmt.Sprintf("%s ", zero.String())
			} else {
				str += fmt.Sprintf("%s ", fl[r][c])
			}
		}
		if r != (n*3-1) && c == n*4 {
			str += "\n"
		}
	}
	return str
}

// Populate a Flat structure given a Cubi.
// 
// A Cubi is composed of a location in space (captured by a Vec) and
// of a description on how to paint that location (captured by a CVec,
// aka color vector).
//
// We use the Vec and CVec to find the indices in Flat that need to be
// populated, and we use the CVec to determine the string representation
// the location's color.
func (fl *Flat) PaintCubi(cbi Cubi, n uint) {
	var offset float64
	if n%2 == 0 {
		offset = 0.5
	}
	signX := float64(-1 * Sign(cbi.pv[Xax]))
	signY := float64(-1 * Sign(cbi.pv[Yax]))
	signZ := float64(-1 * Sign(cbi.pv[Zax]))
	if cbi.cv[Xax] > 0 {
		r := int(n/2) - int(math.Ceil(float64(cbi.pv[Zax])+signZ*offset))
		c := int(n/2) - int(math.Ceil(float64(cbi.pv[Yax])+signY*offset))
		off_r := int(n)
		off_c := int(n * 2)
		(*fl)[r+off_r][c+off_c] = cbi.cv[Xax].Abs().String()
	} else if cbi.cv[Xax] < 0 {
		r := int(n/2) - int(math.Ceil(float64(cbi.pv[Zax])+signZ*offset))
		c := int(n/2) + int(math.Floor(float64(cbi.pv[Yax])+signY*offset))
		off_r := int(n)
		off_c := 0
		(*fl)[r+off_r][c+off_c] = cbi.cv[Xax].Abs().String()
	}
	if cbi.cv[Yax] > 0 {
		r := int(n/2) - int(math.Ceil(float64(cbi.pv[Zax])+signZ*offset))
		c := int(n/2) + int(math.Floor(float64(cbi.pv[Xax])+signX*offset))
		off_r := int(n)
		off_c := int(n)
		(*fl)[r+off_r][c+off_c] = cbi.cv[1].Abs().String()
	} else if cbi.cv[Yax] < 0 {
		r := int(n/2) - int(math.Ceil(float64(cbi.pv[Zax])+signZ*offset))
		c := int(n/2) - int(math.Ceil(float64(cbi.pv[Xax])+signX*offset))
		off_r := int(n)
		off_c := int(n * 3)
		(*fl)[r+off_r][c+off_c] = cbi.cv[1].Abs().String()
	}
	if cbi.cv[Zax] > 0 {
		r := int(n/2) + int(math.Floor(float64(cbi.pv[Yax])+signY*offset))
		c := int(n/2) + int(math.Floor(float64(cbi.pv[Xax])+signX*offset))
		off_r := 0
		off_c := int(n)
		(*fl)[r+off_r][c+off_c] = cbi.cv[2].Abs().String()
	} else if cbi.cv[Zax] < 0 {
		r := int(n/2) - int(math.Ceil(float64(cbi.pv[Yax])+signY*offset))
		c := int(n/2) + int(math.Floor(float64(cbi.pv[Xax])+signX*offset))
		off_r := int(n * 2)
		off_c := int(n)
		(*fl)[r+off_r][c+off_c] = cbi.cv[2].Abs().String()
	}
}

func (fl *Flat) PaintCube(cb Cube) {
	(*fl) = make([][]string, cb.n*3)
	for idx := range *fl {
		(*fl)[idx] = make([]string, cb.n*4)
	}
	for idx := range cb.cubis {
		fl.PaintCubi(cb.cubis[idx], cb.n)
	}
}

type Vec [3]int       // a vector in 3D
type CVec [3]Color    // a "color vector
type Matrix [3][3]int // a matrix

// Multiplying a matrix by a Cubi
// 
// The function multiplies the matrix by the position vector and
// it multiplies the matrix by the color vector
func (m *Matrix) Mult(cbi Cubi) Cubi {
	var ret Cubi
	for i := 0; i <= 2; i++ {
		ret.pv[i] = m[i][0]*cbi.pv[0] + m[i][1]*cbi.pv[1] + m[i][2]*cbi.pv[2]
		ret.cv[i] = Color(m[i][0]*int(cbi.cv[0]) + m[i][1]*int(cbi.cv[1]) + m[i][2]*int(cbi.cv[2]))
	}
	return ret
}

type Cubi struct {
	pv Vec  // position vector
	cv CVec // color vector
}

type Cube struct {
	n     uint
	cubis []Cubi
}

// A string representation of a Rubik's cube
func (cb Cube) String() string {
	var fl Flat
	fl.PaintCube(cb)
	return fl.String()
}

type Axis uint

const (
	Xax Axis = iota
	Yax
	Zax
)

func (ax Axis) String() string {
	names := [...]string{"Xax", "Yax", "Zax"}
	return names[ax]
}

// Returns a 90 degree rotation matrix about an axis,
// either counter-clockwise or clockwise
func GetRotationMatrix(a Axis, counter bool) Matrix {
	var ret Matrix
	if a == Xax {
		if counter {
			ret[0][0] = 1
			ret[1][2] = 1
			ret[2][1] = -1
		} else {
			ret[0][0] = 1
			ret[1][2] = -1
			ret[2][1] = 1
		}
	} else if a == Yax {
		if counter {
			ret[0][2] = -1
			ret[1][1] = 1
			ret[2][0] = 1
		} else {
			ret[0][2] = 1
			ret[1][1] = 1
			ret[2][0] = -1
		}
	} else if a == Zax {
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

// Rotates a cube by multiplying the relevant little cubes by a rotation matrix
func (cb Cube) Rotate(a Axis, idx int, counter bool) Cube {
	ret := cb.New()
	m := GetRotationMatrix(a, counter)
	for cb_idx := range cb.cubis {
		if cb.cubis[cb_idx].pv[a] == idx {
			// We rotate via matrix multiplication
			ret.cubis[cb_idx] = m.Mult(cb.cubis[cb_idx])
		}
	}
	return ret
}

// Returns a cube of size n^3 in its initial configuration.
// Handles cubes of n>2.  Cannot handle the trivial 1x1 cube
//
// Initializes the cube to:
// green  at the center   (y>0)
// blue   to the back     (y<0)
// red    to the left     (x<0)
// orange to the right    (x>0)
// yellow upward          (z>0)
// white  downward        (z<0)
func (cb *Cube) Init(n uint) {
	(*cb).n = n
	(*cb).cubis = make([]Cubi, int(math.Pow(float64(n), 3)-math.Pow(float64(n)-2, 3)))

	ncubi := 0
	for x := -int(n) / 2; x <= int(n)/2; x++ {
		if x == 0 && n%2 == 0 {
			continue
		}
		for y := -int(n) / 2; y <= int(n)/2; y++ {
			if y == 0 && n%2 == 0 {
				continue
			}
			for z := -int(n) / 2; z <= int(n)/2; z++ {
				if z == 0 && n%2 == 0 {
					continue
				}

				// Determine whether it's a center cubi (not an exterior-facing cubi)
				if x < int(n)/2 && x > -int(n)/2 &&
					y < int(n)/2 && y > -int(n)/2 &&
					z < int(n)/2 && z > -int(n)/2 {
					continue
				}

				var xc, yc, zc Color
				if x > 0 {
					xc = orange
				} else if x < 0 {
					xc = -red
				}
				if y > 0 {
					yc = green
				} else if y < 0 {
					yc = -blue
				}
				if z > 0 {
					zc = yellow
				} else if z < 0 {
					zc = -white
				}
				cb.cubis[ncubi] = Cubi{cv: CVec{xc, yc, zc}, pv: Vec{x, y, z}}
				ncubi += 1
			}
		}
	}
}

func (cb *Cube) Shuffle(times uint) {
	axes := [...]Axis{Xax, Yax, Zax}
	idxs := make([]int, cb.n+(cb.n+1)%2)
	for idx := -int(cb.n) / 2; idx <= int(cb.n)/2; idx++ {
		idxs[idx+int(cb.n)/2] = idx
	}
	dirs := [...]bool{true, false}
	var perms uint
	for perms < times {
		ax := axes[rand.Intn(len(axes))]  // pick an axis
		idx := idxs[rand.Intn(len(idxs))] // pick an index
		dir := dirs[rand.Intn(len(dirs))] // pick a direction
		if cb.n%2 == 0 && idx == 0 {
			continue
		}
		(*cb) = cb.Rotate(ax, idx, dir)
		perms += 1
	}
}

// A copy-constructor
func (cb Cube) New() Cube {
	var ret Cube
	ret.n = cb.n
	ret.cubis = make([]Cubi, len(cb.cubis))
	for idx := range cb.cubis {
		ret.cubis[idx] = cb.cubis[idx]
	}
	return ret
}

func init() {
	printInColors = true
	rand.Seed(42) // Pseudo random
}

func PrintInColors(flag bool) {
	printInColors = flag
}
