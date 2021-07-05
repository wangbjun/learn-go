package main

import "fmt"

// 冒泡排序，前一个和后一个比较，然后进行交换
func main() {
	var arr = []int{1, 15, 2, 7, 3, 5, 10, 3, 3, 12}
	var length = len(arr)
	for i := 0; i < length; i++ {
		var flag = false
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	fmt.Printf("%v\n", arr)
}
