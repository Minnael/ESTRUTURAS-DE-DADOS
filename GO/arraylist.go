package main

import "fmt"

type List[T any] interface {
	append(value T)          //TERMINADO
	get(pos int) T           //TERMINADO
	show()                   //TERMINADO
	pop() T                  //TERMINADO
	update(value T, pos int) //TERMINADO
	insert(value T, pos int) //TERMINADO
	remove(pos int)          //TERMINADO
	reverse()                //TERMINADO
}

type ArrayList[T any] struct {
	values   []T
	inserted int
}

func (list *ArrayList[T]) doubleArray() {
	newArray := make([]T, len(list.values)*2)
	for i := 0; i < list.inserted; i++ {
		newArray[i] = list.values[i]
	}
	list.values = newArray
}

func (list *ArrayList[T]) append(val T) {
	if list.inserted == 0 {
		list.values[0] = val
		list.inserted++
	} else {
		if list.inserted == len(list.values) {
			list.doubleArray()
		}
		list.values[list.inserted] = val
		list.inserted++
	}
}

func (list *ArrayList[T]) pop() T {
	aux := list.values[list.inserted-1]
	list.inserted--
	return aux
}

func (list *ArrayList[T]) show() {
	for i := 0; i < list.inserted; i++ {
		fmt.Print("|", list.values[i], "|    ")
	}
}

func (list *ArrayList[T]) get(pos int) T {
	if list.inserted == 0 {
		fmt.Println("EMPTY LIST!")
		var aux T
		return aux
	}

	if pos >= 0 && pos <= list.inserted-1 {
		return list.values[pos]
	} else {
		fmt.Println("INVALID POSITION!")
		var aux T
		return aux
	}
}

func (list *ArrayList[T]) update(val T, pos int) {
	if pos >= 0 && pos <= list.inserted-1 {
		list.values[pos] = val
	} else {
		fmt.Println("INVALID POSITION!")
	}
}

func (list *ArrayList[T]) insert(val T, pos int) {
	if pos >= 0 && pos < list.inserted {
		if list.inserted == len(list.values) {
			list.doubleArray()
		}
		for i := list.inserted; i > pos; i-- {
			list.values[i] = list.values[i-1]
		}
		list.values[pos] = val
		list.inserted++
	} else {
		fmt.Println("INVALID POSITION!")
	}
}

func (list *ArrayList[T]) remove(pos int) {
	if pos < 0 || pos > list.inserted {
		fmt.Println("INVALID POSITION!")
	} else {
		for i := pos; i < list.inserted-1; i++ {
			list.values[i] = list.values[i+1]
		}
		list.inserted--
	}
}

func (list *ArrayList[T]) reverse() {
	for i := 0; i < list.inserted/2; i++ {
		aux := list.values[list.inserted-1-i]
		list.values[list.inserted-1-i] = list.values[i]
		list.values[i] = aux
	}
}

func main() {
	list := ArrayList[int]{values: make([]int, 5), inserted: 0}

	list.append(10)  //0
	list.append(20)  //1
	list.append(30)  //2
	list.append(40)  //3
	list.append(50)  //4
	list.append(60)  //5
	list.append(70)  //6
	list.append(80)  //7
	list.append(90)  //8
	list.append(100) //9

	fmt.Println(list.pop())

	list.reverse()

	list.show()
}
