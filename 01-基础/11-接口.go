package main

import (
	"fmt"
)

// 定义一个接口 Animal，包含一个 Speak 方法
type Animal interface {
	Speak() string
}

// 定义一个函数，传入一个 Animal 类型的参数，并调用其 Speak 方法
func Speak(a Animal) {
	fmt.Println(a.Speak())
}

// 定义一个 Dog 结构体，包含一个 Name 字段
type Dog struct {
	Name string
}

// 实现 Animal 接口的 Speak 方法
func (d Dog) Speak() string {
	return "Woof!"
}

// 定义一个 Cat 结构体，包含一个 Name 字段
type Cat struct {
	Name string
}

// 实现 Animal 接口的 Speak 方法
func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	dog := Dog{Name: "Rex"}
	cat := Cat{Name: "Whiskers"}

	// Speak函数接收Amimal接口，不管哪个结构体，只要实现了 Animal 接口的 Speak 方法，就可以调用
	Speak(dog) // 输出：Woof!
	Speak(cat) // 输出：Meow!
}
