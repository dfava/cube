package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/dfava/cube/internal"
)

type move struct {
	axis internal.Axis
	idx  int
	dir  internal.Direction
	desc string
}

func main() {
	internal.PrintInColors(true)
	var n uint = 3
	cb := internal.New(n)
	history := []internal.Cube{cb}
	moves := []move{}

	scanner := bufio.NewScanner(os.Stdin)
	printHelp(n)

	showCube := true
	for {
		if showCube {
			fmt.Printf("\nCube state (moves: %d):\n%s\n", len(history)-1, history[len(history)-1])
		}
		showCube = true
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			showCube = false
			continue
		}

		parts := strings.Fields(input)
		cmd := strings.ToLower(parts[0])

		switch cmd {
		case "h", "help":
			printHelp(n)
		case "n", "new":
			if len(parts) < 2 {
				fmt.Println("Invalid size. Usage: new <size>")
				showCube = false
				continue
			}
			newSize, err := strconv.ParseUint(parts[1], 10, 32)
			if err != nil || newSize < 1 {
				fmt.Printf("Invalid size: %s. Please provide a positive integer.\n", parts[1])
				showCube = false
				continue
			}
			n = uint(newSize)
			cb = internal.New(n)
			history = []internal.Cube{cb}
			moves = []move{}
			fmt.Printf("Created a new %dx%d cube.\n", n, n)
		case "q", "quit", "exit":
			fmt.Println("Goodbye!")
			return
		case "r", "reset":
			cb = internal.New(n)
			history = []internal.Cube{cb}
			moves = []move{}
			fmt.Println("Cube reset.")
		case "u", "undo":
			if len(history) > 1 {
				history = history[:len(history)-1]
				moves = moves[:len(moves)-1]
				fmt.Println("Undo successful.")
			} else {
				fmt.Println("Nothing to undo.")
			}
		case "s", "shuffle":
			// We'll do 20 random moves
			current := history[len(history)-1]
			for i := 0; i < 20; i++ {
				ax, idx, dir := randomMove(n)
				current = current.Turn(ax, idx, dir)
				if n%2 == 1 && idx == 0 {
					current = current.Rotate(ax, !dir)
				}
				history = append(history, current)
				moves = append(moves, move{ax, idx, dir, fmt.Sprintf("shuffle %s %d %s", ax, idx, dir)})
			}
			fmt.Println("Cube shuffled (20 moves added to history).")
		case "p", "playback":
			if len(moves) == 0 {
				fmt.Println("No history to play back.")
				showCube = false
				continue
			}
			fmt.Println("Playing back history:")
			for i, m := range moves {
				fmt.Printf("Step %d: %s\n", i+1, m.desc)
				fmt.Println(history[i+1])
				fmt.Println("--------------------")
			}
		case "x", "y", "z":
			if len(parts) < 3 {
				fmt.Println("Invalid move. Usage: <axis> <index> <direction: c|cc>")
				showCube = false
				continue
			}
			var ax internal.Axis
			switch cmd {
			case "x":
				ax = internal.Xax
			case "y":
				ax = internal.Yax
			case "z":
				ax = internal.Zax
			}
			idx, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Printf("Invalid index: %s\n", parts[1])
				showCube = false
				continue
			}
			// Validate index range
			minIdx := -int(n) / 2
			maxIdx := int(n) / 2
			if idx < minIdx || idx > maxIdx || (n%2 == 0 && idx == 0) {
				fmt.Printf("Invalid index: %d. Range for a %dx%d cube is %d to %d", idx, n, n, minIdx, maxIdx)
				if n%2 == 0 {
					fmt.Print(" (excluding 0)")
				}
				fmt.Println(".")
				showCube = false
				continue
			}
			var dir internal.Direction
			dirPart := strings.ToLower(parts[2])
			if dirPart == "c" {
				dir = internal.Clock
			} else if dirPart == "cc" {
				dir = internal.Counterclock
			} else {
				fmt.Printf("Invalid direction: %s. Use 'c' or 'cc'.\n", parts[2])
				showCube = false
				continue
			}

			current := history[len(history)-1]
			next := current.Turn(ax, idx, dir)
			if n%2 == 1 && idx == 0 {
				next = next.Rotate(ax, !dir)
			}
			history = append(history, next)
			moves = append(moves, move{ax, idx, dir, fmt.Sprintf("%s %d %s", ax, idx, dir)})
		default:
			fmt.Printf("Unknown command: %s. Type 'h' for help.\n", cmd)
			showCube = false
		}
	}
}

func printHelp(n uint) {
	fmt.Println("Rubik's Cube CLI")
	fmt.Println("Commands:")
	fmt.Println("  x <idx> <c|cc>  : Turn about X-axis at index <idx> (c: clockwise, cc: counter-clockwise)")
	fmt.Println("  y <idx> <c|cc>  : Turn about Y-axis at index <idx>")
	fmt.Println("  z <idx> <c|cc>  : Turn about Z-axis at index <idx>")
	fmt.Println("  u, undo         : Undo the last turn")
	fmt.Println("  s, shuffle      : Shuffle the cube (20 random moves)")
	fmt.Println("  r, reset        : Reset the cube to initial state")
	fmt.Println("  n, new <size>   : Create a new cube of size <size>")
	fmt.Println("  p, playback     : Play back the history of turns")
	fmt.Println("  h, help         : Show this help")
	fmt.Println("  q, quit         : Exit the application")
	fmt.Printf("\nIndex range for a %dx%d cube is %d to %d", n, n, -int(n)/2, int(n)/2)
	if n%2 == 0 {
		fmt.Print(" (excluding 0)")
	}
	fmt.Println(".")
}

func randomMove(n uint) (internal.Axis, int, internal.Direction) {
	axes := [...]internal.Axis{internal.Xax, internal.Yax, internal.Zax}
	idxs := []int{}
	for idx := -int(n) / 2; idx <= int(n)/2; idx++ {
		if n%2 == 0 && idx == 0 {
			continue
		}
		idxs = append(idxs, idx)
	}
	dirs := [...]internal.Direction{internal.Counterclock, internal.Clock}

	ax := axes[rand.Intn(len(axes))]
	idx := idxs[rand.Intn(len(idxs))]
	dir := dirs[rand.Intn(len(dirs))]
	return ax, idx, dir
}
