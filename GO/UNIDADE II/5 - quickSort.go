package main

import "fmt"

func quickSort(lista []int, inicio int, fim int) {
	if inicio < fim {
		pivo := particiona(lista, inicio, fim)
		quickSort(lista, inicio, pivo-1)
		quickSort(lista, pivo+1, fim)
	}
}

func particiona(lista []int, inicio int, fim int) int {
	pivo := lista[fim]
	i := inicio
	for j := inicio; j < fim; j++ {
		if lista[j] < pivo {
			lista[j], lista[i] = lista[i], lista[j]
			i++
		}
	}
	lista[fim], lista[i] = lista[i], lista[fim]
	return i
}

func main() {
	lista := []int{4, 7, 2, 6, 4, 1, 8, 3}
	quickSort(lista, 0, len(lista)-1)
	fmt.Println(lista)
}
