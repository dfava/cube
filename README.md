# cube

A Rubik's cube implementation in Go.  Supports cubes of size `n x n x n` for `n` larger than one.

## Use

Installing the package:

```
go get -v github.com/dfava/cube
```

then using it:

```
package main

import (
	"fmt"
	"github.com/dfava/cube"
)

func main() {
	var cb Cube
	cb.Init(3)
	fmt.Println(cb)
}
```

For more, see `example.go`.


## Representation

I wanted a representation where we could leverage the cube's geometry.  By geometry I mean that I wanted to be able to use rotation matrices and basic linear algebra to manipulate the cube.  The cube is an object in three dimensions, in principle, we should be able to capture exactly that.  In practice, a representation that capture the cube's geometry would generalize better.  It should be as easy to create and manipulate a `7x7x7` cube (or any other configuration) as it is a `3x3x3`, which is the standard Rubik's cube.

### Why not a matrix?

A cube has six sides, so, one obvious possible representation is as an array of length 6 containing elements that are `nxn`.  In other words, to flatten out the cube into 2D arrays.
Another possible representation is to create an `nxnxn` matrix.

These representations relate the geometry of the cube with the memory layout of its implementation.  This gives us the sense that we are accurately capturing the essence of what it is to be a cube.  This "sense" is more an "illusion."  I would say that the essence of the Rubik's cube is not its static, physical configuration but its dynamism, meaning, *how its configuration is allowed to change*.  The essence of a Rubik's cube are the operations, the moves we can make.

### A list of vectors

The most natural way to position something in space is through a vector.  Since we are in 3D, we'll use a three dimensional vector: `[x, y, z]`.  We will break down a cube into the little cubes that constitute the cube.  I'll call these little cubes *cubis*.

For example, the Rubik's cube is made of `3*3*3-1=27-1=26` cubis.
There are `3*3*3` little cubes in the Rubik's cube, but the little cube at the center doesn't play a role (we never get to see it! So we do not need to represent the internal pieces).  That is why we subtract one from `3*3*3` and we say the Rubik's cube is composed of `26` cubis as opposed to `27`.

In general, an `nxnxn` cube has `n^3 - (n-2)^3` cubis.  These are the external facing little cubes.  We represent each cubi with two vectors: a *position vector* and a *color vector*.
The position vector captures where the cubi is in space.
The color vector captures how to color the cubi.

### Coordinate system and Canonical cube

I'm taking the axis to be

```
    z
    |
    |
    |
    ._________ x
   /
  /
 /
y
```

And the canonical cube to be as follows.  Take a die with the `1` facing you and two facing down.  Then three will be facing right, four facing left, five up, and six away from you.

1. Green 
2. White
3. Orange
4. Red
5. Yellow
6. Blue
 
The Rubik's cube is then represented in with the following array of vectors:

```
[{[-1 -1 -1] [-4 -6 -2]} {[-1 -1 0] [-4 -6 0]} {[-1 -1 1] [-4 -6 5]} {[-1 0 -1] [-4 0 -2]} {[-1 0 0] [-4 0 0]} {[-1 0 1] [-4 0 5]} {[-1 1 -1] [-4 1 -2]} {[-1 1 0] [-4 1 0]} {[-1 1 1] [-4 1 5]} {[0 -1 -1] [0 -6 -2]} {[0 -1 0] [0 -6 0]} {[0 -1 1] [0 -6 5]} {[0 0 -1] [0 0 -2]} {[0 0 1] [0 0 5]} {[0 1 -1] [0 1 -2]} {[0 1 0] [0 1 0]} {[0 1 1] [0 1 5]} {[1 -1 -1] [3 -6 -2]} {[1 -1 0] [3 -6 0]} {[1 -1 1] [3 -6 5]} {[1 0 -1] [3 0 -2]} {[1 0 0] [3 0 0]} {[1 0 1] [3 0 5]} {[1 1 -1] [3 1 -2]} {[1 1 0] [3 1 0]} {[1 1 1] [3 1 5]}]
```

Each element in the array is a tuple of vectors.  Take the first one for example:

```
[-1 -1 -1] [-4 -6 -2]
```

The first vector, `[-1 -1 -1]` says that the cubi is located at coordinates `x=-1`, `y=-1`, and `z=-1`.  The second vector says that the cube is colored as `Red` (number 4) on the left wall (negative on the `x`-axis), `Blue` (number 6) on the wall facing you (negative on the `y`-axis), and `White` (number 2) on the wall facing down (negative on the `z`-axis).

The array of vectors can be flattened out like this:

```
        y y y
        y y y
        y y y
r r r | g g g | o o o | b b b
r r r | g g g | o o o | b b b
r r r | g g g | o o o | b b b
        w w w
        w w w
        w w w
```

Some observations about colors:

- Corner cubis have all non-zero x,y,z components in the color vector
- In-between-corners cubis have two out of three non-zero components
- Center-of-face cubis have one non-zero component
- The cubi at the center of the cube would have had all components as zero in its color vector---but since the center cube has no colors, it doesn't need to be represented at all.


### Rotation

With the representation as described above, rotation becomes trivial.  We simply multiply vectors (position and color vectors) by a rotation matrix.  This is the function that performs a move of the cube.  You pass in the axis, an index in the axis, and a direction (counter-clockwise or clockwise).  It rotates the corresponding vectors via matrix multiplication, where `m` is a rotation matrix for a given axis and direction.

```
func (cb Cube) Rotate(a Axis, idx int, c bool) Cube {
  ret := cb.New()
  m := GetRotationMatrix(a, c)
  for cb_idx := range cb.cubis {
    if cb.cubis[cb_idx].pv[a] == idx {
      // Rotate
      ret.cubis[cb_idx] = m.Mult(cb.cubis[cb_idx])
    }
  }
  return ret
}
```
