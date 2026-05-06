package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Gin 提供了强大的绑定系统，可以将请求数据解析到 Go 结构体中并自动验证。
无需手动调用 c.PostForm() 或读取 c.Request.Body，只需定义一个带标签的结构体，让 Gin 来完成工作
*/

type LoginForm struct {
	User 		string `form:"user" binding:"required"`
	Password 	string `form:"password" binding:"required"`
}

func main() {
	route := gin.Default()

	route.POST("/login", func(ctx *gin.Context) {
		var form LoginForm

		if err := ctx.ShouldBind(&form); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "logged in"})
	})

	route.Run()
}