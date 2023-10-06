package main

import (
	"errors"
	"fmt"
)

func buscabinaria(lista []int, valor int) (int, error) {
	esquerda := 0
	direita := len(lista) - 1
	for esquerda <= direita {
		meio := (esquerda + direita) / 2
		if lista[meio] == valor {
			return meio, errors.New("")
		} else if lista[meio] > valor {
			esquerda = meio + 1
		} else {
			direita = meio - 1
		}
	}
	return 0, errors.New("ITEM NÃO ENCONTRADO!")
}

func main() {
	lista := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(buscabinaria(lista, 4))
}
