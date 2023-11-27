package main

import "fmt"

func bubbleSort(list []int) {
	condition := true
	for condition {
		condition = false
		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				condition = true
			}
		}
	}
}

func main() {
	list := []int{10, 9, 8, 11, 2, 1, 8, -1, 14, 1, 20}
	bubbleSort(list)
	fmt.Println(list)
}
