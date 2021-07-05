package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randNew()
	return
	rand.Seed(time.Now().UnixNano())
	// 下面这些高阶函数都是使用的全局种子源
	fmt.Printf("%v\n", rand.Int63())     // 产生一个0到2的63次方之间的正整数,类型是int64
	fmt.Printf("%v\n", rand.Int63n(100)) // 产生一个0到n之间的正整数,n不能大于2的63次方，类型是int64
	fmt.Printf("%v\n", rand.Int31())     // 这个和前面的区别就产生的是int32类型的整数
	fmt.Printf("%v\n", rand.Int31n(100)) // 同上

	// 这个比较有意思了，它至少产生一个32位的正整数，因为int类型在64位机器上等于int64，在32位机器上就是int32
	fmt.Printf("%v\n", rand.Int())
	fmt.Printf("%v\n", rand.Intn(100)) // 同上

	fmt.Printf("%v\n", rand.Float32()) // 产生一个0到1.0的浮点数，float32类型
	fmt.Printf("%v\n", rand.Float64()) // 产生一个0到1.0的浮点数，float64类型
}

func randNew() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Printf("%v\n", r.Intn(100))
	fmt.Printf("%v\n", r.Intn(100))
	fmt.Printf("%v\n", r.Intn(100))
	fmt.Printf("%v\n", r.Intn(100))
	fmt.Printf("%v\n", r.Intn(100))

	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		m[r.Intn(1000)]++
	}

	var count = 0
	for k, v := range m {
		if v >= 2 {
			fmt.Printf("%v --- %v\n", k, v)
			count++
		}
	}
	println(count)
}
