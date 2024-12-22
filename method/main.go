package main

import "fmt"

func main() {

	var person2 Person
	person2.Name = "张三"
	person2.Age = 28
	person2.sayHi2()
	person2.sayHi()
	person2.speak()

	var person3 Person
	person3.Name = "张三"
	person3.speak()
	person3.Cal()
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) sayHi() {
	p.Name = "HPP"
	p.Age = 18
	fmt.Println("你好啊!!", p.Age, "岁的", p.Name)
}

func (p Person) sayHi2() {
	p.Name = "HPP"
	p.Age = 18
	fmt.Println("你好啊!!", p.Age, "岁的", p.Name)
}

func (p *Person) speak() {
	fmt.Println(p.Name, "是一个大好人")
}

/*
*
计算10000以内的和
*/
func (p *Person) Cal() {
	sum := 0
	for i := 0; i <= 1000; i++ {
		sum += i
	}
	fmt.Println(p.Name, "计算出1000以内的和为：", sum)
}
