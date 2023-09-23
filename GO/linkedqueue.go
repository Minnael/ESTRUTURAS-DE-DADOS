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

func (queue *LinkedQueue[T]) enqueue(val T) {
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

func (queue *LinkedQueue[T]) dequeue() (T, error) {
	if queue.inserted == 0 {
		var sup T
		return sup, errors.New("EMPTY QUEUE")
	}

	aux := queue.head
	queue.head = queue.head.next
	queue.inserted--
	return aux.value, errors.New("")
}

func (queue *LinkedQueue[T]) front() (T, error) {
	if queue.inserted == 0 {
		var sup T
		return sup, errors.New("EMPTY QUEUE")
	}

	return queue.head.value, errors.New("")
}

func (queue *LinkedQueue[T]) isempty() bool {
	if queue.inserted == 0 {
		return true
	} else {
		return false
	}
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

	queue.dequeue()
	queue.dequeue()
	queue.dequeue()

	queue.show()
}
