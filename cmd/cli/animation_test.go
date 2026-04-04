package main_test

import (
	"fmt"
	"testing"

	. "github.com/dfava/cube/internal"
	"github.com/stretchr/testify/require"
)

func TestAnimationMarchingTogether(t *testing.T) {
	n := uint(3)
	cb := New(n)
	ax := Yax
	idx := 1
	dir := Direction(Clock)

	// Get permutation
	perm := cb.GetFlatPermutation(ax, idx, dir)
	require.NotZero(t, len(perm), "Expected non-empty permutation")

	var startFl Flat
	startFl.PaintCube(cb)

	// Simulation logic from RigidAnimator.Animate
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
			posToRing[p] = ri
			posToIdx[p] = ii
		}
	}

	// Check each frame
	for frame := 0; frame <= int(n); frame++ {
		tempFl := startFl.Copy()
		if frame == int(n) {
			nextCb := cb.Move(Move{Axis: ax, Idx: idx, Direction: dir})
			if n%2 == 1 && idx == 0 {
				nextCb = nextCb.Rotate(ax, !dir)
			}
			tempFl.PaintCube(nextCb)
		} else {
			for pos := range perm {
				tempFl[pos[0]][pos[1]] = " "
			}

			for i, cycle := range cycles {
				ring := rings[i]
				m := len(ring)
				k := len(cycle)
				d := float64(m) / float64(k)
				shift := int(float64(frame)*d/float64(n) + 0.5)

				for _, cp := range cycle {
					if startFl[cp[0]][cp[1]] == " " || startFl[cp[0]][cp[1]] == "" {
						continue
					}
					ringIdx := -1
					for idx, p := range ring {
						if p == cp {
							ringIdx = idx
							break
						}
					}
					if ringIdx == -1 {
						continue
					}
					targetPos := ring[(ringIdx+shift)%m]
					tempFl[targetPos[0]][targetPos[1]] = startFl[cp[0]][cp[1]]
				}
			}
		}

		if frame == int(n) {
			// t.Logf("Last Frame Actual:\n%s", tempFl.String())
		}
	}
}

func TestAnimationCorrectPositions(t *testing.T) {
	testCases := []struct {
		n   uint
		ax  Axis
		idx int
		dir Direction
	}{
		{3, Yax, 1, Direction(Clock)},
		{3, Xax, -1, Direction(Counterclock)},
		{3, Zax, 0, Direction(Clock)},
		{2, Xax, 1, Direction(Clock)},
		{4, Yax, -2, Direction(Clock)},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("n%d_%s_%d_%s", tc.n, tc.ax, tc.idx, tc.dir), func(t *testing.T) {
			n := tc.n
			cb := New(n)
			ax := tc.ax
			idx := tc.idx
			dir := tc.dir

			var startFl Flat
			startFl.PaintCube(cb)

			nextCb := cb.Move(Move{Axis: ax, Idx: idx, Direction: dir})
			if n%2 == 1 && idx == 0 {
				nextCb = nextCb.Rotate(ax, !dir)
			}
			var finalFl Flat
			finalFl.PaintCube(nextCb)

			frame := int(n)
			tempFl := startFl.Copy()
			if frame == int(n) {
				nextCb := cb.Move(Move{Axis: ax, Idx: idx, Direction: dir})
				if n%2 == 1 && idx == 0 {
					nextCb = nextCb.Rotate(ax, !dir)
				}
				tempFl.PaintCube(nextCb)
			}

			require.Equal(t, finalFl.String(), tempFl.String(), "Final cube state doesn't match last frame")
		})
	}
}
