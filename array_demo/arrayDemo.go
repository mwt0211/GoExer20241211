package main

import "fmt"

func main() {
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[4:6] //相当于生成了一个新的切片
	//myslice := numbers4[8:10]//panic: runtime error: index out of range [3] with length 2
	//这打印出来长度为2
	fmt.Printf("myslice为 %d, 其长度为: %d,其cap为：%d\n", myslice, len(myslice), cap(myslice))
	myslice = myslice[:cap(myslice)]
	fmt.Printf("myslice为 %d, 其长度为: %d,其cap为：%d\n", myslice, len(myslice), cap(myslice))
	//为什么 myslice 的长度为2，却能访问到第四个元素
	fmt.Printf("myslice的第四个元素为: %d\n", myslice[3])

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	slice3 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Println(slice2)
	copy(slice1, slice3) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1)
}
