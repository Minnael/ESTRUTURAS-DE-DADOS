package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil}
}

func (list *LinkedList) size() int {
	aux := list.head
	size := 0
	for aux != nil {
		size++
		aux = aux.next
	}
	return size
}

func (list *LinkedList) append(val int) {
	newNode := Node{value: val, next: nil}
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

func (list *LinkedList) insert(val int, pos int) {
	size := list.size()
	newNode := Node{value: val, next: nil}

	if pos < 0 || pos > size {
		fmt.Println("POSIÇÃO INVÁLIDA!")
	}

	aux := list.head
	for i := 0; i < pos-1; i++ {
		aux = aux.next
	}

	newNode.next = aux.next
	aux.next = &newNode
	fmt.Println(aux.next.value)
}

func (list *LinkedList) pop() {
	if list.head == nil {
		fmt.Println("LISTA VAZIA!")
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

func (list *LinkedList) get(pos int) int {
	if pos < 0 || pos > list.size() {
		fmt.Println("POSIÇÃO INVÁLIDA!")
		var none int
		return none
	}

	aux := list.head
	for i := 0; i < pos; i++ {
		aux = aux.next
	}

	if aux == nil {
		fmt.Println("POSIÇÃO INVÁLIDA!")
		var none int
		return none
	}
	return aux.value
}

func (list *LinkedList) remove(pos int) {
	if pos < 0 || pos > list.size() {
		fmt.Println("POSIÇÃO INVÁLIDA!")
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

func (list *LinkedList) update(val int, pos int) {
	if pos < 0 || pos > list.size() {
		fmt.Println("POSIÇÃO INVÁLIDA!")
	}
	aux := list.head
	if pos == 0 {
		aux.value = val
		return
	}
	for i := 0; i < pos-1; i++ {
		aux = aux.next
	}
	aux.next.value = val
}

func (list *LinkedList) inverter() {
	aux := list.head
	sup := list.head.next
	for i := 0; i < list.size()-1; i++ {
		aux = sup.next
	}
}

func (list *LinkedList) show() {
	size := list.size()
	for i := 0; i < size; i++ {
		fmt.Print("  ", list.get(i))
	}
}

func main() {
	lista := NewLinkedList()
	lista.append(189)  //0
	lista.append(29)   //1
	lista.append(19)   //2
	lista.append(109)  //3
	lista.append(24)   //4
	lista.append(1009) //5
	lista.append(37)   //6

	lista.inverter()
}
