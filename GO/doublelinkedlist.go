package main

import "fmt"

type List[T any] interface {
	append(value T)
}

type Node[T any] struct {
	value      int
	prev, next *Node[T]
}

type DoubleLinkedList[T any] struct {
	head, tail *Node[T]
}

func main() {
	fmt.Println("Hello World")
}
