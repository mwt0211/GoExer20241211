package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	url := "http://localhost:8090/download/111.pdf"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatal("create request failed:", err)
	}
	httpClient := http.DefaultClient
	res, err := httpClient.Do(req)
	if err != nil {
		log.Println("request failed:", err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("status code error: %d %s", res.StatusCode, res.Status)
	}
	//写入新文件
	fileName := "F:\\课外学习\\Java面试资料合集\\newFile.pdf"
	file1, err := os.Create(fileName)
	if err != nil {
		log.Fatal("create file failed:", err)
	}
	defer file1.Close()
	_, err = io.Copy(file1, res.Body)
	if err != nil {
		log.Fatal("copy file failed:", err)
	}

	log.Println("copied file Suc:", fileName)

}
