package main

import (
	"errors"
	"fmt"
)

type List[T any] interface {
	append(val T)           //TERMINADO
	get(pos int) (T, error) //TERMINADO
	show()                  //TERMINADO
	pop() (T, error)        //TERMINADO
	update(val T, pos int)  //TERMINADO
	insert(val T, pos int)  //TERMINADO
	remove(pos int)         //TERMINADO
	reverse()               //TERMINADO
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	head     *Node[T]
	inserted int
}

func (list *LinkedList[T]) append(val T) {
	newNode := Node[T]{value: val, next: nil}
	if list.head == nil {
		list.head = &newNode
		list.inserted++
		return
	}
	aux := list.head
	for aux.next != nil {
		aux = aux.next
	}
	aux.next = &newNode
	list.inserted++
}

func (list *LinkedList[T]) get(pos int) (T, error) {
	if list.head == nil {
		var aux T
		return aux, errors.New("EMPTY LIST")
	}
	if pos < 0 || pos > list.inserted-1 {
		var aux T
		return aux, errors.New("INVALID POSITION")
	}
	aux := list.head
	for i := 0; i < pos; i++ {
		aux = aux.next
	}
	return aux.value, errors.New("")
}

func (list *LinkedList[T]) show() {
	aux := list.head
	for i := 0; i < list.inserted; i++ {
		fmt.Print("|", aux.value, "|  ->  ")
		aux = aux.next
	}
	fmt.Print("nil")
}

func (list *LinkedList[T]) pop() (T, error) {
	if list.head == nil {
		var aux T
		return aux, errors.New("EMPTY LIST")
	}
	if list.head.next == nil {
		sup := list.head.value
		list.head = nil
		list.inserted--
		return sup, errors.New("")
	}
	aux := list.head
	for aux.next.next != nil {
		aux = aux.next
	}
	sup := aux.value
	aux.next = aux.next.next
	list.inserted--
	return sup, errors.New("")
}

func (list *LinkedList[T]) update(val T, pos int) {
	if list.head == nil {
		fmt.Println("EMPTY LIST")
		return
	}
	if pos < 0 || pos > list.inserted-1 {
		fmt.Println("INVALID POSITION")
		return
	}
	aux := list.head
	for i := 0; i < pos; i++ {
		aux = aux.next
	}
	aux.value = val
}

func (list *LinkedList[T]) insert(val T, pos int) {
	if pos < 0 || pos > list.inserted {
		fmt.Println("INVALID POSITION")
		return
	}
	newNode := Node[T]{value: val, next: nil}
	if pos == 0 {
		newNode.next = list.head
		list.head = &newNode
		list.inserted++
		return
	}
	aux := list.head
	for i := 1; i < pos; i++ {
		aux = aux.next
	}
	newNode.next = aux.next
	aux.next = &newNode
	list.inserted++
}

func (list *LinkedList[T]) remove(pos int) {
	if list.head == nil {
		fmt.Println("EMPTY LIST")
		return
	}
	if pos < 0 || pos > list.inserted-1 {
		fmt.Println("INVALID POSITION")
		return
	}
	if pos == 0 {
		list.head = list.head.next
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

func (list *LinkedList[T]) reverse() {
	current := list.head
	var prev, next *Node[T]
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	list.head = prev
}

func unique(list *LinkedList[int]) {
	if list.head == nil {
		fmt.Println("EMPTY LIST")
		return
	}

	aux := list.head

	for aux.next != nil {
		if aux.value == aux.next.value {
			aux.next = aux.next.next
			list.inserted--
		} else {
			aux = aux.next
		}
	}
}

func main() {
	list := LinkedList[int]{head: nil, inserted: 0}

	list.append(100)
	list.append(200)
	list.append(300)
	list.append(400)
	list.append(500)

	list.reverse()

	list.show()
}
