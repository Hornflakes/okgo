package main

import (
	"golang.org/x/tour/tree"

	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func PrintWalk(t *tree.Tree, name string) {
	ch := make(chan int)

	go func() {
		Walk(t, ch)
		close(ch)
	}()

	fmt.Printf("PrintWalk %v: ", name)
	for v := range ch {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 != ok2 {
			return false
		}

		if !ok1 {
			return true
		}

		if v1 != v2 {
			return false
		}
	}

	return true
}

func main() {
	tree1 := &tree.Tree{
		Value: 3,
		Left: &tree.Tree{
			Value: 1,
			Left: &tree.Tree{
				Value: 1,
			},
			Right: &tree.Tree{
				Value: 2,
			},
		},
		Right: &tree.Tree{
			Value: 8,
			Left: &tree.Tree{
				Value: 5,
			},
			Right: &tree.Tree{
				Value: 13,
			},
		},
	}
	tree2 := &tree.Tree{
		Value: 8,
		Left: &tree.Tree{
			Value: 3,
			Left: &tree.Tree{
				Value: 1,
				Left: &tree.Tree{
					Value: 1,
				},
				Right: &tree.Tree{
					Value: 2,
				},
			},
			Right: &tree.Tree{
				Value: 5,
			},
		},
		Right: &tree.Tree{
			Value: 13,
		},
	}

	fmt.Println("Println tree1:", tree1)
	fmt.Println("Println tree2:", tree2)

	PrintWalk(tree1, "tree1")
	PrintWalk(tree2, "tree2")

	same := Same(tree1, tree2)
	fmt.Println("Same:", same)
}
