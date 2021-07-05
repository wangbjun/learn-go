package main

import "fmt"

func main() {
	// var s [5]int
	// s[0] = 1
	// s[2] = 3
	// for k, v := range s {
	//	println(strconv.Itoa(k) + " ---> " + strconv.Itoa(v))
	//
	//	time.Sleep(1)
	//
	//	fmt.Printf("%d ---> %d\n", k, v)
	// }

	var m = make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	for k, v := range m {
		fmt.Printf("%s ---> %d\n", k, v)
	}

	mp := map[string]int{
		"e": 1,
		"f": 2,
	}

	// 删除元素
	delete(mp, "e")
	// 加减
	mp["f"] = mp["f"] + 1
	mp["f"]++

	for k, v := range mp {
		fmt.Printf("%s ---> %d\n", k, v)
	}
}
