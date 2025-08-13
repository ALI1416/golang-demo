package main

import "fmt"

/* 结构体 */
type Person struct {
	Name string
	Age  int
}

/* 结构体方法 */
func (p Person) SayHello() {
	fmt.Println("Hello, my name is", p.Name)
}

/* 结构体指针方法 */
func (p *Person) SetAge1(age int) {
	p.Age = age
}

/* 结构体方法 */
func (p Person) SetAge2(age int) {
	p.Age = age
}

// main 函数演示了Go语言中结构体的四种初始化方式：
// 1. 直接初始化结构体
// 2. 先声明后赋值
// 3. 使用new关键字创建指针
// 4. 使用取地址符&初始化指针
// 同时展示了结构体方法和指针方法的调用区别，
// 以及Go语言按值传递的特性。
func main() {
	/* 直接初始化结构体 */
	person := Person{Name: "John", Age: 30}

	/* 先声明后赋值 */
	var person2 Person
	person2.Name = "Amy"
	person2.Age = 25
	fmt.Printf("person2: %T\n", person2)

	/* 使用new关键字创建指针 */
	// person3返回的是结构体指针
	// person3.Name = "xiaoming" 底层 (*person3).name = "xiaoming"
	person3 := new(Person)
	person3.Name = "xiaoming"
	person3.Age = 30
	fmt.Printf("person3: %T\n", person3)

	/* 使用取地址符&初始化指针 */
	person4 := &Person{Name: "liuqiang", Age: 30}
	fmt.Printf("person4: %T\n", person4)

	// 调用构体方法
	person.SayHello()

	// 调用结构体指针方法
	// 指针接收者：
	// 当结构体方法使用指针接收者（即方法接收者为*StructType形式）定义时，该方法会接收一个指向结构体的指针
	// 这意味着方法内部操作的是原始结构体的内存地址，而不是其副本
	// 因此，对结构体字段的任何修改都会直接反映到原始结构体上
	person4.SetAge1(40)
	fmt.Println(person4.Age) // 40

	// 在Go语言中，函数参数的传递方式是按值传递
	// 这意味着当你将一个变量传递给函数时，实际上是传递了这个变量的副本
	// 因此，如果在函数内部修改了这个副本，原始变量不会受到影响
	person.SetAge2(40)
	fmt.Println(person.Age) // 30

}
