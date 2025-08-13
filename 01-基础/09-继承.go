package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// 继承
type Student struct {
	Person *Person
	School string
}

// 继承
type Teacher struct {
	Person Person
	School string
}

// 继承的时候指定结构体指针和结构体的区别。
// 继承结构体指针，在赋值操作的时候，可以修改原来的值，而继承结构体的时候，赋值会把结构体对象复制一份

func main() {
	stu := Student{
		Person: &Person{Name: "John", Age: 20},
		School: "MIT",
	}
	fmt.Println(stu)

	stu.Person.Age = 21
	fmt.Println(stu.Person.Age)

	tea := Teacher{
		Person: Person{Name: "Amy", Age: 30},
		School: "Harvard",
	}
	fmt.Println(tea)
	tea.Person.Age = 31
	fmt.Println(tea.Person.Age)

	// 指针类型可以修改原来的值
	tea1 := tea
	tea1.Person.Age = 32
	fmt.Println(tea1.Person.Age) // 32
	fmt.Println(tea.Person.Age)  // 31

	// 结构体类型会复制一份新的值
	stu1 := stu
	stu1.Person.Age = 22
	fmt.Println(stu1.Person.Age) // 22
	fmt.Println(stu.Person.Age)  // 22

}
