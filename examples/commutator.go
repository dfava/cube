// Copyright 2020 Daniel S. Fava. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	. "github.com/dfava/cube"
)

// The commutator of g and h is [g,h]:
//
// [g,h] = g^(-1) . h^(-1) . g . h
//
// On an commutative algebra, the commutator is the identity.
//
// The turns of a rubiks cube are a non-commutative algebra,
// so the commutator is not the identity,
// but its the "closest thing" to the identiy.
//
// We pick some turns to represent g and h and then we compute
// g.h as well as [g,h]
// Note that [g,h] is closer to the identify than g.h
func main() {
	//PrintInColors(false)
	var cb Cube
	cb.Init(3)
	fmt.Println(cb)
	fmt.Println()

	cb = cb.Turn(Xax, -1, Counterclock) // g
	cb = cb.Turn(Yax, -1, Counterclock) // h
	fmt.Println("g . h")
	fmt.Println(cb)
	fmt.Println()

	cb.Init(3)

	cb = cb.Turn(Xax, -1, Clock) // g^(-1)
	//fmt.Println(cb)
	//fmt.Println()
	cb = cb.Turn(Yax, -1, Clock) // h^(-1)
	//fmt.Println("g^(-1) . h^(-1)")
	//fmt.Println(cb)
	//fmt.Println()

	cb = cb.Turn(Xax, -1, Counterclock) // g
	//fmt.Println(cb)
	//fmt.Println()
	cb = cb.Turn(Yax, -1, Counterclock) // h
	fmt.Println("g^(-1) . h^(-1) . g . h")
	fmt.Println(cb)
	fmt.Println()

	return
}
