// Copyright 2020 Daniel S. Fava. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cube_test

import (
	"fmt"
	. "github.com/dfava/cube"
	"testing"
	"time"
)

func testSolver(t *testing.T, sizes []uint) {
	for _, n := range sizes {
		for _, shuffle := range [...]uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
			var cube Cube
			cube.Init(n)
			cube.Shuffle(shuffle)
			var init Cube
			init.Init(n)
			fmt.Printf("n=%d, shuffle=%d\n", n, shuffle)
			fmt.Println("Starting position:")
			fmt.Println(cube)
			start := time.Now()
			edges := (cube.GetPath(init)).(map[string]string)
			end := time.Now()
			step := 0
			for end := init.String(); end != ""; end = edges[end] {
				fmt.Println(step)
				fmt.Println(end)
				fmt.Println()
				step += 1
			}
			fmt.Println(end.Sub(start))
		}
	}
}

func TestSolverOdd(t *testing.T) {
	//testSolver(t, []uint{3, 5, 7, 9})
	testSolver(t, []uint{3}) //, 5, 7, 9})
}

func TestSolverEven(t *testing.T) {
	testSolver(t, []uint{2, 4, 6, 8})
}
