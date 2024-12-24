package main

import (
	"encoding/json"
	"fmt"
	"log"
	"unsafe"
)

func main() {
	var monster Monster
	var monsterdeliver Monster
	version := 1

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
	fmt.Println("--------------------")
	cmd := newCommand(
		"version",
		&version,
		"show version",
	)
	//cmd2 := newCommand(
	//	"version",
	//	&version,
	//	"show version",
	//)
	fmt.Println(cmd)
	//todo:输出值不为传入值
	fmt.Printf("%d", cmd.Var)
	fmt.Println("******************")
	cal := new(Cal)
	result := cal.getQuarter(3)
	fmt.Printf("%d^2的计算结果为:%d\n", result.Num, result.Ans)
	log.Printf("%d^2的计算结果为:%d\n", result.Num, result.Ans)

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

type Command struct {
	Name    string `json:"name"`    // 指令名称
	Var     *int   `json:"var"`     // 指令绑定的变量
	Comment string `json:"comment"` // 指令的注释
}

func newCommand(name string, varRef *int, comment string) *Command {
	return &Command{
		Name:    name,
		Var:     varRef,
		Comment: comment,
	}
}

type Result struct {
	Num, Ans int
}
type Cal int

func (cal *Cal) getQuarter(num int) *Result {
	return &Result{
		Num: num,
		Ans: num * num,
	}

}
