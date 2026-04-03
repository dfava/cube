// Copyright 2020 Daniel S. Fava. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/dfava/cube/internal"
)

// Turn middle cubi at (0,1,1) so that the green and the yellow are reversed
func turn_middle(cb internal.Cube) internal.Cube {
	// Turn
	cb = cb.Turn(internal.Xax, 0, internal.Counterclock)
	cb = cb.Turn(internal.Zax, -1, internal.Counterclock)
	cb = cb.Turn(internal.Zax, -1, internal.Counterclock)
	cb = cb.Turn(internal.Xax, 0, internal.Clock)
	// Notice that these next four moves are a commutator!
	cb = cb.Turn(internal.Zax, -1, internal.Clock)
	cb = cb.Turn(internal.Xax, 0, internal.Counterclock)
	cb = cb.Turn(internal.Zax, -1, internal.Counterclock)
	cb = cb.Turn(internal.Xax, 0, internal.Clock)
	//fmt.Println(cb)
	//fmt.Println()

	//
	cb = cb.Turn(internal.Zax, 1, internal.Counterclock)
	cb = cb.Turn(internal.Zax, 1, internal.Counterclock)
	//fmt.Println(cb)
	//fmt.Println()

	// Reverse turn
	cb = cb.Turn(internal.Xax, 0, internal.Counterclock)
	cb = cb.Turn(internal.Zax, -1, internal.Clock)
	cb = cb.Turn(internal.Xax, 0, internal.Clock)
	cb = cb.Turn(internal.Zax, -1, internal.Counterclock)
	cb = cb.Turn(internal.Xax, 0, internal.Counterclock)
	cb = cb.Turn(internal.Zax, -1, internal.Clock)
	cb = cb.Turn(internal.Zax, -1, internal.Clock)
	cb = cb.Turn(internal.Xax, 0, internal.Clock)
	//fmt.Println(cb)
	//fmt.Println()

	cb = cb.Turn(internal.Zax, 1, internal.Clock)
	cb = cb.Turn(internal.Zax, 1, internal.Clock)
	return cb
}

func main() {
	internal.PrintInColors(false)
	cb := internal.New(3)
	fmt.Println(cb)
	fmt.Println()

	cb = turn_middle(cb)
	fmt.Println(cb)
	fmt.Println()
}
