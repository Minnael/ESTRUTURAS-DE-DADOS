package main

import "fmt"

type Queue[T any] interface {
	enqueue(val T)
	dequeue() T
	front() T
	isempty() bool
	show()
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedQueue[T any] struct {
	head     *Node[T]
	inserted inta
}

func (queue *LinkedQueue[T]) enqueue(val T) {
	newNode := Node[T]{value: val, next: nil}
	if queue.inserted == 0 {
		queue.head = &newNode
		queue.inserted++
		return
	}

	aux := queue.head

	for aux.next != nil {
		aux = aux.next
	}
	aux.next = &newNode
	queue.inserted++
}

func (queue *LinkedQueue[T]) dequeue() T {
	if queue.inserted == 0 {
		fmt.Println("DEQUEU -> EMPTY QUEUE!")
		var sup T
		return sup
	}

	aux := queue.head
	queue.head = queue.head.next
	queue.inserted--
	return aux.value
}

func (queue *LinkedQueue[T]) front() T {
	if queue.inserted == 0 {
		fmt.Println("FRONT -> EMPTY QUEUE!")
		var sup T
		return sup
	}

	return queue.head.value
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

}

func main() {
	queue := LinkedQueue[int]{head: nil, inserted: 0}

	queue.enqueue(10)
	queue.enqueue(20)
	queue.enqueue(30)
	queue.enqueue(40)
	queue.enqueue(50)

	fmt.Println(queue.isempty())

	queue.show()
}
