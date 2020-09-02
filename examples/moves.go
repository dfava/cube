// Copyright 2020 Daniel S. Fava. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	. "github.com/dfava/cube"
)

func main() {
	//PrintInColors(false)
	fmt.Println("Printing all possible moves of a Rubik's cube,")
	fmt.Println("starting from the initial configuration.")
	fmt.Println()
	var n uint
	n = 3
	var cb Cube
	cb.Init(n)
	var c2 Cube
	for _, ax := range [...]Axis{Xax, Yax, Zax} {
		for idx := -int(n) / 2; idx <= int(n)/2; idx++ {
			for _, dir := range [...]Direction{Counterclock, Clock} {
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
