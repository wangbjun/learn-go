package main

import "fmt"

func main() {
	p := people{name: "who"}
	run(&p)

	d := dog{name: "旺财"}
	run(&d)
}

func run(d duck) {
	fmt.Println(d.say())
}

type duck interface {
	say() string
}

type people struct {
	name string
}

func (*people) say() string {
	return "哈哈"
}

type dog struct {
	name string
}

func (*dog) say() string {
	return "汪汪"
}
