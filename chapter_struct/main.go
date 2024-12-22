package main

import (
	"encoding/json"
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

	//序列化
	jsonMonster := Monster{"牛魔", 100000, "铁扇公主的芭蕉扇"}
	marshal, _ := json.Marshal(jsonMonster)
	fmt.Println("marshal", string(marshal))
	//反序列化
	var jsonMonster2 Monster
	err := json.Unmarshal(marshal, &jsonMonster2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jsonMonster2.Skill, jsonMonster2.Age, jsonMonster2.Name)

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
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}
