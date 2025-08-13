package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// 指定序列化后的字段
type Person2 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 序列化，对象转json
	person := Person{Name: "John", Age: 30}
	jsonData, err := json.Marshal(person) // 返回的是一个字节切片[]uint8
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Printf("jsonData: %T\n", jsonData) // jsonData: []uint8
	// string() 将字节切片转换为字符串
	fmt.Println(string(jsonData)) // {"Name":"John","Age":30}

	// 反序列化,json转对象
	// 反引号（`）包围的字符串是 原始字符串字面量，
	// 与普通的双引号字符串（" "）不同，
	// 它不会对字符串内容中的特殊字符（如换行符、制表符等）进行转义处理，
	// 也不需要使用反斜杠（\）来转义
	json_str := `{"Name":"John","Age":30}`
	var person2 Person
	// json_str 是字符串类型,需要转换为字节切片
	err = json.Unmarshal([]byte(json_str), &person2)
	if err != nil {
		fmt.Println("Error unmarshalling from JSON:", err)
		return
	}
	fmt.Println(person2)

	// 序列化，对象转json，指定序列化后的字段
	person3 := Person2{Name: "Amy", Age: 30}
	jsonData, err = json.Marshal(person3)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData)) // {"name":"John","age":30}

	// 反序列化，json转对象，指定序列化后的字段
	json_str = `{"name":"Amy","age":30}`
	var person4 Person2
	err = json.Unmarshal([]byte(json_str), &person4)
	if err != nil {
		fmt.Println("Error unmarshalling from JSON:", err)
		return
	}
	fmt.Println(person4) // {John 30}
}
