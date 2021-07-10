package main

import "fmt"

// 选择排序
func main() {
	var arr = []int{1, 15, 2, 7, 3, 5, 10, 3, 3, 12}
	var length = len(arr)

	for i := 0; i < length; i++ {
		min := arr[i]
		for j := i + 1; j < length; j++ {
			if arr[j] < min {
				min = arr[j]
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}

	fmt.Printf("%v\n", arr)
}
