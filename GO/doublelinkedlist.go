package main

import "fmt"

type List[T any] interface {
	append(val T) //
	show()        //
	pop()         //
	update(val T, pos int)
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
	if list.size == 0 {
		fmt.Println("EMPTY LIST")
		return
	}
	if list.size == 1 {
		list.head = nil
		list.tail = nil
		return
	}

	aux := list.head

	for aux.next.next != nil {
		aux = aux.next
	}
	aux.next = nil
	list.size--
}

func (list *DoubleLinkedList[T]) update(val T, pos int) {
	if pos < 0 || pos > list.size-1 {
		fmt.Println("INVALID POSITION")
		return
	}

	if pos == 0 {
		list.head.value = val
		return
	}

	aux := list.head
	for i := 0; i < pos-1; i++ {
		aux = aux.next
	}
	aux.next.value = val
}

func main() {
	list := DoubleLinkedList[int]{head: nil, tail: nil, size: 0}

	list.append(5)
	list.append(10)
	list.append(15)
	list.append(20)
	list.append(25)

	list.update(50, 5)
	list.show()
}
