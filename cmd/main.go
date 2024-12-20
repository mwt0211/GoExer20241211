// package 定义包名 main 包名
package main

import (
	"fmt"
	"image/color"
	"mwt.com/util"
)

var pllette = []color.Color{color.White, color.Black}
var str3 string

// str1:=""
var str2 = ""

const (
	writeIndex = 0
	blackIndex = 1
	redIndex   = 2
)

// func 定义函数 main 函数名
func main() {

	util.PrinytStr()
	util.TestIota()
	cal := util.Cal(3.23, 1.05, '/')
	fmt.Printf("运算结果保留两位小数为：%.2f\n", cal)
	fmt.Println("运算结果为：", cal)
	fmt.Println("utils包中的Num变量为", util.Num)
	jugeString := util.JugeString("asdcfg", "adcfgs")
	fmt.Println("jugeString", jugeString)
}
