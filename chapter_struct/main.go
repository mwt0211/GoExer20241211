package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var monster Monster
	var monsterdeliver Monster

	monster.Name = "张三"
	monster.Age = 18

	monsterdeliver.Age = 19
	monsterdeliver.Name = "HPP"

	monster2 := monster
	monster3 := &monsterdeliver

	monster2.Name = "李四"
	monster3.Name = "王五"
	fmt.Println("monster2", monster2)             //monster2 {李四 18}
	fmt.Println("monster", monster)               //monster {张三 18}
	fmt.Println("monster3")                       //monster3 &{王五 19}
	fmt.Println("monsterdeliver", monsterdeliver) //monsterdeliver {王五 19}
	var a byte = 'a'
	c := a / 10
	fmt.Println("c", c)
	var bol = false
	fmt.Println("布尔类型的长度为", unsafe.Sizeof(bol), "个字节")
}

type Person struct {
	Name  string
	Age   int
	Score [5]float64
	ptr   *int
	slice []int
	map1  map[string]string
}

type Monster struct {
	Name string
	Age  int
}
