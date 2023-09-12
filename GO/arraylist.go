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

func (list *ArrayList) doubleArray(sizeArray int) {
	doubleValues := make([]int, 2*sizeArray)
	for i := 0; i < sizeArray; i++ {
		doubleValues[i] = list.values[i]
	}
	list.values = doubleValues
}

func (list *ArrayList) append(value int) {
	sizeArray := len(list.values)
	if list.inserted == sizeArray {
		list.doubleArray(sizeArray)
	}
	list.values[list.inserted] = value
	list.inserted++
}

func (list *ArrayList) show(pos int) int {
	if pos < 0 || pos > list.inserted {
		panic("POSICAO INVÁLIDA!")
	} else {
		return list.values[pos]
	}
}

func (list *ArrayList) insert(value int, pos int) {
	if pos >= 0 && pos <= list.inserted {
		sizeArray := len(list.values)
		if list.inserted == sizeArray {
			list.doubleArray(sizeArray)
		}
		for i := list.inserted + 1; i > pos; i-- {
			list.values[i] = list.values[i-1]
		}
		list.values[pos] = value
		list.inserted++
	}
}

func (list *ArrayList) pop() {
	if list.inserted == 0 {
		fmt.Println("LISTA SEM NENHUM ELEMENTO!")
	} else {
		var aux int
		list.values[list.inserted-1] = aux
		list.inserted--
	}
}

func (list *ArrayList) update(value int, pos int) {
	if pos >= 0 && pos <= list.inserted {
		list.values[pos] = value
	} else {
		fmt.Println("POSIÇÃO INVÁLIDA!")
	}
}

func (list *ArrayList) remove(pos int) {
	if pos >= 0 && pos <= list.inserted {
		for i := pos; i < list.inserted-1; i++ {
			list.values[i] = list.values[i+1]
		}
		var aux int
		list.values[list.inserted-1] = aux
		list.inserted--
	} else {
		fmt.Println("POSIÇÃO INVÁLIDA!")
	}
}

func (list *ArrayList) reverse() {
	for i := 0; i < list.inserted/2; i++ {
		list.values[i], list.values[list.inserted-i-1] = list.values[list.inserted-i-1], list.values[i]
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
	list.append(6)

	//list.insert(50, 4)

	list.reverse()

	fmt.Println("TAMANHO DA LISTA:", list.size())
	for i := 0; i < list.inserted; i++ {
		fmt.Println("POSICAO:", list.show(i))
	}
}
