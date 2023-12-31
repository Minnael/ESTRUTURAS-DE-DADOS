package main

import "fmt"

func insertionSort(list []int) []int {
	for i := 0; i < len(list)-1; i++ {
		number := list[i+1]
		for i >= 0 && list[i] > number {
			list[i+1], list[i] = list[i], list[i+1]
			i--
		}
	}
	return list
}

func main() {
	list := []int{3, 2, 1}
	fmt.Println(insertionSort(list))
}
