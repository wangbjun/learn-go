package main

import (
	"fmt"
	"imooc/queue"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(4)
	q.Push(3)

	fmt.Println(q)

	q.Push(5)

	fmt.Println(q)

	q.Pop()

	q.Push("ABC")

	q.Push("0.15")

	fmt.Println(q)
}
