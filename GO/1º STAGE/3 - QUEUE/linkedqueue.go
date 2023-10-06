package main

import (
	"errors"
	"fmt"
)

type Queue[T any] interface {
	enqueue(val T)
	dequeue() (T, error)
	front() (T, error)
	isempty() bool
	size() int
	show()
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedQueue[T any] struct {
	head, rear *Node[T]
	inserted   int
}

func (queue *LinkedQueue[T]) enqueue(val T) { // O(1) BASTA FAZER AJUSTES NOS PONTEIROS COM ENFASE NO REAR, Ω(1)
	newNode := Node[T]{value: val, next: nil}
	if queue.inserted == 0 {
		queue.head = &newNode
		queue.rear = &newNode
		queue.inserted++
		return
	}
	queue.rear.next = &newNode
	queue.rear = &newNode
	queue.inserted++
}

func (queue *LinkedQueue[T]) dequeue() (T, error) { // O(1) BASTA FAZER AJUSTE DE PONTEIROS EM HEAD, Ω(1)
	if queue.inserted == 0 {
		var sup T
		return sup, errors.New("EMPTY QUEUE")
	}

	aux := queue.head
	queue.head = queue.head.next
	queue.inserted--
	return aux.value, errors.New("")
}

func (queue *LinkedQueue[T]) front() (T, error) { // O(1) BASTA RETONAR O PRIMEIRO ITEM DA QUEUE, Ω(1)
	if queue.inserted == 0 {
		var sup T
		return sup, errors.New("EMPTY QUEUE")
	}

	return queue.head.value, errors.New("")
}

func (queue *LinkedQueue[T]) isempty() bool { // O(1) BASTA RETORNAR VERIFICAR QUEUE.INSERTED, Ω(1)
	if queue.inserted == 0 {
		return true
	} else {
		return false
	}
}

func (queue *LinkedQueue[T]) size() int { // O(1) BASTA RETORNAR QUEUE.INSERTED, Ω(1)
	return queue.inserted
}

func (queue *LinkedQueue[T]) show() {
	aux := queue.head
	for aux != nil {
		fmt.Print("|", aux.value, "| -> ")
		aux = aux.next
	}
	fmt.Print("nil")

}

func main() {
	queue := LinkedQueue[int]{head: nil, inserted: 0}

	queue.enqueue(10)
	queue.enqueue(20)
	queue.enqueue(30)
	queue.enqueue(40)
	queue.enqueue(50)

	fmt.Println(queue.rear.value)
}
