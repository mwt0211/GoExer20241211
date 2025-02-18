package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"strings"
)

type ErrorMessage struct {
	Code     int         `json:"code"`
	Reason   string      `json:"reason"`
	Message  string      `json:"message"`
	Metadata interface{} `json:"metadata"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	log.Println("实例化成功")
	fmt.Println("实例化成功")

	url := ""
	r.GET("/waiwang/download/:flag", func(c *gin.Context) {
		flag := c.DefaultQuery("flag", "API")
		fmt.Println(flag)
		var name = ""
		switch flag {
		case "API":
			url = "https://dl.google.com/go/go1.23.4.src.tar.gz"
			index := strings.LastIndex(url, "/")
			name = url[index+1 : len(url)]
			fmt.Println(name)
			//break
		case "FILE":
			url = ""
		case "ZIP":
			url = ""
		case "DOC":
			url = ""
		case "xls":
			url = ""
		default:
			url = ""
		}
		request, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
		if err != nil {
			log.Fatal("create request error:", err)
		}
		response, err := http.DefaultClient.Do(request)
		//client := http.DefaultClient
		//respBytes, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal("request error:", err)
		}
		if err != nil {
			log.Println("request failed:", err)
		}

		defer response.Body.Close()
		//判断响应体状态
		if response.StatusCode != http.StatusOK {
			log.Printf("status code error: %d %s", response.StatusCode, response.Status)
		}
		//设置响应头
		//c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Type", response.Header.Get("Content-Type"))
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name))
		c.Writer.WriteHeader(http.StatusOK)
		_, err = io.Copy(c.Writer, response.Body)
		if err != nil {
			log.Fatal("copy error:", err)
		}

	})
	err := r.Run("127.0.0.1:8089")
	if err != nil {
		log.Fatal("start server error:", err)
		return
	}
	log.Println("start server Suc:")
	fmt.Println("start server Suc:")

}
