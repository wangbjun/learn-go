package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 演示懒汉式
	go func() {
		s := NewLazySingleton()
		s.Say()
	}()
	singleton := NewLazySingleton()
	singleton.Say()

	// 演示饿汉式
	go func() {
		h := NewHungrySingleton()
		h.Say()
	}()
	hungry := NewHungrySingleton()
	hungry.Say()

	time.Sleep(time.Second)
}

// 定义一个结构体（相当于面向对象语言里面的class）
type Singleton struct {
	name string
}

// 相当于class的方法
func (s Singleton) Say() {
	fmt.Printf("%s: Say Hello\n", s.name)
}

// 由于Go并不是完全面向对象的语言，只能通过包级别的全局变量模拟类变量
var lazyInstance *Singleton

// 为了Go协程并发安全，可以加锁，如果不加锁第一次实例化的时候会出现data race操作,但是其实并不影响功能
// 第一加锁方式是使用mutex
// mu.lock
// defer mu.unlock
// 但是这样并不优雅，相当于每次New的时候都得加锁，影响性能，所以最好的方式是使用once
var once sync.Once

// 懒汉式，所谓懒汉式就是比较懒，等你第一次调用的时候才会去实例化对象
func NewLazySingleton() *Singleton {
	once.Do(func() {
		fmt.Println("init lazy")
		lazyInstance = &Singleton{name: "lazy"}
	})
	return lazyInstance
}

var hungryInstance = Singleton{name: "hungry"}

// 饿汉式，所谓饿汉式就是在你调用New之前就已经初始化好了，这个更简单，而且也不需要锁
func NewHungrySingleton() *Singleton {
	return &hungryInstance
}
