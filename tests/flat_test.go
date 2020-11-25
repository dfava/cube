// Copyright 2020 Daniel S. Fava. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cube_test

import (
	"fmt"
	. "github.com/dfava/cube"
	"testing"
)

func TestFromFile2(t *testing.T) {
	var fl Flat
	fl.FromFile("cube2.txt")
	fmt.Println(fl)
	cube := fl.Cube()
	fmt.Println()
	fmt.Println(cube)
}

func TestFromFile3(t *testing.T) {
	var fl Flat
	fl.FromFile("cube3.txt")
	fmt.Println(fl)
	cube := fl.Cube()
	fmt.Println()
	fmt.Println(cube)
}

func cube2Flat2Cube(t *testing.T, sizes []uint) {
	for _, n := range sizes {
		for _, shuffle := range [...]uint{0, 1, 2, 10, 13} {
			var cube Cube
			cube.Init(n)
			cube.Shuffle(shuffle)
			var fl Flat
			fl.PaintCube(cube)
			fmt.Println(fl)
			var other = fl.Cube()
			fmt.Println(other)
			if cube.String() != other.String() {
				t.Errorf("flattening and reconstructing failed! n=%d", n)
			}
		}
	}
}

func TestCube2Flat2CubeOdd(t *testing.T) {
	cube2Flat2Cube(t, []uint{3, 5, 7, 9})
}

func TestCube2Flat2CubeEven(t *testing.T) {
	cube2Flat2Cube(t, []uint{2, 4, 6, 8})
}
