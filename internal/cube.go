// Copyright 2020 Daniel S. Fava. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package internal

import (
	"math"
	"math/rand"
)

type vec [3]int       // a vector in 3D
type cVec [3]Color    // a "color vector
type matrix [3][3]int // a matrix

func (a cVec) or(b cVec) cVec {
	var c cVec
	or := func(a Color, b Color) Color {
		if a == zero {
			return b
		}
		return a
	}
	c[0] = or(a[0], b[0])
	c[1] = or(a[1], b[1])
	c[2] = or(a[2], b[2])
	return c
}

// Multiplying a matrix by a cubi
//
// The function multiplies the matrix by the position vector and
// it multiplies the matrix by the color vector
func (m *matrix) mult(cbi cubi) cubi {
	var ret cubi
	for i := 0; i <= 2; i++ {
		ret.pv[i] = m[i][0]*cbi.pv[0] + m[i][1]*cbi.pv[1] + m[i][2]*cbi.pv[2]
		ret.cv[i] = Color(m[i][0]*int(cbi.cv[0]) + m[i][1]*int(cbi.cv[1]) + m[i][2]*int(cbi.cv[2]))
	}
	return ret
}

type cubi struct {
	pv vec  // position vector
	cv cVec // color vector
}

type Cube struct {
	n     uint
	cubis []cubi
}

func New(n uint) Cube {
	if n <= 1 {
		panic("n must be greater than 1")
	}
	ret := Cube{n: n}
	ret.cubis = make([]cubi, int(math.Pow(float64(n), 3)-math.Pow(float64(n)-2, 3)))
	ret.Reset()
	return ret
}

// A copy-constructor
func (cube Cube) Copy() Cube {
	var ret Cube
	ret.n = cube.n
	ret.cubis = make([]cubi, len(cube.cubis))
	for idx := range cube.cubis {
		ret.cubis[idx] = cube.cubis[idx]
	}
	return ret
}

// Resets the cube to:
// green  at the center   (y>0)
// blue   to the back     (y<0)
// red    to the left     (x<0)
// orange to the right    (x>0)
// yellow upward          (z>0)
// white  downward        (z<0)
//
// Cannot handle the trivial 1x1 cube
func (cube *Cube) Reset() {
	n := cube.n
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

				// Give xc a color if x is at the Cubi has a face on
				// either side of the x axis.  Same for yc, zc and the y, z axes.
				extremity := int(n / 2)
				var xc, yc, zc Color
				if x == extremity {
					xc = orange
				} else if x == -extremity {
					xc = -red
				}
				if y == extremity {
					yc = green
				} else if y == -extremity {
					yc = -blue
				}
				if z == extremity {
					zc = yellow
				} else if z == -extremity {
					zc = -white
				}
				cube.cubis[ncubi] = cubi{cv: cVec{xc, yc, zc}, pv: vec{x, y, z}}
				ncubi += 1
			}
		}
	}
}

// A string representation of a Rubik's cube
func (cube Cube) String() string {
	var fl Flat
	fl.PaintCube(cube)
	return fl.String()
}

func (cube Cube) GetSize() uint {
	return cube.n
}

// Performs a move on a cube by turning part of the
// cube about an Axis in a particular direction
func (cube Cube) Turn(a Axis, idx int, counter Direction) Cube {
	ret := cube.Copy()
	m := GetRotationMatrix(a, counter)
	for cube_idx := range cube.cubis {
		if cube.cubis[cube_idx].pv[a] == idx {
			// We rotate via matrix multiplication
			ret.cubis[cube_idx] = m.mult(cube.cubis[cube_idx])
		}
	}
	return ret
}

// Perform all moves on a cube
func (cube Cube) GetAllTurns() []Cube {
	var ret []Cube
	for _, ax := range [...]Axis{Xax, Yax, Zax} {
		for idx := -int(cube.n) / 2; idx <= int(cube.n)/2; idx++ {
			for _, dir := range [...]Direction{Counterclock, Clock} {
				if cube.n%2 == 0 && idx == 0 {
					continue
				}
				newCube := cube.Turn(ax, idx, dir)
				if cube.n%2 == 1 && idx == 0 {
					// Preserves a canonical cube
					newCube = newCube.Rotate(ax, !dir)
				}
				ret = append(ret, newCube)
			}
		}
	}
	return ret
}

// Rotate the whole cube in 3 dimensions
func (cube Cube) Rotate(a Axis, counter Direction) Cube {
	ret := cube.Copy()
	m := GetRotationMatrix(a, counter)
	for cube_idx := range cube.cubis {
		// We rotate via matrix multiplication
		ret.cubis[cube_idx] = m.mult(cube.cubis[cube_idx])
	}
	return ret
}

// Perform all rotations on a cube
// There are 24 equivalence classes (rotations) of a cube.
func (cube Cube) GetAllRotations() []Cube {
	var ret []Cube

	current := cube
	explored := make(map[string]Cube)
	var toExplore []Cube

	for {
		explored[current.String()] = current

		for _, ax := range [...]Axis{Xax, Yax, Zax} {
			for _, dir := range [...]Direction{Counterclock, Clock} {
				other := current.Rotate(ax, dir)
				if _, in := explored[other.String()]; !in {
					toExplore = append(toExplore, other)
				}
			}
		}

		if len(toExplore) == 0 {
			// Nothing more to explore
			for k := range explored {
				ret = append(ret, explored[k])
			}
			return ret
		}

		current = toExplore[0]
		toExplore = toExplore[1:]
	}
}

func (cube *Cube) Shuffle(times uint) {
	axes := [...]Axis{Xax, Yax, Zax}
	idxs := make([]int, cube.n+(cube.n+1)%2)
	for idx := -int(cube.n) / 2; idx <= int(cube.n)/2; idx++ {
		idxs[idx+int(cube.n)/2] = idx
	}
	dirs := [...]Direction{Counterclock, Clock}
	var perms uint
	for perms < times {
		ax := axes[rand.Intn(len(axes))]  // pick an axis
		idx := idxs[rand.Intn(len(idxs))] // pick an index
		dir := dirs[rand.Intn(len(dirs))] // pick a direction
		if cube.n%2 == 0 && idx == 0 {
			continue
		}
		(*cube) = cube.Turn(ax, idx, dir)
		if idx == 0 { // Preserve the cube's orientation
			(*cube) = cube.Rotate(ax, !dir)
		}
		perms += 1
	}
}

func (cube Cube) IsSolved() bool {
	// We could do with [3][2]Color, I'm wasting a bit of memory to simplify
	// the algorithm:
	// Signs -1 and 1 map to array indexes 0 and 2 as opposed to 0 and 1
	var sideColor [3][3]Color
	extremity := int((cube.n) / 2)
	for idx := range cube.cubis {
		for _, ax := range [...]Axis{Xax, Yax, Zax} {
			for sign := range []int{-1, 1} {
				if cube.cubis[idx].pv[ax] == sign*extremity {
					if sideColor[ax][sign+1] == zero {
						sideColor[ax][sign+1] = cube.cubis[idx].cv[ax]
					} else if sideColor[ax][sign+1] != cube.cubis[idx].cv[ax] {
						return false
					}
				}
			}
		}
	}
	return true
}

func (cube Cube) IsCanonical() bool {
	if cube.n%2 == 0 { // Only odd sized cubes can be canonical
		return false
	}
	extremity := int(cube.n / 2)
	canon := true
	for _, cbi := range cube.cubis {
		switch cbi.pv {
		case vec{-extremity, 0, 0}:
			canon = canon && (cbi.cv[Xax] == -red)
		case vec{extremity, 0, 0}:
			canon = canon && (cbi.cv[Xax] == orange)
		case vec{0, -extremity, 0}:
			canon = canon && (cbi.cv[Yax] == -blue)
		case vec{0, extremity, 0}:
			canon = canon && (cbi.cv[Yax] == green)
		case vec{0, 0, -extremity}:
			canon = canon && (cbi.cv[Zax] == -white)
		case vec{0, 0, extremity}:
			canon = canon && (cbi.cv[Zax] == yellow)
		}
	}
	return canon
}
