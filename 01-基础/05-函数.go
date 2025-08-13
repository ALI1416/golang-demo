package main

import "fmt"

// 函数定义
func add(a int, b int) int {
	return a + b
}

// 多个返回值函数
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	result := add(2, 3)
	fmt.Println(result) // 5

	// 匿名函数
	addfunc := func(a int, b int) int {
		return a + b
	}

	result2 := addfunc(1, 2)
	fmt.Println(result2)

	a, b := swap("a", "b")
	fmt.Println(a, b)

}
