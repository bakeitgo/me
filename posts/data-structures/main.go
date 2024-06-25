package main

import (
	bubble_sort "data_structures/sort/bubble"
	"fmt"
)
func seedArray(size int) (arr []int) {
	for i := 0; i < size; i++ {
		arr = append(arr, i)
	}
	return
}

func main() {
	arr := []int{10,3,5,7,2,4,6}
	ar := bubble_sort.BubbleImplISaw(arr)
	for _, v := range ar {
		fmt.Println(v)
	}
}
