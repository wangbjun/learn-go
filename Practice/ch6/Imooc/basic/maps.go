package main

import "fmt"

func main() {
	m := map[string]string{
		"name":     "abc",
		"location": "china",
	}

	fmt.Println(m)

	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println(m2)
	fmt.Println(m3)

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println(m["name"])
	fmt.Println(m["noExist"])
}
