package main

import "fmt"

var soma int

func mergeSort(lista []int) []int {
	if len(lista) > 1 {
		meio := len(lista) / 2
		esquerda := mergeSort(lista[0:meio])
		direita := mergeSort(lista[meio:len(lista)])
		return merge(lista, esquerda, direita)
	}
	return lista
}

func merge(lista []int, esquerda []int, direita []int) []int {
	i, j, lista_aux := 0, 0, make([]int, 0, len(esquerda)+len(direita))

	for i < len(esquerda) && j < len(direita) {
		if esquerda[i] < direita[j] {
			lista_aux = append(lista_aux, esquerda[i])
			i++
		} else {
			lista_aux = append(lista_aux, direita[j])
			j++
		}
	}
	lista_aux = append(lista_aux, esquerda[i:]...)
	lista_aux = append(lista_aux, direita[j:]...)

	for i := 0; i < len(lista_aux); i++ {
		lista[i] = lista_aux[i]
	}

	return lista_aux
}

func main() {
	lista := []int{5, 4, 3, 2, 1, -5, -4, -3, -2, -1}
	mergeSort(lista)
	fmt.Println(lista)
}
