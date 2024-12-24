package main

import (
	"log"
	"net/http"
	"net/rpc"
)

// 模拟RPC服务端
func main() {
	rpc.Register(new(Cal))
	rpc.HandleHTTP()
	log.Printf("Serving RPC server on port %d", 1234)
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("Error serving: ", err)
	}

}

type Result struct {
	Num, Ans int
}
type Cal int

//RPC 需要满足的条件
// func (t *T) MethodName(argType T1, replyType *T2) error
/*
即需要满足以下 5 个条件：

方法类型（T）是导出的（首字母大写）
方法名（MethodName）是导出的
方法有2个参数(argType T1, replyType *T2)，均为导出/内置类型
方法的第2个参数一个指针(replyType *T2)
方法的返回值类型是 error
*/
func (cal *Cal) Square(num int, res *Result) error {
	res.Num = num
	res.Ans = num * num
	log.Printf("服务端接收输入的值为%d", res.Num)
	log.Printf("服务端计算的值为%d", res.Ans)
	return nil
}
