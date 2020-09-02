// Copyright 2020 Daniel S. Fava. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	. "github.com/dfava/cube"
)

func main() {
	//PrintInColors(false)
	fmt.Println("Cubes of different sizes")
	var cb Cube
	cb.Init(2)
	fmt.Println(cb)
	fmt.Println()
	cb.Init(3)
	fmt.Println(cb)
	fmt.Println()
	cb.Init(4)
	fmt.Println(cb)
	fmt.Println()
	cb.Init(5)
	fmt.Println(cb)
	fmt.Println()
}
