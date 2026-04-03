// Copyright 2020 Daniel S. Fava. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	. "github.com/dfava/cube/internal"
)

func main() {
	//PrintInColors(false)
	cb := New(3)
	fmt.Println(cb)
	fmt.Println()

	cb = swap_top_corners(cb)
	cb = cb.Turn(Zax, 1, Clock)

	cb = swap_top_corners(cb)
	cb = cb.Turn(Zax, 1, Clock) //cb = cb.Turn(Zax, 1, Counterclock)

	fmt.Println(cb)
	fmt.Println()
}
