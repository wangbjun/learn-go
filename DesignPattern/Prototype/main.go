package main

import "fmt"

// 定义一个结构体（相当于面向对象语言里面的class）
type Person struct {
	name string
}

// 相当于class的方法
func (s Person) Say() {
	fmt.Printf(
		"%s: Say Hello\n", s.name)
}

// 此处克隆是复制了该对象的指针指向的内存区域
func (s *Person) Clone() Person {
	clone := *s
	return clone
}

// 原型模式是指用一个已经创建的实例作为原型，通过复制该原型对象来创建一个和原型相同或相似的新对象
// 像Java等其它面向对象的语言都提供了clone的方法，但是Go并没有这个方法，但是我们可以自己实现
func main() {
	person := Person{name: "Hi"}

	// clone该对象
	clone := person.Clone()
	clone.name = "Why"

	// 打印结果，克隆出来的对象并不影响之前的对象
	person.Say()
	clone.Say()
}
