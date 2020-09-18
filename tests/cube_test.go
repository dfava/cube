// Copyright 2020 Daniel S. Fava. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cube_test

import (
	//"fmt"
	. "github.com/dfava/cube"
	"math/rand"
	"testing"
)

func TestShuffleCanonical(t *testing.T) {
	for _, n := range []uint{3, 5, 7, 9} {
		for _, shuffle := range [...]uint{0, 1, 2, 10, 20, 100} {
			var cube Cube
			cube.Init(n)
			cube.Shuffle(shuffle)
			if !cube.IsCanonical() {
				t.Errorf("Cube not canonical! n=%d", n)
				t.Errorf("\n%s", cube.String())
			}
		}
	}
}

type turnOperands struct {
	ax  Axis
	idx int
	dir Direction
}

func TestTurn(t *testing.T) {
	var times uint = 100
	for _, n := range []uint{3, 5, 7, 9} {
		axes := [...]Axis{Xax, Yax, Zax}
		idxs := make([]int, n+(n+1)%2)
		for idx := -int(n) / 2; idx <= int(n)/2; idx++ {
			idxs[idx+int(n)/2] = idx
		}
		dirs := [...]Direction{Counterclock, Clock}
		var perms []turnOperands
		var num_perms int
		for num_perms := 0; num_perms < int(times); num_perms += 1 {
			ax := axes[rand.Intn(len(axes))]  // pick an axis
			idx := idxs[rand.Intn(len(idxs))] // pick an index
			dir := dirs[rand.Intn(len(dirs))] // pick a direction
			if n%2 == 0 && idx == 0 {
				continue
			}
			perms = append(perms, turnOperands{ax: ax, idx: idx, dir: dir})
		}
		var cube Cube
		cube.Init(n)
		var final Cube
		final.Init(n)
		// Perform turns
		for num_perms = 0; num_perms < int(times); num_perms += 1 {
			//fmt.Println(perms[num_perms])
			final = final.Turn(perms[num_perms].ax, perms[num_perms].idx, perms[num_perms].dir)
		}
		//fmt.Println()
		//fmt.Println(final)
		// Perform, in reverse, the reverse of the turns
		for num_perms = int(times) - 1; num_perms >= 0; num_perms -= 1 {
			//fmt.Println(perms[num_perms])
			final = final.Turn(perms[num_perms].ax, perms[num_perms].idx, !perms[num_perms].dir)
		}
		// Make sure init and final are the same
		if cube.String() != final.String() {
			t.Errorf("turns failed! n=%d", n)
		}
	}
}
