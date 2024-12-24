package main

import (
	"log"
	"net/rpc"
)

// 模拟RPC客户端
//同步调用
//func main() {
//	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
//	if err != nil {
//		log.Fatal("建立链接失败:", err)
//	}
//	var reply Result
//	err2 := client.Call("Cal.Square", 14, &reply)
//	if err2 != nil {
//		log.Fatal("RPC调用 Cal.Square 方法失败:", err2)
//	}
//	log.Printf("输入的值为%d", reply.Num)
//	log.Printf("计算的值为%d", reply.Ans)
//}

type Result struct {
	Num, Ans int
}

// 异步调用
func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("建立链接失败:", err)
	}
	var reply Result
	asyncCall := client.Go("Cal.Square", 14, &reply, nil)
	log.Printf("输入的值为%d,计算的结果为%d", reply.Num, reply.Ans)
	<-asyncCall.Done
	log.Printf("输入的值为%d,计算的结果为%d", reply.Num, reply.Ans)
}
