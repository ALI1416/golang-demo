package main

import "fmt"

func main() {
	// Map中的元素没有固定顺序
	// 传递Map时只复制引用，不复制数据
	// 键必须是可比较的类型（如数字、字符串、布尔等）
	// 值可以是任何类型

	// 创建Map的不同方式
	studentScores1 := map[string]int{}
	studentScores2 := make(map[string]int)
	studentScores := map[string]int{
		"张三": 85,
		"李四": 92,
		"王五": 78,
	}
	fmt.Println(studentScores1, studentScores2, studentScores)

	// 添加和修改元素
	studentScores["张三"] = 90
	studentScores["刘六"] = 60
	fmt.Println("学生成绩:", studentScores)

	// 获取元素
	zhangScore := studentScores["张三"]
	fmt.Println("张三的成绩:", zhangScore)

	// 检查键是否存在
	score, ok := studentScores["赵六"]
	if ok {
		fmt.Println("赵六的成绩:", score)
	} else {
		fmt.Println("赵六不在成绩单中")
	}

	// 删除元素
	delete(studentScores, "王五")
	fmt.Println("删除王五后:", studentScores)

	// 7. 遍历Map
	for name, score := range studentScores {
		fmt.Printf("%s: 成绩=%d\n", name, score)
	}

	//8. 只获取key
	for k := range studentScores {
		fmt.Printf("%s\n", k)
	}

	// 8. 获取Map长度
	fmt.Println("学生人数:", len(studentScores))

}
