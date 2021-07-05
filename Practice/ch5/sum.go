package main

import "fmt"

/**
可变函数
*/
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}

	return total
}

func max(vals ...int) int {
	if len(vals) <= 0 {
		return 0
	}
	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}

	return max
}

func main() {
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(4, 5, 6))
	fmt.Println(max(4, 15, 6))
	fmt.Println(max())
}
