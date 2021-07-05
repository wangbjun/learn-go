package main

import "fmt"

func main() {
	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Println("请输入姓名 ")
	fmt.Scanln(&name)

	fmt.Println("请输入年龄 ")
	fmt.Scanln(&age)

	fmt.Println("请输入工资 ")
	fmt.Scanln(&sal)
	fmt.Println("是否通过考试 ")
	fmt.Scanln(&isPass)

	fmt.Printf("\n姓名是%v\n年龄是%v\n工资是%v\n能过考试了吗？%v\n", name, age, sal, isPass)

}
