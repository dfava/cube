package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dfava/cube/internal"
	"golang.org/x/term"
)

const animSpeed = 250 * time.Millisecond

type Animator interface {
	Animate(cb internal.Cube, ax internal.Axis, idx int, dir internal.Direction, n uint, helpVisible bool)
}

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
	var animator Animator = DummyAnimator{}

	// cmdHistory stores previously entered commands for up-arrow navigation.
	cmdHistory := []string{}
	const maxHistory = 1000

	// Put terminal in raw mode for x/term to handle input properly.
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error putting terminal in raw mode: %v\r\n", err)
		return
	}
	defer func() {
		_ = term.Restore(int(os.Stdin.Fd()), oldState)
	}()

	t := term.NewTerminal(os.Stdin, "> ")

	printHelp(n)
	helpVisible := true
	showCube := true

	for {
		if showCube {
			clearScreen()
			if helpVisible {
				printHelp(n)
			}
			printAxes()
			fmt.Printf("\r\nCube state (moves: %d):\r\n%s\r\n", len(history)-1, history[len(history)-1])
		}
		showCube = true

		input, err := t.ReadLine()
		if err != nil {
			break
		}

		input = strings.TrimSpace(input)
		if input == "" {
			showCube = false
			continue
		}

		// Add to command history if it's different from the last one.
		if len(cmdHistory) == 0 || cmdHistory[len(cmdHistory)-1] != input {
			cmdHistory = append(cmdHistory, input)
			if len(cmdHistory) > maxHistory {
				cmdHistory = cmdHistory[1:]
			}
		}

		parts := strings.Fields(input)
		cmd := strings.ToLower(parts[0])

		switch cmd {
		case "h", "help":
			helpVisible = true
			printHelp(n)
		case "n", "new":
			if len(parts) < 2 {
				fmt.Println("Invalid size. Usage: new <size>\r")
				showCube = false
				continue
			}
			newSize, err := strconv.ParseUint(parts[1], 10, 32)
			if err != nil || newSize < 1 {
				fmt.Printf("Invalid size: %s. Please provide a positive integer.\r\n", parts[1])
				showCube = false
				continue
			}
			n = uint(newSize)
			cb = internal.New(n)
			history = []internal.Cube{cb}
			moves = []move{}
			helpVisible = true
			fmt.Printf("Created a new %dx%d cube.\r\n", n, n)
		case "q", "quit", "exit":
			fmt.Println("Goodbye!\r")
			return
		case "r", "reset":
			cb = internal.New(n)
			history = []internal.Cube{cb}
			moves = []move{}
			helpVisible = true
			fmt.Println("Cube reset.\r")
		case "u", "undo":
			if len(history) > 1 {
				history = history[:len(history)-1]
				moves = moves[:len(moves)-1]
				fmt.Println("Undo successful.\r")
			} else {
				fmt.Println("Nothing to undo.\r")
			}
		case "s", "shuffle":
			current := history[len(history)-1]
			for range 20 {
				ax, idx, dir := randomMove(n)
				current = current.Move(internal.Move{Axis: ax, Idx: idx, Direction: dir})
				if n%2 == 1 && idx == 0 {
					current = current.Rotate(ax, !dir)
				}
				history = append(history, current)
				moves = append(moves, move{ax, idx, dir, fmt.Sprintf("shuffle %s %d %s", ax, idx, dir)})
			}
			fmt.Println("Cube shuffled (20 moves added to history).\r")
		case "p", "playback":
			if len(moves) == 0 {
				fmt.Println("No history to play back.\r")
				showCube = false
				continue
			}
			fmt.Println("Playing back history:\r")
			for i, m := range moves {
				fmt.Printf("Step %d: %s\r\n", i+1, m.desc)
				animator.Animate(history[i], m.axis, m.idx, m.dir, n, helpVisible)
				time.Sleep(3 * animSpeed)
			}
		case "x", "y", "z":
			if len(parts) < 3 {
				fmt.Println("Invalid move. Usage: <axis> <index> <direction: c|cc>\r")
				showCube = false
				continue
			}
			ax, _ := internal.ParseAxis(cmd)
			idx, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Printf("Invalid index: %s\r\n", parts[1])
				showCube = false
				continue
			}
			minIdx := -int(n) / 2
			maxIdx := int(n) / 2
			if idx < minIdx || idx > maxIdx || (n%2 == 0 && idx == 0) {
				fmt.Printf("Invalid index: %d. Range for a %dx%d cube is %d to %d", idx, n, n, minIdx, maxIdx)
				if n%2 == 0 {
					fmt.Print(" (excluding 0)")
				}
				fmt.Println(".\r")
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
				fmt.Printf("Invalid direction: %s. Use 'c' or 'cc'.\r\n", parts[2])
				showCube = false
				continue
			}

			current := history[len(history)-1]
			next := current.Move(internal.Move{Axis: ax, Idx: idx, Direction: dir})
			if n%2 == 1 && idx == 0 {
				next = next.Rotate(ax, !dir)
			}
			animator.Animate(current, ax, idx, dir, n, helpVisible)
			history = append(history, next)
			moves = append(moves, move{ax, idx, dir, fmt.Sprintf("%s %d %s", ax, idx, dir)})
		default:
			fmt.Printf("Unknown command: %s. Type 'h' for help.\r\n", cmd)
			helpVisible = false
			showCube = false
		}
	}
}

func printHelp(n uint) {
	fmt.Println("Rubik's Cube CLI\r")
	fmt.Println("Commands:\r")
	fmt.Println("  x <idx> <c|cc>  : Turn about X-axis at index <idx> (c: clockwise, cc: counter-clockwise)\r")
	fmt.Println("  y <idx> <c|cc>  : Turn about Y-axis at index <idx>\r")
	fmt.Println("  z <idx> <c|cc>  : Turn about Z-axis at index <idx>\r")
	// Added a small tip about history
	fmt.Println("  [Up Arrow]      : Recall previous command\r")
	fmt.Println("  u, undo         : Undo the last turn\r")
	fmt.Println("  s, shuffle      : Shuffle the cube (20 random moves)\r")
	fmt.Println("  r, reset        : Reset the cube to initial state\r")
	fmt.Println("  n, new <size>   : Create a new cube of size <size>\r")
	fmt.Println("  p, playback     : Play back the history of turns\r")
	fmt.Println("  h, help         : Show this help\r")
	fmt.Println("  q, quit         : Exit the application\r")
	fmt.Printf("\r\nIndex range for a %dx%d cube is %d to %d", n, n, -int(n)/2, int(n)/2)
	if n%2 == 0 {
		fmt.Print(" (excluding 0)")
	}
	fmt.Println(".\r")
}

func printAxes() {
	fmt.Println("\r\nAxis directions:\r")
	fmt.Println("   z\r")
	fmt.Println("   |\r")
	fmt.Println("   +--- x\r")
	fmt.Println("  /\r")
	fmt.Println(" y\r")
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

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
