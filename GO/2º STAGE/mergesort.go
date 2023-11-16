package main

import (
	"fmt"
)

func mergeSort(lista []int) []int {
	if len(lista) <= 1 {
		return lista
	}

	meio := len(lista) / 2
	esquerda := mergeSort(lista[0:meio])
	direita := mergeSort(lista[meio:len(lista)])

	return merge(esquerda, direita)
}

func merge(esquerda, direita []int) []int {
	resultado := make([]int, 0, len(esquerda)+len(direita))
	i, j := 0, 0

	for i < len(esquerda) && j < len(direita) {
		if esquerda[i] < direita[j] {
			resultado = append(resultado, esquerda[i])
			i++
		} else {
			resultado = append(resultado, direita[j])
			j++
		}
	}
	for i < len(esquerda) {
		resultado = append(resultado, esquerda[i])
		i++
	}

	for j < len(direita) {
		resultado = append(resultado, direita[j])
		j++
	}

	return resultado
}

func main() {
	lista := []int{4, 7, 2, 6, 4, 1, 8, 3}
	fmt.Println("Original:", lista)

	lista = mergeSort(lista)
	fmt.Println("Ordenado:", lista)
}
