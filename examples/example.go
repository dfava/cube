package main

import (
	"fmt"
	. "github.com/dfava/cube"
)

func printCubesOfDifferentSizes() {
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

func printAllRotations(n uint) {
	var cb Cube
	cb.Init(n)
	var c2 Cube
	for _, ax := range [...]Axis{Xax, Yax, Zax} {
		for idx := -int(n) / 2; idx <= int(n)/2; idx++ {
			for _, dir := range [...]bool{true, false} {
				if n%2 == 0 && idx == 0 {
					continue
				}
				fmt.Println(ax, idx, dir)
				fmt.Println(cb)
				fmt.Println()
				c2 = cb.Rotate(ax, idx, dir)
				fmt.Println(c2)
				fmt.Println()
			}
		}
	}
}

func main() {
	//PrintInColors(false)
	var cb Cube
	cb.Init(3)
	fmt.Println("Cubes of different sizes")
	printCubesOfDifferentSizes()
	fmt.Println()
	fmt.Println("Performing all moves on a Rubik's cube starting from the initial configuration")
	printAllRotations(3)
	fmt.Println()
	fmt.Println("Performing moves at random")
	cb.Init(3)
	fmt.Println(cb)
	for perm := 0; perm < 10; perm++ {
		cb.Permute(1)
		fmt.Println(cb)
		fmt.Println()
	}
	return
}
