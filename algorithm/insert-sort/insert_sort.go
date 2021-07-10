package main

import "fmt"

// 插入排序
func main() {
	var arr = []int{17, 15, 2, 7, 3, 5, 10, 3, 3, 12}
	var length = len(arr)

	for i := 1; i < length; i++ {
		value := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > value {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = value
	}

	fmt.Printf("%v\n", arr)
}
