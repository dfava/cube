package main

import (
	"github.com/dfava/cube/internal"
)

type DummyAnimator struct{}

func (a DummyAnimator) Animate(cb internal.Cube, ax internal.Axis, idx int, dir internal.Direction, n uint, helpVisible bool) {
	// Dummy animator doesn't animate, it just does nothing.
	// The final state will be printed by the main loop.
}
