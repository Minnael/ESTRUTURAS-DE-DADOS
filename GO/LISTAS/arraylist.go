package main

import "fmt"

type ArrayList struct {
	values   []int
	inserted int
}

func Init(size int) ArrayList {
	return ArrayList{values: make([]int, size), inserted: 0}
}

func (list *ArrayList) size() int {
	return list.inserted
}

func (list *ArrayList) append(value int) {
	sizeArray := len(list.values)
	if list.inserted == sizeArray {
		doubleValues := make([]int, 2*sizeArray)
		for i := 0; i < sizeArray; i++ {
			doubleValues[i] = list.values[i]
		}
		list.values = doubleValues
	}
	list.values[list.inserted] = value
	list.inserted++
}

func (list *ArrayList) show(pos int) int {
	if pos < 0 || pos > list.inserted {
		panic("POSICAO INVÁLIDA")
	} else {
		return list.values[pos]
	}
}

func main() {
	list := Init(5)

	list.append(0)
	list.append(1)
	list.append(2)
	list.append(3)
	list.append(4)
	list.append(5)

	fmt.Println("TAMANHO DA LISTA:", list.size())
	fmt.Println("POSICAO 0:", list.show(1))
}
