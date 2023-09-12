package main

import "fmt"

type List[T any] interface {
	append(val T)          //TERMINADO
	size() int             //TERMINADO
	get(pos int) T         //TERMINADO
	show()                 //TERMINADO
	pop()                  //TERMINADO
	update(val T, pos int) //TERMINADO
	insert(val T, pos int) //TERMINADO
	remove(pos int)        //TERMINADO
	reverse()              //TERMINADO
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
}

func (list *LinkedList[T]) append(val T) {
	newNode := Node[T]{value: val, next: nil}
	if list.head == nil {
		list.head = &newNode
		return
	}

	aux := list.head

	for aux.next != nil {
		aux = aux.next
	}
	aux.next = &newNode
}

func (list *LinkedList[T]) size() int {
	var size int
	aux := list.head
	for aux != nil {
		aux = aux.next
		size++
	}
	return size
}

func (list *LinkedList[T]) get(pos int) T {
	size := list.size() - 1
	if pos < 0 || pos > size {
		fmt.Println("INVALID POSITION")
		var sup T
		return sup
	}
	if list.head == nil {
		fmt.Println("EMPTY LIST")
		var sup T
		return sup
	}
	if pos == 0 {
		return list.head.value
	}

	aux := list.head
	for aux.next != nil {
		aux = aux.next
	}
	return aux.value
}

func (list *LinkedList[T]) show() {
	size := list.size()
	aux := list.head
	for i := 0; i < size; i++ {
		fmt.Print("|", aux.value, "|  ->  ")
		aux = aux.next
	}
	print("nil")
}

func (list *LinkedList[T]) pop() {
	if list.head == nil {
		fmt.Println("EMPTY LIST")
		return
	}
	if list.head.next == nil {
		list.head = nil
		return
	}
	aux := list.head
	for aux.next.next != nil {
		aux = aux.next
	}
	aux.next = nil
}

func (list *LinkedList[T]) update(val T, pos int) {
	if pos < 0 || pos > list.size()-1 {
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
	if pos < 0 || pos > list.size()-1 {
		fmt.Println("INVALID POSITION")
		return
	}

	newNode := Node[T]{value: val, next: nil}
	if pos == 0 {
		sup := list.head
		list.head = &newNode
		newNode.next = sup
		return
	}

	aux := list.head

	for i := 0; i < pos-1; i++ {
		aux = aux.next
	}
	sup := aux.next
	aux.next = &newNode
	newNode.next = sup
}

func (list *LinkedList[T]) remove(pos int) {
	if pos < 0 || pos > list.size()-1 {
		fmt.Println("INVALID POSITION")
		return
	}
	if pos == 0 {
		list.head = list.head.next
		return
	}

	aux := list.head

	for i := 0; i < pos-1; i++ {
		aux = aux.next
	}
	aux.next = aux.next.next
}

func (list *LinkedList[T]) reverse() {
	current := list.head
	var prev *Node[T]
	var next *Node[T]
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	list.head = prev
}

func main() {
	list := LinkedList[int]{head: nil}

	list.append(5)
	list.append(10)
	list.append(15)
	list.append(20)
	list.append(25)

	list.reverse() //INVERTE LISTA
	list.show()    //EXIBIR LISTA
}
