package main

import "fmt"

func insertionSort(list []int) []int {
	for i := 0; i < len(list)-1; i++ {
		number := list[i+1]
		for i >= 0 && list[i] > number {
			list[i+1] = list[i]
			i = i - 1
		}
		list[i+1] = number
	}
	return list
}

func main() {
	list := []int{7, 5, 8, 3, 40, 10, 27, 13, 0, 20, 4, 1, 9, 14, 101, 28381, 31, 77, 91, 82, 97, 444, 900, 7}

	fmt.Println(insertionSort(list))
}
