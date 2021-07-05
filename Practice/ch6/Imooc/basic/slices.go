package main

import "fmt"

func updateSlices(s []int) {
	s[0] = 100
}

func printSlices(s []int) {
	fmt.Printf("slice len is %d, cap is %d\n", len(s), cap(s))
}

func main() {
	arr := [...]int{1, 2, 4, 5, 67, 8, 23, 3, 4}
	s := arr[2:6]

	fmt.Println(s)

	updateSlices(s) //slice是引用

	fmt.Println(s)
	fmt.Println(arr)

	s2 := s[1:3]

	s3 := s2[1:5] //向后扩展

	fmt.Println(s2)
	fmt.Println(s3)

	s3 = append(s2, 10)

	fmt.Println(s3)
	fmt.Println(arr)

	var s4 []int
	for i := 0; i < 50; i++ {
		printSlices(s4)
		s4 = append(s4, i*2+1)
	}
}
