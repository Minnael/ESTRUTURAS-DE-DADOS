package main

import "fmt"

type List[T any] interface {
	append(val T)
	show()
	pop()
}

type Node[T any] struct {
	value      T
	prev, next *Node[T]
}

type DoubleLinkedList[T any] struct {
	head, tail *Node[T]
	size       int
}

func (list *DoubleLinkedList[T]) append(val T) {
	newNode := Node[T]{value: val, prev: nil, next: nil}
	if list.size == 0 {
		list.head = &newNode
		list.tail = &newNode
		list.size++
	} else {
		list.tail.next = &newNode
		list.tail = &newNode
		list.size++
	}
}

func (list *DoubleLinkedList[T]) show() {
	aux := list.head
	for aux != nil {
		fmt.Print("|", aux.value, "|   ")
		aux = aux.next
	}
}

func (list *DoubleLinkedList[T]) pop() {

}

func main() {
	list := DoubleLinkedList[int]{head: nil, tail: nil, size: 0}

	list.append(5)
	list.append(10)
	list.append(15)
	list.append(20)
	list.append(25)
	list.append(30)
	list.append(35)

	list.show()
}
