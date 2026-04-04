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
	cb = cb.Move(internal.Move{Axis: internal.Xax, Idx: 0, Direction: internal.Counterclock})
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: -1, Direction: internal.Counterclock})
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: -1, Direction: internal.Counterclock})
	cb = cb.Move(internal.Move{Axis: internal.Xax, Idx: 0, Direction: internal.Clock})
	// Notice that these next four moves are a commutator!
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: -1, Direction: internal.Clock})
	cb = cb.Move(internal.Move{Axis: internal.Xax, Idx: 0, Direction: internal.Counterclock})
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: -1, Direction: internal.Counterclock})
	cb = cb.Move(internal.Move{Axis: internal.Xax, Idx: 0, Direction: internal.Clock})
	//fmt.Println(cb)
	//fmt.Println()

	//
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: 1, Direction: internal.Counterclock})
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: 1, Direction: internal.Counterclock})
	//fmt.Println(cb)
	//fmt.Println()

	// Reverse turn
	cb = cb.Move(internal.Move{Axis: internal.Xax, Idx: 0, Direction: internal.Counterclock})
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: -1, Direction: internal.Clock})
	cb = cb.Move(internal.Move{Axis: internal.Xax, Idx: 0, Direction: internal.Clock})
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: -1, Direction: internal.Counterclock})
	cb = cb.Move(internal.Move{Axis: internal.Xax, Idx: 0, Direction: internal.Counterclock})
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: -1, Direction: internal.Clock})
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: -1, Direction: internal.Clock})
	cb = cb.Move(internal.Move{Axis: internal.Xax, Idx: 0, Direction: internal.Clock})
	//fmt.Println(cb)
	//fmt.Println()

	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: 1, Direction: internal.Clock})
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: 1, Direction: internal.Clock})
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
