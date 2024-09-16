package main

import "fmt"

func reverseArray(arr []int) []int {
	/*
		make : built-in function
		- primary used for memory allocation (slices, maps, channels)
	*/
	reversed := make([]int, len(arr))
	for i := range arr {
		reversed[i] = arr[len(arr) - 1 - i] // this is cool please check it once
	}
	return reversed
}

func main() {
	arr := []int{1, 3, 2, 6, -1, 4, 1, 8, 2}
	result := reverseArray(arr)
	fmt.Println(result)
}
