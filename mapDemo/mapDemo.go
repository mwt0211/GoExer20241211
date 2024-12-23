package main

import (
	"fmt"
	"sync"
)

func main() {
	var m map[string]int
	var str []map[string]interface{}
	m = map[string]int{"张三": 24, "咯i四": 25}
	//todo:需要完善key为String，value为任意类型的声明，并向map中添加具体测试值
	//str=append(str, map[string]interface{
	//	"name": "HPP",
	//	"age": 18,
	//})
	fmt.Println(str == nil)
	fmt.Println(m)
	fmt.Println("*******************")

	//sync.Map 不能使用 make 创建
	var scene sync.Map
	// 将键值对保存到sync.Map
	//sync.Map 将键和值以 interface{} 类型进行保存。
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	//遍历需要提供一个匿名函数，参数为 k、v，类型为 interface{}，每次 Range() 在遍历一个元素时，都会调用这个匿名函数把结果返回。
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}
