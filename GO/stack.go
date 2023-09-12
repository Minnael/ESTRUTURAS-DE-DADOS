package main

import "fmt"

type Stac[T any] interface {
	push(val T) //TERMINADO
	size() int  //TERMINADO
	top() T     //TERMINADO
	pop()       //TERMINADO
	empty() int //TERMINADO
	clear()     //TERMINADO
	show()      //TERMINADO
}

type Node[T any] struct {
	value T
	back  *Node[T]
}

type Stack[T any] struct {
	head *Node[T]
}

func (stack *Stack[T]) push(val T) {
	newNode := Node[T]{value: val, back: nil}
	if stack.head == nil {
		stack.head = &newNode
		return
	}

	aux := stack.head

	stack.head = &newNode
	stack.head.back = aux
}

func (stack *Stack[T]) size() int {
	sup := 0
	aux := stack.head
	for aux != nil {
		sup++
		aux = aux.back
	}
	return sup
}

func (stack *Stack[T]) top() T {
	if stack.head == nil {
		fmt.Println("EMPTY STACK")
		var aux T
		return aux
	}
	return stack.head.value
}

func (stack *Stack[T]) pop() {
	if stack.head == nil {
		fmt.Println("EMPTY STACK")
		return
	}
	stack.head = stack.head.back
}

func (stack *Stack[T]) empty() int {
	if stack.head == nil {
		fmt.Println("EMPTY STACK -> TRUE")
		return 1
	} else {
		fmt.Println("EMPTY STACK -> FALSE")
		return 0
	}
}

func (stack *Stack[T]) clear() {
	aux := stack.head
	for aux != nil {
		aux = aux.back
		stack.head = aux
	}
}

func (stack *Stack[T]) show() {
	if stack.head == nil {
		fmt.Println("EMPTY STACK")
		return
	}

	aux := stack.head

	for aux.back != nil {
		fmt.Println("|", aux.value, "|")
		aux = aux.back
	}
	fmt.Println("|", aux.value, "|")
}

func main() {
	stack := Stack[int]{head: nil}

	stack.push(10)
	stack.push(50)
	stack.push(40)
	stack.push(20)

	stack.empty()

	stack.clear()

	stack.empty()

}
