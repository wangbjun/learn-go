package main

import "fmt"

// 工厂方法模式: 定义一个创建产品对象的工厂接口，将产品对象的实际创建工作推迟到具体子工厂类当中
func main() {
	var f factoryInterface
	f = new(productFactory)
	phone := f.createProduct("phone")
	phone.info()
}

// 工厂接口
type factoryInterface interface {
	createProduct(string) productInterface
}

// 产品接口
type productInterface interface {
	info()
}

type productFactory struct {
}

// 实现工厂方法
func (productFactory) createProduct(t string) productInterface {
	switch t {
	case "phone":
		return &phone{}
	case "computer":
		return &computer{}
	default:
		return &phone{}
	}
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
