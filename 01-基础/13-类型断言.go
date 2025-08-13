package main

import "fmt"

func main() {
	// x 是一个空接口，可以接受任何类型
	var x interface{} = "Hello, Go!"

	// 类型断言，检查 x 是否是一个 string 类型
	value, ok := x.(string)
	if ok {
		fmt.Println("x is a string:", value) // 输出：x is a string: Hello, Go!
	} else {
		fmt.Println("x is not a string")
	}

	// 类型断言，检查 x 是否是一个 int 类型
	value2, ok2 := x.(int)
	if ok2 {
		fmt.Println("x is an int:", value2)
	} else {
		fmt.Println("x is not an int") // 输出：x is not an int
	}

}
