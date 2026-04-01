// https://go.dev/tour/moretypes/18

package main

import (
	"math/rand"

	"golang.org/x/tour/pic"
)

func one(x, y int) uint8 {
	return uint8((x + y) / 2)
}
func two(x, y int) uint8 {
	return uint8(x * y)
}
func three(x, y int) uint8 {
	return uint8(x ^ y)
}

var funcs = []func(int, int) uint8{
	one,
	two,
	three,
}

func Pic(dx, dy int) (pic [][]uint8) {
	pic = make([][]uint8, dy)
	f := funcs[rand.Intn(3)]

	for y := range pic {
		pic[y] = make([]uint8, dx)

		for x := range pic[y] {
			pic[y][x] = f(x, y)
		}
	}

	return
}

func main() {
	pic.Show(Pic)
}
