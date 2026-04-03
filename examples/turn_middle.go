// Copyright 2020 Daniel S. Fava. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	. "github.com/dfava/cube/internal"
)

// Turn middle cubi at (0,1,1) so that the green and the yellow are reversed
func turn_middle(cb Cube) Cube {
	// Turn
	cb = cb.Turn(Xax, 0, Counterclock)
	cb = cb.Turn(Zax, -1, Counterclock)
	cb = cb.Turn(Zax, -1, Counterclock)
	cb = cb.Turn(Xax, 0, Clock)
	// Notice that these next four moves are a commutator!
	cb = cb.Turn(Zax, -1, Clock)
	cb = cb.Turn(Xax, 0, Counterclock)
	cb = cb.Turn(Zax, -1, Counterclock)
	cb = cb.Turn(Xax, 0, Clock)
	//fmt.Println(cb)
	//fmt.Println()

	//
	cb = cb.Turn(Zax, 1, Counterclock)
	cb = cb.Turn(Zax, 1, Counterclock)
	//fmt.Println(cb)
	//fmt.Println()

	// Reverse turn
	cb = cb.Turn(Xax, 0, Counterclock)
	cb = cb.Turn(Zax, -1, Clock)
	cb = cb.Turn(Xax, 0, Clock)
	cb = cb.Turn(Zax, -1, Counterclock)
	cb = cb.Turn(Xax, 0, Counterclock)
	cb = cb.Turn(Zax, -1, Clock)
	cb = cb.Turn(Zax, -1, Clock)
	cb = cb.Turn(Xax, 0, Clock)
	//fmt.Println(cb)
	//fmt.Println()

	cb = cb.Turn(Zax, 1, Clock)
	cb = cb.Turn(Zax, 1, Clock)
	return cb
}

func main() {
	//PrintInColors(false)
	cb := New(3)
	fmt.Println(cb)
	fmt.Println()

	cb = turn_middle(cb)
	fmt.Println(cb)
	fmt.Println()
}
