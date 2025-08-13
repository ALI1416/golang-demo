package main

import "fmt"

func main() {
	/* if-else */
	score := 85
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 60 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	/* for */ // 无while
	// 标准for循环
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	// 类似while的循环
	sum := 0
	for sum < 5 {
		fmt.Println(sum)
		sum++
	}
	// 遍历切片
	numbers := []int{1, 2, 3}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	/* switch-case */ // 自动break
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Other")
	}

	/* 没有三元表达式 */

}
