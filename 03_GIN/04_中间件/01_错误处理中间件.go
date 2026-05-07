package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 集中式的错误处理中间件通过在每个请求后运行并检查通过 c.Error(err) 添加到 Gin 上下文中的任何错误来解决这个问题。
// 如果发现错误，它会发送一个带有正确状态码的结构化 JSON 响应。

func MyErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last().Err

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
		}
	}
}


func main() {

	router := gin.Default()

	router.Use(MyErrorHandler())

	router.GET("/ok", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Everything is fine!",
		})
	})

	router.GET("/error", func(ctx *gin.Context) {
		ctx.Error(errors.New("something went wrong"))
	})

	router.Run()

	/*
	curl http://localhost:8080/ok

	curl http://localhost:8080/error
	
	*/
}