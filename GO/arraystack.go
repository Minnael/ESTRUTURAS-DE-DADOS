package main

import "fmt"

type Array[T any] interface {
	push(val T) //TERMINADO
	pop() T
	top() T
	empty() bool
	clear()
	show() //TERMINADO
	balparenteses(par string) bool
}

type ArrayStack[T any] struct {
	values   []T
	inserted int
}

func (stack *ArrayStack[T]) DoublyArray() {
	newArray := make([]T, 2*stack.inserted)
	for i := 0; i < stack.inserted; i++ {
		newArray[i] = stack.values[i]
	}
	stack.values = newArray
}

func (stack *ArrayStack[T]) push(val T) {
	if len(stack.values) == stack.inserted {
		stack.DoublyArray()
	}
	stack.values[stack.inserted] = val
	stack.inserted++
}

func (stack *ArrayStack[T]) pop() T {
	aux := stack.values[stack.inserted-1]
	stack.inserted--
	return aux
}

func (stack *ArrayStack[T]) show() {
	for i := 0; i < stack.inserted; i++ {
		fmt.Print("|", stack.values[i], "|  ->  ")
	}
	fmt.Print("nil")
}

func parenteses(par string) bool {
	stack := ArrayStack[string]{values: make([]string, len(par)), inserted: 0}
	for i := 0; i < len(par); i++ {
		if string(par[i]) == "(" {
			stack.push("(")
		} else if string(par[i]) == ")" {
			if stack.top() == "(" {
				stack.pop()
			} else {
				return false
			}
		}
	}
	if stack.inserted == 0 {
		return true
	} else {
		return false
	}
}

func (stack *ArrayStack[T]) top() T {
	if stack.inserted == 0 {
		var aux T
		return aux
	}
	return stack.values[stack.inserted-1]
}

func main() {
	par := "(()))"

	fmt.Println(parenteses(par))

}
