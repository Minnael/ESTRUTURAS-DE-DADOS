package main

import "fmt"

type Stack[T any] interface {
	push(val T) //TERMINADO
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

type LinkedStack[T any] struct {
	head     *Node[T]
	inserted int
}

func (stack *LinkedStack[T]) push(val T) {
	newNode := Node[T]{value: val, back: nil}
	if stack.head == nil {
		stack.head = &newNode
		stack.inserted++
		return
	}

	aux := stack.head

	stack.head = &newNode
	stack.head.back = aux
	stack.inserted++
}

func (stack *LinkedStack[T]) top() T {
	if stack.head == nil {
		fmt.Println("EMPTY STACK")
		var aux T
		return aux
	}
	return stack.head.value
}

func (stack *LinkedStack[T]) pop() {
	if stack.head == nil {
		fmt.Println("EMPTY STACK")
		stack.inserted--
		return
	}
	stack.head = stack.head.back
	stack.inserted--
}

func (stack *LinkedStack[T]) empty() int {
	if stack.head == nil {
		fmt.Println("EMPTY STACK -> TRUE")
		return 1
	} else {
		fmt.Println("EMPTY STACK -> FALSE")
		return 0
	}
}

func (stack *LinkedStack[T]) clear() {
	stack.head = nil
}

func (stack *LinkedStack[T]) show() {
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
	stack := LinkedStack[int]{head: nil}

	stack.push(10)
	stack.push(20)
	stack.push(30)
	stack.push(40)
	stack.push(50)

	stack.show()
}
