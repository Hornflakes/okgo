// https://go.dev/tour/moretypes/23

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	res := make(map[string]int)

	fields := strings.Fields(s)
	for _, f := range fields {
		res[f]++
	}

	return res
}

func main() {
	wc.Test(WordCount)
}
