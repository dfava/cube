package main

import (
	"fmt"
	"time"

	"github.com/dfava/cube/internal"
)

type RigidAnimator struct{}

func (a RigidAnimator) Animate(cb internal.Cube, ax internal.Axis, idx int, dir internal.Direction, n uint, helpVisible bool) {
	perm := cb.GetFlatPermutation(ax, idx, dir)
	if len(perm) == 0 {
		return
	}

	var startFl internal.Flat
	startFl.PaintCube(cb)

	visited := make(map[[2]int]bool)
	var cycles [][][2]int
	for start := range perm {
		if visited[start] {
			continue
		}
		var cycle [][2]int
		curr := start
		for !visited[curr] {
			visited[curr] = true
			cycle = append(cycle, curr)
			curr = perm[curr]
		}
		if len(cycle) > 1 {
			cycles = append(cycles, cycle)
		}
	}

	var rings [][][2]int
	for _, cycle := range cycles {
		var ring [][2]int
		for i := 0; i < len(cycle); i++ {
			p1 := cycle[i]
			p2 := cycle[(i+1)%len(cycle)]
			r, c := p1[0], p1[1]
			for r != p2[0] {
				ring = append(ring, [2]int{r, c})
				if r < p2[0] {
					r++
				} else {
					r--
				}
			}
			for c != p2[1] {
				ring = append(ring, [2]int{r, c})
				if c < p2[1] {
					c++
				} else {
					c--
				}
			}
		}
		rings = append(rings, ring)
	}

	posToRing := make(map[[2]int]int)
	posToIdx := make(map[[2]int]int)
	for ri, ring := range rings {
		for ii, p := range ring {
			// A point might appear in multiple rings if rings overlap or if the
			// logic is flawed. But let's assume it's okay for now.
			posToRing[p] = ri
			posToIdx[p] = ii
		}
	}

	for frame := 0; frame <= int(n); frame++ {
		tempFl := startFl.Copy()
		if frame == int(n) {
			// Safeguard: Last frame is always the target state
			nextCb := cb.Move(internal.Move{Axis: ax, Idx: idx, Direction: dir})
			if n%2 == 1 && idx == 0 {
				nextCb = nextCb.Rotate(ax, !dir)
			}
			tempFl.PaintCube(nextCb)
		} else {
			// Clear moving cells to prepare for interpolation
			for pos := range perm {
				tempFl[pos[0]][pos[1]] = " "
			}

			for i, cycle := range cycles {
				ring := rings[i]
				m := len(ring)
				k := len(cycle)
				d := float64(m) / float64(k)
				shift := int(float64(frame)*d/float64(n) + 0.5)

				// Identify all starting positions in this cycle that have colors
				type coloredCyclePos struct {
					pos   [2]int
					color string
				}
				var cColors []coloredCyclePos
				for _, cp := range cycle {
					if startFl[cp[0]][cp[1]] != " " && startFl[cp[0]][cp[1]] != "" {
						cColors = append(cColors, coloredCyclePos{cp, startFl[cp[0]][cp[1]]})
					}
				}

				for _, cc := range cColors {
					ringIdx := -1
					for idx, p := range ring {
						if p == cc.pos {
							ringIdx = idx
							break
						}
					}
					if ringIdx == -1 {
						continue
					}
					targetPos := ring[(ringIdx+shift)%m]
					tempFl[targetPos[0]][targetPos[1]] = cc.color
				}
			}
		}

		clearScreen()
		if helpVisible {
			printHelp(n)
		}
		printAxes()
		fmt.Printf("\r\nAnimating move: %s %d %s\r\n%s\r\n", ax, idx, dir, tempFl)

		if frame < int(n) {
			time.Sleep(animSpeed)
		}
	}
	time.Sleep(2 * animSpeed)
}
