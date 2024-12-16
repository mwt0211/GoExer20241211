package util

import "fmt"

func cal(a float64, b float64, operate byte) float64 {
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
