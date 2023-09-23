package main

import (
	"errors"
	"fmt"
)

type Queue[T any] interface {
	Enqueue(value T)     //TERMINADO
	Dequeue() (T, error) //TERMINADO
	Front() (T, error)   //TERMINADO
	IsEmpty() bool       //TERMINADO
	Size() int           //TERMINADO
}

type ArrayQueue[T any] struct {
	values   []T
	inserted int
}

func (queue *ArrayQueue[T]) DoubleArray() {
	newArray := make([]T, queue.inserted*2)
	for i := 0; i < queue.inserted-1; i++ {
		newArray[i] = queue.values[i]
	}
	queue.values = newArray
}

func (queue *ArrayQueue[T]) Enqueue(val T) {
	if queue.inserted == len(queue.values) {
		queue.DoubleArray()
	}
	queue.values[queue.inserted] = val
	queue.inserted++
}

func (queue *ArrayQueue[T]) Dequeue() (T, error) {
	if queue.inserted == 0 {
		var aux T
		return aux, errors.New("EMPTY QUEUE")
	}
	aux := queue.values[0]
	queue.values = queue.values[1:]
	queue.inserted--
	return aux, errors.New("")
}

func (queue *ArrayQueue[T]) Front() (T, error) {
	if queue.inserted == 0 {
		var aux T
		return aux, errors.New("EMPTY QUEUE")
	}
	return queue.values[0], errors.New("")
}

func (queue *ArrayQueue[T]) IsEmpty() bool {
	if queue.inserted == 0 {
		return true
	}
	return false
}

func (queue *ArrayQueue[T]) Size() int {
	return queue.inserted
}

func main() {
	queue := ArrayQueue[int]{values: make([]int, 5), inserted: 0}

	queue.Enqueue(100)
	queue.Enqueue(200)
	queue.Enqueue(300)
	queue.Enqueue(400)
	queue.Enqueue(500)

	fmt.Println(queue.IsEmpty())
}
