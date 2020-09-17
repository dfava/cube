// Copyright 2020 Daniel S. Fava. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cube

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

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

const (
	row = iota
	col
)

const (
	neg = -1
	pos = 1
)

func sign(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}

type projection struct {
	axis   [2]Axis
	offset [2]int
	sign   [2]int
	fun    [2]func(float64) float64
}

// Project down from a side of the cube (specified by an axis and a sign)
// taking the size of the cube (n) into account.
// Returns an offset into a Flat structure and other information
// needed to calculate the projection
func project(n uint, axis Axis, sign bool) projection {
	var proj projection
	if axis == Xax {
		proj.axis[row] = Zax
		proj.axis[col] = Yax
		if sign {
			proj.offset[row] = int(n)
			proj.offset[col] = int(n * 2)
			proj.sign[row] = neg
			proj.sign[col] = neg
		} else {
			proj.offset[row] = int(n)
			proj.offset[col] = 0
			proj.sign[row] = neg
			proj.sign[col] = pos
		}
	} else if axis == Yax {
		proj.axis[row] = Zax
		proj.axis[col] = Xax
		if sign {
			proj.offset[row] = int(n)
			proj.offset[col] = int(n)
			proj.sign[row] = neg
			proj.sign[col] = pos
		} else {
			proj.offset[row] = int(n)
			proj.offset[col] = int(n * 3)
			proj.sign[row] = neg
			proj.sign[col] = neg
		}
	} else {
		proj.axis[row] = Yax
		proj.axis[col] = Xax
		if sign {
			proj.offset[row] = 0
			proj.offset[col] = int(n)
			proj.sign[row] = pos
			proj.sign[col] = pos
		} else {
			proj.offset[row] = int(n * 2)
			proj.offset[col] = int(n)
			proj.sign[row] = neg
			proj.sign[col] = pos
		}
	}

	for _, idx := range [...]int{row, col} {
		if proj.sign[idx] == neg {
			proj.fun[idx] = math.Ceil
		} else {
			proj.fun[idx] = math.Floor
		}
	}
	return proj
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
func (fl *Flat) PaintCubi(cubi Cubi, n uint) {
	var offset float64
	if n%2 == 0 {
		offset = 0.5
	}
	sign := [3]float64{
		float64(-1 * sign(cubi.pv[Xax])),
		float64(-1 * sign(cubi.pv[Yax])),
		float64(-1 * sign(cubi.pv[Zax]))}

	for _, ax := range [...]Axis{Xax, Yax, Zax} {
		if cubi.cv[ax] == 0 {
			continue
		}
		proj := project(n, ax, cubi.cv[ax] > 0)
		r := int(n/2) + proj.sign[row]*int(proj.fun[row](float64(cubi.pv[proj.axis[row]])+sign[proj.axis[row]]*offset)) + proj.offset[row]
		c := int(n/2) + proj.sign[col]*int(proj.fun[col](float64(cubi.pv[proj.axis[col]])+sign[proj.axis[col]]*offset)) + proj.offset[col]
		(*fl)[r][c] = cubi.cv[ax].Abs().String()
	}
}

func (fl *Flat) PaintCube(cube Cube) {
	(*fl) = make([][]string, cube.n*3)
	for idx := range *fl {
		(*fl)[idx] = make([]string, cube.n*4)
	}
	for idx := range cube.cubis {
		fl.PaintCubi(cube.cubis[idx], cube.n)
	}
}

func (fl *Flat) FromFile(fname string) {
	buf, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = buf.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	snl := bufio.NewScanner(buf)
	var iterations uint
	var n uint

	var zeros []string
	for snl.Scan() {
		colors := func(ss []string) []string { // missing Python right now...
			for idx, s := range ss {
				if s == "|" {
					ss = append(ss[:idx], ss[idx+1:]...)
				}
			}
			return ss
		}(strings.Fields(snl.Text()))
		if iterations == 0 {
			n = uint(len(colors))
			(*fl) = make([][]string, n*3)
			zeros = func() []string {
				ret := make([]string, n)
				for i := range ret {
					ret[i] = " "
				}
				return ret
			}()
			//strings.Fields(strings.Repeat("z ", int(n)))
		}
		if iterations < n || iterations >= 2*n {
			colors = append(append(append(zeros, colors...), zeros...), zeros...)
		}
		(*fl)[iterations] = colors
		iterations += 1
	}
	err = snl.Err()
	if err != nil {
		log.Fatal(err)
	}
	if iterations != n*3 {
		log.Fatal(errors.New("invalid cube size"))
	}
}

func (fl Flat) ToCube() Cube {
	n := len(fl) / 3

	var cube Cube
	cube.n = uint(n)
	cube.cubis = make([]Cubi, int(math.Pow(float64(n), 3)-math.Pow(float64(n)-2, 3)))

	extremity := n / 2

	preCube := make(map[Vec]CVec)
	for r := 0; r < n*3; r++ {
		c := 0
		for ; c < n*4; c++ {
			if fl[r][c] == " " || fl[r][c] == "" {
				continue
			}
			//fmt.Println(r, c, fl[r][c])
			var x, y, z int
			var xc, yc, zc Color
			if r < n {
				// Size 5 (yellow)
				//fmt.Println("yellow")
				proj := project(uint(n), Zax, true)
				x = (c - proj.offset[col] - n/2) * proj.sign[col]
				y = (r - proj.offset[row] - n/2) * proj.sign[row]
				z = extremity
				zc = ToColor(fl[r][c])
			} else if r >= 2*n {
				// Size w (white)
				//fmt.Println("white")
				proj := project(uint(n), Zax, false)
				x = (c - proj.offset[col] - n/2) * proj.sign[col]
				y = (r - proj.offset[row] - n/2) * proj.sign[row]
				z = -extremity
				zc = -ToColor(fl[r][c])
			} else {
				// Could be one of sides 4 (red), 1 (green), 3 (orange), or 6 (blue)
				if c < n {
					// Size 4 (red)
					//fmt.Println("red")
					proj := project(uint(n), Xax, false)
					x = -extremity
					y = (c - proj.offset[col] - n/2) * proj.sign[col]
					z = (r - proj.offset[row] - n/2) * proj.sign[row]
					xc = -ToColor(fl[r][c])
				} else if c < 2*n {
					// Size 1 (green)
					//fmt.Println("green")
					proj := project(uint(n), Yax, true)
					x = (c - proj.offset[col] - n/2) * proj.sign[col]
					y = extremity
					z = (r - proj.offset[row] - n/2) * proj.sign[row]
					yc = ToColor(fl[r][c])
				} else if c < 3*n {
					// Size 3 (orange)
					//fmt.Println("orange")
					proj := project(uint(n), Xax, true)
					x = extremity
					y = (c - proj.offset[col] - n/2) * proj.sign[col]
					z = (r - proj.offset[row] - n/2) * proj.sign[row]
					xc = ToColor(fl[r][c])
				} else {
					// Size 6 (blue)
					//fmt.Println("blue")
					proj := project(uint(n), Yax, false)
					x = (c - proj.offset[col] - n/2) * proj.sign[col]
					y = -extremity
					z = (r - proj.offset[row] - n/2) * proj.sign[row]
					yc = -ToColor(fl[r][c])
				}
			}
			//fmt.Println(x, y, z, xc, yc, zc)
			cvec, ok := preCube[Vec{x, y, z}]
			if !ok {
				preCube[Vec{x, y, z}] = CVec{xc, yc, zc}
			} else {
				preCube[Vec{x, y, z}] = cvec.or(CVec{xc, yc, zc})
			}
		}
		ncubi := 0
		for pvec, cvec := range preCube {
			//fmt.Println(pvec, cvec)
			cube.cubis[ncubi] = Cubi{cv: cvec, pv: pvec}
			ncubi += 1
		}
	}

	return cube
}
