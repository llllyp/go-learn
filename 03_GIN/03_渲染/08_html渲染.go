package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 使用 html/template 包进行 HTML 渲染
func main() {

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	// 在不同的模板中使用同名模板模板
	

	router.Run()
}