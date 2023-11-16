package main

import "fmt"

func insertionSort(list []int) []int {
	for i := 0; i < len(list)-1; i++ {
		number := list[i+1]
		for i >= 0 && list[i] > number {
			list[i+1] = list[i]
			i--
		}
		list[i+1] = number
	}
	return list
}

func main() {
	list := []int{7, 5}

	fmt.Println(insertionSort(list))
}
