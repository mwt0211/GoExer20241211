package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	////指定文件服务器的路径
	//fs := http.FileServer(http.Dir("F:\\课外学习\\Java面试资料合集"))
	//server := &http.Server{
	//	Addr:    ":8089",
	//	Handler: fs,
	//}
	//log.Fatal(server.ListenAndServe())
	r := gin.Default()
	r.Static("/download", "F:\\课外学习\\Java面试资料合集")
	err := r.Run(":8090")
	if err != nil {
		log.Fatal("failed to start server:", err)
	}
}
