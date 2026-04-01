// https://go.dev/tour/methods/20

package main

import (
	"fmt"
	"math"
	"slices"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func SqrtCompareVals(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	last2Zs := make([]float64, 2)
	var prevZ float64

	fmt.Println("Compare Vals")
	z := 1.0
	for i := 1; ; i++ {
		prevZ = z

		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %v: %v\n", i, z)

		last2Zs[0] = last2Zs[1]
		last2Zs[1] = prevZ

		if slices.Contains(last2Zs, z) {
			return z, nil
		}
	}
}

func SqrtEpsilon(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	var prevZ float64

	fmt.Println("\nEpsilon")
	z := 1.0
	for i := 1; ; i++ {
		prevZ = z

		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %v: %v\n", i, z)

		if math.Abs(z-prevZ) < 1e-10 {
			return z, nil
		}
	}
}

func main() {
	x := -2.0

	z1, err1 := SqrtCompareVals(x)
	z2, err2 := SqrtEpsilon(x)

	fmt.Printf(
		"\nCompare Vals Sqrt: %v (err: %v)\nEpsilon Sqrt: %v (err: %v)\nStandard Sqrt: %v\n",
		z1, err1,
		z2, err2,
		math.Sqrt(x),
	)
}
