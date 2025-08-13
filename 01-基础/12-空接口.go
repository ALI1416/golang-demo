package main

import (
	"fmt"
)

// 定义一个函数接收空接口(泛型)
func print(a interface{}) {
	fmt.Println(a)
}

func main() {
	print(1)
	print("hello")
	print(true)

	// 定义一个空接口类型的切片
	a := []interface{}{"nihao", 2, true}
	print(a)

	// 定义一个map，key为string，value为空接口
	b := map[string]interface{}{"name": "张三", "age": 20, "gender": "男"}
	print(b)

	// 定义一个结构体，包含一个字段，类型为空接口
	c := struct {
		Name interface{}
	}{Name: "张三"}
	print(c)

}
