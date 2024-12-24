package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	//http://localhost:9090/
	engine.GET("/", func(c *gin.Context) {
		c.String(200, "hello world\n")
		//status状态位
		c.String(http.StatusOK, "who are you?")

	})
	//解析路径参数
	//http://localhost:9090/user/HPP
	engine.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello %s\n", name)
	})
	//获取Query参数
	//http://localhost:9090/users?name=HPP&age=20
	engine.GET("users", func(c *gin.Context) {
		name := c.Query("name")
		age := c.DefaultQuery("age", "18")
		c.String(http.StatusOK, "hello 我的名字是%s,我的年龄是 %s\n", name, age)
	})
	//Post请求
	//todo:未绑定成功
	engine.POST("/form", func(c *gin.Context) {
		userName := c.PostForm("username")

		pwd := c.DefaultPostForm("password", "1111111")
		fmt.Println(pwd)
		c.JSON(http.StatusOK, gin.H{
			"userName": userName,
			"passWord": pwd,
		})
	})
	//重定向
	engine.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})
	engine.GET("/index", func(c *gin.Context) {
		c.Request.URL.Path = "/users"
		engine.HandleContext(c)
	})
	engine.Run(":9090")

}
