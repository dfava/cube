package main

import (
	. "github.com/dfava/cube/internal"
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
