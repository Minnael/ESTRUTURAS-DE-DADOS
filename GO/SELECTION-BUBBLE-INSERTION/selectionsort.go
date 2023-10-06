package main

import "fmt"

func selectionSort(lista []int) []int {
	for i := 0; i < len(lista); i++ {
		number := lista[i]
		sup := i
		for j := i; j < len(lista); j++ {
			if number >= lista[j] {
				number = lista[j]
				sup = j
			}
		}
		number = lista[i]
		lista[i] = lista[sup]
		lista[sup] = number
	}
	return lista
}

func main() {
	lista := []int{7, 5, 1, 8, 3, 40, 10, 1, 1, 27, 13, 9, 0}

	fmt.Println(selectionSort(lista))
}
