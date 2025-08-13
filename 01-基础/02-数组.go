package main

import "fmt"

func main() {
	/* 定义、使用数组 */
	var numbers [3]int = [3]int{1, 2, 3}
	fmt.Println(numbers)    // [1 2 3]
	fmt.Println(numbers[1]) // 2

	// 数组在进行数据传递时，是值传递，而非引用传递
	var arr = [3]int{1, 2, 3}
	arr2 := arr
	arr2[0] = 3
	fmt.Println(arr, arr2) // [1 2 3] [3 2 3]

	/* 遍历数组 */
	scores := [5]int{95, 85, 75, 90, 88}
	// 使用传统的for循环
	for i := 0; i < len(scores); i++ {
		fmt.Printf("学生%d的成绩: %d\n", i+1, scores[i])
	}
	// 使用for range获取索引和值
	for index, score := range scores {
		fmt.Printf("学生%d的成绩: %d\n", index+1, score)
	}

	/* 声明切片 */
	var a []string            // nil
	var f = make([]int, 4)    // 初始化4个int默认值
	var b = []int{}           // 空
	var c = []int{1, 2, 3, 4} // 有值
	fmt.Println(a, f, b, c)
	// 添加元素
	c = append(c, 5)
	fmt.Println(c) // [1 2 3 4 5]
	// 切片操作
	// slice[low:high]
	// low: 起始索引（包含该元素）
	// high: 结束索引（不包含该元素）
	d := c[1:3]
	fmt.Println(d)
	fmt.Printf("长度:%d 容量:%d\n", len(d), cap(d)) // 5

}
