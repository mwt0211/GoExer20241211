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

	//类型转换
	//高精度像低精度转换
	var i int64 = 9999999999 //1001010100000010111110001111111111
	b := int8(i)             //11111111
	fmt.Println("高精度像低精度转换的数值为", b)
	//基本类型转string
	iString := fmt.Sprintf("%d", i)
	fmt.Printf("%T,%q\n", iString, iString)
	judgeLevel(0)
}

// 根据分数来匹配对应的等级
func judgeLevel(score float64) string {
	res := ""
	switch {
	case score > 100:
		fmt.Printf(
			"特别棒\n")
	case score > 90:
		fmt.Printf(
			"优秀\n")
	case score > 80:
		fmt.Printf(
			"良好\n")
	case score > 70:
		fmt.Printf(
			"中等\n")
	case score > 60:
		fmt.Printf(
			"及格\n")
	default:
		fmt.Printf(
			"不及格\n")

	}
	return res
}
