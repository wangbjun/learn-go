package main

import "fmt"

func printArray(arr [5]int) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	var arr [5]int
	arr[1] = 5

	arr2 := [3]int{1, 2, 3}

	arr3 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)

	//for i:= 0;i<len(arr3);i++{
	//	fmt.Println(arr3[i])
	//}

	for _, v := range arr3 {
		fmt.Println(v)
	}

	printArray(arr)
	//printArray(arr2) //error
}
