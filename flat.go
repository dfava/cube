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
	signX := float64(-1 * Sign(cubi.pv[Xax]))
	signY := float64(-1 * Sign(cubi.pv[Yax]))
	signZ := float64(-1 * Sign(cubi.pv[Zax]))
	if cubi.cv[Xax] > 0 {
		r := int(n/2) - int(math.Ceil(float64(cubi.pv[Zax])+signZ*offset))
		c := int(n/2) - int(math.Ceil(float64(cubi.pv[Yax])+signY*offset))
		off_r := int(n)
		off_c := int(n * 2)
		(*fl)[r+off_r][c+off_c] = cubi.cv[Xax].Abs().String()
	} else if cubi.cv[Xax] < 0 {
		r := int(n/2) - int(math.Ceil(float64(cubi.pv[Zax])+signZ*offset))
		c := int(n/2) + int(math.Floor(float64(cubi.pv[Yax])+signY*offset))
		off_r := int(n)
		off_c := 0
		(*fl)[r+off_r][c+off_c] = cubi.cv[Xax].Abs().String()
	}
	if cubi.cv[Yax] > 0 {
		r := int(n/2) - int(math.Ceil(float64(cubi.pv[Zax])+signZ*offset))
		c := int(n/2) + int(math.Floor(float64(cubi.pv[Xax])+signX*offset))
		off_r := int(n)
		off_c := int(n)
		(*fl)[r+off_r][c+off_c] = cubi.cv[1].Abs().String()
	} else if cubi.cv[Yax] < 0 {
		r := int(n/2) - int(math.Ceil(float64(cubi.pv[Zax])+signZ*offset))
		c := int(n/2) - int(math.Ceil(float64(cubi.pv[Xax])+signX*offset))
		off_r := int(n)
		off_c := int(n * 3)
		(*fl)[r+off_r][c+off_c] = cubi.cv[1].Abs().String()
	}
	if cubi.cv[Zax] > 0 {
		r := int(n/2) + int(math.Floor(float64(cubi.pv[Yax])+signY*offset))
		c := int(n/2) + int(math.Floor(float64(cubi.pv[Xax])+signX*offset))
		off_r := 0
		off_c := int(n)
		(*fl)[r+off_r][c+off_c] = cubi.cv[2].Abs().String()
	} else if cubi.cv[Zax] < 0 {
		r := int(n/2) - int(math.Ceil(float64(cubi.pv[Yax])+signY*offset))
		c := int(n/2) + int(math.Floor(float64(cubi.pv[Xax])+signX*offset))
		off_r := int(n * 2)
		off_c := int(n)
		(*fl)[r+off_r][c+off_c] = cubi.cv[2].Abs().String()
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
