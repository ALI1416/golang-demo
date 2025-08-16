package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	/* 变量 */
	// 显式声明
	var x int = 10
	var name string = "Golang"
	// 类型推导
	var y = 3.14 // float64
	// 短变量声明（仅在函数内使用）
	z := true
	fmt.Println(x, name, y, z)

	/* 常量 */
	const Pi = 3.14159
	const AppName = "MyApp"
	fmt.Println(Pi, AppName)

	/* 基本数据类型 默认值 */
	var defInt int         // 0
	var defFloat64 float64 // 0
	var defString string   // 空字符串
	var defBool bool       // false
	fmt.Println(defInt, defFloat64, defString, defBool)

	/* 格式化打印 */
	// %v: 默认格式值
	// %T: 变量类型
	// %d: 整数
	// %f: 浮点数
	// %t: 布尔值
	// %s: 字符串
	// %x, %X: 十六进制表示
	age := 25       // int
	price := 19.999 // float64
	// %T 数据类型 %v值
	fmt.Printf("age: %T %v \n", age, age)
	fmt.Printf("price: %T %v \n", price, price)
	// 控制小数位数
	fmt.Printf("2位小数: %.2f 4位小数: %.4f\n", price, price)
	// 多行文本
	message := `
	Hello
	Hello
	Hello
	`
	fmt.Println(message)

	/* 字符串操作 */
	s := "你好李四"
	// 将字符串转为uint8(byte)类型(ASCII字符)
	s_uint8 := []uint8(s)
	// 将字符串转为rune类型(UTF-8字符)
	s_rune := []rune(s)
	fmt.Println(s, s_uint8, s_rune)
	// 字符串长度(uint8)、UTF-8字符长度
	fmt.Println(len(s), utf8.RuneCountInString(s))
	// 字符串拼接
	ss := "张三"
	fmt.Println(s + ", " + ss)
	fmt.Println(fmt.Sprintf("%s, %s", s, ss))
	// 字符串分割
	n := "123-456-789"
	// []string类型
	arr := strings.Split(n, "-")
	fmt.Println(arr)

	/* 遍历字符串 */
	str := "Hello, 世界"
	// 使用for range遍历（推荐，可以正确处理Unicode字符）
	for index, char := range str {
		fmt.Printf("位置 %d: 字符 %c, Unicode码点 %U\n", index, char, char)
	}
	// 使用for循环遍历字节
	for i := 0; i < len(str); i++ {
		fmt.Printf("位置 %d: 字节 %d, 字符 %c\n", i, str[i], str[i])
	}
	// 将字符串转换为rune切片后遍历
	runes := []rune(str)
	for i, r := range runes {
		fmt.Printf("位置 %d: 字符 %c\n", i, r)
	}

}
