package main

import "fmt"

func countingSort(lista []int) {
	max := 0
	for i := 0; i < len(lista); i++ {
		if lista[i] > max {
			max = lista[i]
		}
	}

	A := make([]int, max+1)

	for i := range lista {
		A[lista[i]]++
	}

	index := 0
	for i, countValue := range A {
		for j := 0; j < countValue; j++ {
			lista[index] = i
			index++
		}
	}
}

func main() {
	lista := []int{2, 5, 3, 0, 2, 3, 0, 3}
	countingSort(lista)
	fmt.Println(lista)
}
