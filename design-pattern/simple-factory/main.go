package main

import "fmt"

// 定义工厂类
type factory struct {
}

// 定义产品接口
type product interface {
	info()
}

// 定义具体产品类，实现相关接口
type phone struct {
}

func (phone) info() {
	fmt.Println("info phone")
}

// 定义具体产品类，实现相关接口
type computer struct {
}

func (computer) info() {
	fmt.Println("info computer")
}

// 实现工厂方法
func (factory) createProduct(t string) product {
	switch t {
	case "phone":
		return &phone{}
	case "computer":
		return &computer{}
	default:
		return &phone{}
	}
}

// 简单工厂模式: 定义一个工厂类，它可以根据参数的不同返回不同类的实例，被创建的实例通常都具有共同的父类
func main() {
	factory := &factory{}
	phone := factory.createProduct("phone")
	phone.info()

	computer := factory.createProduct("computer")
	computer.info()
}
