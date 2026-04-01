// https://go.dev/tour/flowcontrol/8

package main

import (
	"fmt"
	"math"
	"slices"
)

func SqrtCompareVals(x float64) (z float64) {
	last2Zs := make([]float64, 2)
	var prevZ float64

	fmt.Println("Compare Vals")
	z = 1.0
	for i := 1; ; i++ {
		prevZ = z

		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %v: %v\n", i, z)

		last2Zs[0] = last2Zs[1]
		last2Zs[1] = prevZ

		if slices.Contains(last2Zs, z) {
			return
		}
	}
}

func SqrtEpsilon(x float64) (z float64) {
	var prevZ float64

	fmt.Println("\nEpsilon")
	z = 1.0
	for i := 1; ; i++ {
		prevZ = z

		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %v: %v\n", i, z)

		if math.Abs(z-prevZ) < 1e-10 {
			return
		}
	}
}

func main() {
	x := 2.0
	fmt.Printf("\nCompare Vals Sqrt: %v\nEpsilon Sqrt: %v\nStandard Sqrt: %v\n", SqrtCompareVals(x), SqrtEpsilon(x), math.Sqrt(x))
}
