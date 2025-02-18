package main

import "fmt"

/*
*
defer 执行顺序
*/
func main() {
	defer fmt.Println("1号门")
	devide(20)
	defer fmt.Println("2号门")
	devide(0)
	defer fmt.Println("3号门")
}
func devide(x float64) float64 {
	return 100 / x
}
