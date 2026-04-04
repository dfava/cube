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
	cb = cb.Move(internal.Move{Axis: internal.Xax, Idx: 1, Direction: internal.Counterclock})
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: -1, Direction: internal.Counterclock})
	cb = cb.Move(internal.Move{Axis: internal.Xax, Idx: 1, Direction: internal.Clock})

	cb = cb.Move(internal.Move{Axis: internal.Xax, Idx: -1, Direction: internal.Counterclock})
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: -1, Direction: internal.Clock})
	cb = cb.Move(internal.Move{Axis: internal.Xax, Idx: -1, Direction: internal.Clock})

	cb = cb.Move(internal.Move{Axis: internal.Xax, Idx: 1, Direction: internal.Counterclock})
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: -1, Direction: internal.Counterclock})
	cb = cb.Move(internal.Move{Axis: internal.Xax, Idx: 1, Direction: internal.Clock})

	//fmt.Println(cb)
	//fmt.Println()

	//
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: 1, Direction: internal.Counterclock})
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: 1, Direction: internal.Counterclock})
	//fmt.Println(cb)
	//fmt.Println()

	// Reverse swap
	cb = cb.Move(internal.Move{Axis: internal.Xax, Idx: 1, Direction: internal.Counterclock})
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: -1, Direction: internal.Clock})
	cb = cb.Move(internal.Move{Axis: internal.Xax, Idx: 1, Direction: internal.Clock})

	cb = cb.Move(internal.Move{Axis: internal.Xax, Idx: -1, Direction: internal.Counterclock})
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: -1, Direction: internal.Counterclock})
	cb = cb.Move(internal.Move{Axis: internal.Xax, Idx: -1, Direction: internal.Clock})

	cb = cb.Move(internal.Move{Axis: internal.Xax, Idx: 1, Direction: internal.Counterclock})
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: -1, Direction: internal.Clock})
	cb = cb.Move(internal.Move{Axis: internal.Xax, Idx: 1, Direction: internal.Clock})

	//fmt.Println(cb)
	//fmt.Println()

	//
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: 1, Direction: internal.Clock})
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: 1, Direction: internal.Clock})
	return cb
}

func main() {
	internal.PrintInColors(false)
	cb := internal.New(3)
	fmt.Println(cb)
	fmt.Println()

	cb = swap_top_corners(cb)
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: 1, Direction: internal.Clock})

	cb = swap_top_corners(cb)
	cb = cb.Move(internal.Move{Axis: internal.Zax, Idx: 1, Direction: internal.Clock}) //cb = cb.Turn(Zax, 1, Counterclock)

	fmt.Println(cb)
	fmt.Println()
}
