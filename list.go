// https://go.dev/tour/generics/2

package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T comparable] struct {
	next *List[T]
	val  T
}

func New[T comparable](v T) *List[T] {
	return &List[T]{
		val:  v,
		next: nil,
	}
}

func (l *List[T]) Push(v T) *List[T] {
	newNode := &List[T]{val: v}
	newNode.next = l
	return newNode
}

func (l *List[T]) Find(target T) bool {
	curr := l
	for curr != nil {
		if curr.val == target {
			return true
		}
		curr = curr.next
	}
	return false
}

func main() {
	list := New(30)
	list = list.Push(20)
	list = list.Push(10)

	fmt.Println("List head value:", list.val)
	fmt.Println("Second value:", list.next.val)

	found := list.Find(20)

	fmt.Println("Found 20:", found)
}
