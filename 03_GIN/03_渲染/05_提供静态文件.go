package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router.Static(relativePath, root)
	//提供整个目录 对 relativePath 的请求会映射到 root 下的文件
	router.Static("assets", "./assets") 

	// router.StaticFS(relativePath, fs) — 类似于 Static，但接受一个 http.FileSystem 接口，
	// 当你需要从嵌入式文件系统提供文件或自定义目录列表行为时使用。
	router.StaticFS("/more_static", http.Dir("my_file_system"))

	// router.StaticFile(relativePath, filePath) — 提供单个文件。适用于 /favicon.ico 或 /robots.txt 等端点。
	router.StaticFile("/favicon,ico", "./resources/favicon.ico")

	router.Run()
}