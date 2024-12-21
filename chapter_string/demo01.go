package main

import "fmt"

func main() {
	//字符串的转义字符
	str1 := "测试\n测试"
	fmt.Println("str1", str1)
	//使用反引号
	str2 := `type Person struct {
	Name  string
	Age   int
	Score [5]float64
	ptr   *int
	slice []int
	map1  map[string]string
}`
	fmt.Println("str2", str2)
	//字符串拼接
	str3 := "hello" + "啊"
	str3 += "HPP"
	fmt.Println("str3", str3)
	//长字符串拼接,涉及到换行时，拼接符号在上
	str4 := "hello" + "啊" + "hello" + "啊" + "hello" + "啊" +
		"hello" + "啊"
	fmt.Println("str4", str4)
}
