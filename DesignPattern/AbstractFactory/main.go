package main

import "fmt"

// 抽象工厂模式: 是一种为访问类提供一个创建一组相关或相互依赖对象的接口，且访问类无须指定所要产品的具体类就能得到同族的不同等级的产品的模式结构
func main() {
	var f factoryInterface
	f = new(factory)
	phone := f.createPhone()
	computer := f.createComputer()
	phone.info()
	computer.info()
}

// 工厂接口
type factoryInterface interface {
	createPhone() productInterface
	createComputer() productInterface
}

// 产品接口
type productInterface interface {
	info()
}

type factory struct {
}

func (factory) createPhone() productInterface {
	return &phone{}
}
func (factory) createComputer() productInterface {
	return &computer{}
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
