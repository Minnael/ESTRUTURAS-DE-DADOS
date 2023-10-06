package main

import "fmt"

type Stack[T any] interface {
	push(val T)  //TERMINADO
	top() T      //TERMINADO
	pop()        //TERMINADO
	empty() bool //TERMINADO
	clear()      //TERMINADO
	show()       //TERMINADO
}

type Node[T any] struct {
	value T
	back  *Node[T]
}

type LinkedStack[T any] struct {
	head     *Node[T]
	inserted int
}

func (stack *LinkedStack[T]) push(val T) { // O(1) BASTA ADICIONAR NO FINAL DA PILHA, Ω(1)
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

func (stack *LinkedStack[T]) top() T { // O(1) BASTA RETORNAR O ITEM DO TOPO, Ω(1)
	if stack.head == nil {
		fmt.Println("EMPTY STACK")
		var aux T
		return aux
	}
	return stack.head.value
}

func (stack *LinkedStack[T]) pop() { // O(1) BASTA RETIRAR O ÚLTIMO ITEM DA PILHA, Ω(1)
	if stack.head == nil {
		fmt.Println("EMPTY STACK")
		stack.inserted--
		return
	}
	stack.head = stack.head.back
	stack.inserted--
}

func (stack *LinkedStack[T]) empty() bool { // O(1) BASTA VERIFICAR STACK.HEAAD OU STACK.INSERTED, Ω(1)
	if stack.head == nil {
		fmt.Println("EMPTY STACK -> TRUE")
		return true
	} else {
		fmt.Println("EMPTY STACK -> FALSE")
		return false
	}
}

func (stack *LinkedStack[T]) clear() { // O(1) BASTA TRABALHAR STACK.HEAD, Ω(1)
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

func parenteses(par string) bool {
	stack := LinkedStack[string]{head: nil}
	size := len(par)
	for i := 0; i < size; i++ {
		if string(par[i]) == "(" {
			stack.push("(")
		}
		if string(par[i]) == ")" && stack.top() == "(" {
			stack.pop()
		}
	}
	if stack.empty() {
		return true
	} else {
		return false
	}
}

func main() {
	stack := LinkedStack[string]{head: nil}

	aux := "(())"

	parenteses(aux)
}
