// Copyright 2020 Daniel S. Fava. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	. "github.com/dfava/cube"
)

// Perform moves at random, starting from the initial configuration
func main() {
	//PrintInColors(false)
	var cb Cube
	cb.Init(3)
	fmt.Print("shuffle: performing moves at random, ")
	fmt.Println("starting from the initial configuration")
	fmt.Println(cb)
	fmt.Println()
	for perm := 0; perm < 5; perm++ {
		cb.Shuffle(1)
		fmt.Println(cb)
		fmt.Println()
	}
	return
}
