package main

import (
	"errors"
	"fmt"
)

type List[T any] interface {
	append(val T)           //TERMINADO
	show()                  //TERMINADO
	pop() (T, error)        //TERMINADO
	update(val T, pos int)  //TERMINADO
	get(pos int) (T, error) //TERMINADO
	remove(pos int)         //TERMINADO
	insert(val T, pos int)  //TERMINADO
	reverse()               //TERMINADO
}

type Node[T any] struct {
	value      T
	prev, next *Node[T]
}

type DoubleLinkedList[T any] struct {
	head, tail *Node[T]
	inserted   int
}

func (list *DoubleLinkedList[T]) append(val T) {
	newNode := Node[T]{value: val, prev: list.tail, next: nil}
	if list.inserted == 0 {
		list.head = &newNode
		list.tail = &newNode
		list.inserted++
		return
	}
	list.tail.next = &newNode
	list.tail = &newNode
	list.inserted++
}

func (list *DoubleLinkedList[T]) show() {
	aux := list.head
	for aux != nil {
		fmt.Print("|", aux.value, "|  ->  ")
		aux = aux.next
	}
	fmt.Print("nil")
}

func (list *DoubleLinkedList[T]) pop() (T, error) {
	if list.inserted == 0 {
		var aux T
		return aux, errors.New("EMPTY LIST")
	}
	if list.inserted == 1 {
		sup := list.head.value
		list.head = nil
		list.tail = nil
		list.inserted--
		return sup, errors.New("")
	}
	sup := list.tail.value
	list.tail.prev.next = nil
	list.tail = list.tail.prev
	list.inserted--
	return sup, errors.New("")
}

func (list *DoubleLinkedList[T]) update(val T, pos int) {
	if pos < 0 || pos > list.inserted-1 {
		if pos == 0 {
			fmt.Println("EMPTY LIST")
		} else {
			fmt.Println("INVALID POSITION")
		}
		return
	}
	aux := list.head
	for i := 0; i < pos; i++ {
		aux = aux.next
	}
	aux.value = val
}

func (list *DoubleLinkedList[T]) get(pos int) (T, error) {
	if pos < 0 || pos > list.inserted-1 {
		var aux T
		if pos == 0 {
			return aux, errors.New("EMPTY LIST")
		} else {
			return aux, errors.New("INVALID LIST")
		}
	}
	aux := list.head
	for i := 0; i < pos; i++ {
		aux = aux.next
	}
	return aux.value, errors.New("")
}

func (list *DoubleLinkedList[T]) remove(pos int) {
	if pos < 0 || pos > list.inserted-1 {
		if pos == 0 {
			fmt.Println("EMPTY LIST")
		} else {
			fmt.Println("INVALID POSITION")
		}
		return
	}
	if pos == 0 {
		list.head = list.head.next
		list.inserted--
		return
	}
	if pos == list.inserted-1 {
		list.tail = list.tail.prev
		list.tail.next = nil
		list.inserted--
		return
	}
	aux := list.head
	for i := 1; i < pos; i++ {
		aux = aux.next
	}
	aux.next = aux.next.next
	list.inserted--
}

func (list *DoubleLinkedList[T]) insert(val T, pos int) {
	newNode := Node[T]{value: val, prev: nil, next: nil}
	if pos < 0 || pos > list.inserted {
		if list.inserted == 0 {
			list.head = &newNode
			list.tail = &newNode
			list.inserted++
		} else {
			fmt.Println("INVALID POSITION")
		}
		return
	}
	if pos == 0 {
		list.head.prev = &newNode
		newNode.next = list.head
		list.head = &newNode
	} else if list.inserted == pos {
		newNode.prev = list.tail
		list.tail.next = &newNode
		list.tail = &newNode
	} else {
		aux := list.head
		for i := 1; i < pos; i++ {
			aux = aux.next
		}
		newNode.next = aux.next
		newNode.prev = aux
		aux.next.prev = &newNode
		aux.next = &newNode
	}
	list.inserted++
}

func (list *DoubleLinkedList[T]) reverse() {
	current := list.head
	prev := list.head.prev
	next := list.head.next
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	list.head = prev
}

func main() {
	list := DoubleLinkedList[int]{head: nil, tail: nil, inserted: 0}

	list.append(100)
	list.append(200)
	list.append(300)
	list.append(400)
	list.append(500)

	list.insert(600, 5)
	list.reverse()

	list.show()
}
