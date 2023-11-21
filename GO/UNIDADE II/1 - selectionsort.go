package main

import "fmt"

func selectionSort(list []int) []int {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return list
}

func main() {
	list := []int{7, 5, 8, 3, 40, 10, 27, 13, 0, 20, 4, 1, 9, 14, 101, 28381, 31, 77, 91, 82, 97, 444, 900, 7}

	fmt.Println(selectionSort(list))
}
