package main

import (
	"fmt"
)

func main() {
	/* 指针 */
	a := 10
	// 获取指针地址&a
	b := &a
	// 获取指针的值*b
	c := *b
	fmt.Println(a, b, c)
	fmt.Printf("a的类型是%T,b的类型是%T,c的类型是%T\n", a, b, c)

	/* make和new */
	// 使用make创建切片、map和channel
	slice := make([]int, 3)   // 创建长度为3的切片
	m := make(map[string]int) // 创建map
	ch := make(chan int, 2)   // 创建带缓冲的channel
	fmt.Println(slice, m, ch)

	// 使用new创建各种类型
	ptr := new(int)   // 创建指向int零值的指针
	fmt.Println(*ptr) // 输出0

	// 对比make和new创建切片的区别
	s1 := make([]int, 3) // 创建初始化的切片，可直接使用
	s2 := new([]int)     // 创建指向nil切片的指针，*s2是空切片
	fmt.Println(s1, *s2) // 输出[0 0 0] []

	// *s2需要先append才能使用
	*s2 = append(*s2, 1, 2, 3)
	fmt.Println(*s2) // 输出[1 2 3]
}
