package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/local/file", func(ctx *gin.Context) {
		ctx.File("local/file.go")
	})

	var fs http.FileSystem = http.Dir("/Users/lyp/temp/test")
	router.GET("/fs/file", func(ctx *gin.Context) {
		ctx.FileFromFS("file.go", fs)
	})

	router.GET("/download", func(ctx *gin.Context) {
		ctx.FileAttachment("local/aaa.txt", "bbb.txt")
	})

	router.Run()

	/*
	curl -v http://localhost:8080/download

	curl http://localhost:8080/local/file
	
	*/

}