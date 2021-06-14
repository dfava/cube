// Copyright 2020 Daniel S. Fava. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	. "github.com/dfava/cube"
)

func swap_top_corners(cb Cube) Cube {
	// Swap cubi (1,1,1) with cubi (-1,1,1)
	cb = cb.Turn(Xax, 1, Counterclock)
	cb = cb.Turn(Zax, -1, Counterclock)
	cb = cb.Turn(Xax, 1, Clock)

	cb = cb.Turn(Xax, -1, Counterclock)
	cb = cb.Turn(Zax, -1, Clock)
	cb = cb.Turn(Xax, -1, Clock)

	cb = cb.Turn(Xax, 1, Counterclock)
	cb = cb.Turn(Zax, -1, Counterclock)
	cb = cb.Turn(Xax, 1, Clock)

	//fmt.Println(cb)
	//fmt.Println()

	//
	cb = cb.Turn(Zax, 1, Counterclock)
	cb = cb.Turn(Zax, 1, Counterclock)
	//fmt.Println(cb)
	//fmt.Println()

	// Reverse swap
	cb = cb.Turn(Xax, 1, Counterclock)
	cb = cb.Turn(Zax, -1, Clock)
	cb = cb.Turn(Xax, 1, Clock)

	cb = cb.Turn(Xax, -1, Counterclock)
	cb = cb.Turn(Zax, -1, Counterclock)
	cb = cb.Turn(Xax, -1, Clock)

	cb = cb.Turn(Xax, 1, Counterclock)
	cb = cb.Turn(Zax, -1, Clock)
	cb = cb.Turn(Xax, 1, Clock)

	//fmt.Println(cb)
	//fmt.Println()

	//
	cb = cb.Turn(Zax, 1, Clock)
	cb = cb.Turn(Zax, 1, Clock)
	return cb
}

func main() {
	//PrintInColors(false)
	var cb Cube
	cb.Init(3)
	fmt.Println(cb)
	fmt.Println()

	cb = swap_top_corners(cb)
	cb = cb.Turn(Zax, 1, Clock)

	cb = swap_top_corners(cb)
	cb = cb.Turn(Zax, 1, Clock) //cb = cb.Turn(Zax, 1, Counterclock)

	fmt.Println(cb)
	fmt.Println()
}
