package util

import (
	"fmt"
	"sort"
	"strings"
)

var Num string = "1"

func PrinytStr() {
	// fmt 包名 . 调用 Print 函数,并且输出定义的字符串
	age := 1
	fmt.Println("Hello Golang From mwt.com")
	fmt.Printf("%T", age)
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	//fmt.Println(x * y)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))
	fmt.Println(x * y)

}
func TestIota() {
	type weekday int
	const (
		sunday weekday = iota
		Monday
		Tune
		Wed
		Thur
		Fri
		Sat
	)
	day := [...]string{sunday: "星期日", Monday: "星期1", Tune: "星期2", Wed: "星期3", Thur: "星期4", Fri: "星期5", Sat: "星期6"}
	//day := [...]string{Monday: "星期1", Tune: "星期2", Wed: "星期3", Thur: "星期4", Fri: "星期5", Sat: "星期6"}
	fmt.Println(day[Tune])
}

/*
*
算术计算
*/
func Cal(a float64, b float64, operate byte) float64 {
	var res float64
	switch operate {
	case '+':
		res = a + b
		break
	case '-':
		res = a - b
		break
	case '*':
		res = a * b
		break
	case '/':
		res = a / b
		break
	default:
		fmt.Println("操作符号有错误")

	}
	return res

}

/*
*
判断两个字符串是否具有相同的字符，但是其顺序并不相同
*/
func JugeString(str1, str2 string) bool {
	var flag = false
	//判断两个字符串是否相同
	if str1 == str2 {
		flag = false
		return flag
	}
	//判断两个字符串长度是否相同
	if str1 == str2 {
		flag = false
		return flag
	}
	//判断两个字符串排序之后是否相同
	str1Chars := strings.Split(str1, "")
	str2Chars := strings.Split(str2, "")
	sort.Strings(str1Chars)
	sort.Strings(str2Chars)
	fmt.Println("str1Chars---", str1Chars)
	fmt.Println("str2Chars---", str2Chars)
	flag = strings.Join(str1Chars, "") == strings.Join(str2Chars, "")
	return flag
}
