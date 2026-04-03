// Copyright 2020 Daniel S. Fava. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/dfava/cube/internal"
)

func swap_top_corners(cb internal.Cube) internal.Cube {
	// Swap cubi (1,1,1) with cubi (-1,1,1)
	cb = cb.Turn(internal.Xax, 1, internal.Counterclock)
	cb = cb.Turn(internal.Zax, -1, internal.Counterclock)
	cb = cb.Turn(internal.Xax, 1, internal.Clock)

	cb = cb.Turn(internal.Xax, -1, internal.Counterclock)
	cb = cb.Turn(internal.Zax, -1, internal.Clock)
	cb = cb.Turn(internal.Xax, -1, internal.Clock)

	cb = cb.Turn(internal.Xax, 1, internal.Counterclock)
	cb = cb.Turn(internal.Zax, -1, internal.Counterclock)
	cb = cb.Turn(internal.Xax, 1, internal.Clock)

	//fmt.Println(cb)
	//fmt.Println()

	//
	cb = cb.Turn(internal.Zax, 1, internal.Counterclock)
	cb = cb.Turn(internal.Zax, 1, internal.Counterclock)
	//fmt.Println(cb)
	//fmt.Println()

	// Reverse swap
	cb = cb.Turn(internal.Xax, 1, internal.Counterclock)
	cb = cb.Turn(internal.Zax, -1, internal.Clock)
	cb = cb.Turn(internal.Xax, 1, internal.Clock)

	cb = cb.Turn(internal.Xax, -1, internal.Counterclock)
	cb = cb.Turn(internal.Zax, -1, internal.Counterclock)
	cb = cb.Turn(internal.Xax, -1, internal.Clock)

	cb = cb.Turn(internal.Xax, 1, internal.Counterclock)
	cb = cb.Turn(internal.Zax, -1, internal.Clock)
	cb = cb.Turn(internal.Xax, 1, internal.Clock)

	//fmt.Println(cb)
	//fmt.Println()

	//
	cb = cb.Turn(internal.Zax, 1, internal.Clock)
	cb = cb.Turn(internal.Zax, 1, internal.Clock)
	return cb
}

func main() {
	internal.PrintInColors(false)
	cb := internal.New(3)
	fmt.Println(cb)
	fmt.Println()

	cb = swap_top_corners(cb)
	cb = cb.Turn(internal.Zax, 1, internal.Clock)

	cb = swap_top_corners(cb)
	cb = cb.Turn(internal.Zax, 1, internal.Clock) //cb = cb.Turn(Zax, 1, Counterclock)

	fmt.Println(cb)
	fmt.Println()
}
