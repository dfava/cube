// Copyright 2020 Daniel S. Fava. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/dfava/cube/internal"
)

func main() {
	internal.PrintInColors(false)
	fmt.Println("Printing all possible moves of a Rubik's cube,")
	fmt.Println("starting from the initial configuration.")
	fmt.Println()
	var n uint
	n = 3
	cb := internal.New(n)
	var c2 internal.Cube
	for _, ax := range [...]internal.Axis{internal.Xax, internal.Yax, internal.Zax} {
		for idx := -int(n) / 2; idx <= int(n)/2; idx++ {
			for _, dir := range [...]internal.Direction{internal.Counterclock, internal.Clock} {
				if n%2 == 0 && idx == 0 {
					continue
				}
				fmt.Printf("Moving about the %s-axis ", ax)
				fmt.Printf("at index %d in the direction %s\n", idx, dir)
				fmt.Println(cb)
				fmt.Println()
				c2 = cb.Turn(ax, idx, dir)
				fmt.Println(c2)
				fmt.Println()
			}
		}
	}
}
