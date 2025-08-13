package main

import (
	"fmt"
	"net/http"
)

// 案例: 空指针引用
func demoNilPointer() {
	fmt.Println("------ 空指针引用案例 ------")
	// defer关键字：
	// defer会将后面的函数调用推迟到当前函数返回之前执行
	// 无论当前函数是正常返回还是因为panic而中断，defer都会确保这个函数被调用
	// 匿名函数：func(){}()
	// 是一个立即定义并执行的匿名函数。括号()表示立即调用这个函数
	defer func() {
		// recover()函数：
		// 这是Go语言内置函数，用于捕获当前goroutine中的panic
		// 如果当前goroutine没有panic，recover()返回nil；
		// 如果有panic，recover()会捕获panic的值并返回，同时停止panic传播
		// 判断逻辑：if r := recover(); r != nil {}
		// 会先调用recover()并将结果赋值给变量r，然后检查r是否为nil。
		// 如果不是nil，说明捕获到了panic
		if r := recover(); r != nil {
			fmt.Println("捕获到panic:", r)
		}
	}()
	var p *int = nil
	fmt.Println("尝试解引用空指针")
	fmt.Println(*p) // 这里会引发panic
}

func main() {
	/* error */
	// 大部分的内置包或者外部包，都有自己的报错处理机制
	// 因此我们使用的任何函数可能报错，这些报错都不应该被忽略
	// 而是在调用函数的地方，优雅地处理报错
	resp, err := http.Get("https://web.404z.cn/api/")
	if err != nil {
		fmt.Println("错误：" + err.Error())
	} else {
		fmt.Println(resp)
	}

	/* panic/recover */
	// 用于处理严重错误（如程序崩溃），不推荐频繁使用
	demoNilPointer()
	fmt.Printf("hello world")
}
