package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
*

​	//捕获标准输入，并转换为字符串
​	reader := bufio.NewReader(os.Stdin)
​	input, err := reader.ReadString('\n')

​	if err != nil {

​    //如果有错误 退出

​		panic(err)
​	}

需求：能打怪升级
*/
var level = 1
var ex = 0

func main() {
	fmt.Println("请输入你的角色名字")
	//捕获标准输入，并转换为字符串
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	//去掉了最后的\n
	name := input[:len(input)-1]
	fmt.Printf("角色创建成功，当前的用户名为%s,用户等级为：%d\n", name, level)
	s := `你遇到了一个怪物，请选择是战斗还是逃跑?
	1.战斗
	2.逃跑`
	fmt.Printf("%s \n", s)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		selector := input[:len(input)-1]
		switch selector {
		case "1":
			ex += 10
			fmt.Println("杀死了某怪物，获取%d点经验值", ex)
			computeLevel()
			fmt.Println("当前等级为%d", level)
		case "2":
			fmt.Printf("你选择了逃跑\n")
			fmt.Printf("%s \n", s)
		case "exit":
			fmt.Printf("你选择了退出，期待下次参与\n")
			os.Exit(1)
		default:
			fmt.Printf("你的输入我不认识，请重新输入" +
				"\n")
		}

	}
}

/*
*
计算当前等级
*/
func computeLevel() {
	if ex < 20 {
		level = 1
	} else if ex < 40 {
		level = 2
	} else if ex < 200 {
		level = 3
	} else {
		level = 4
	}
}
