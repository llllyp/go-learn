package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
json.Marshal 会出于安全考虑将特殊HTML字符替换为Unicode转义序列 -- 例如 < 变成 \u003c
当 JSON 嵌入 HTML时这很好, 但如果正在构建纯 API, 客户端可能期望得到原始字符

c.PureJSON 使用 json.Encoder 并设置 SetEscapeHTML(false)，因此 <、> 和 & 等 HTML 字符会按原样呈现而不会被转义。

*/

func main() {
	route := gin.Default()

	route.GET("/json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"html": "<b>Hello World!</b>",
		})
	})

	route.GET("/purejson", func(ctx *gin.Context) {
		ctx.PureJSON(http.StatusOK, gin.H{
			"html": "<b>Hello World!</b>",
		})
	})

	route.Run()

	/*
	curl http://localhost:8080/json
	{"html":"\u003cb\u003eHello World!\u003c/b\u003e"}

	curl http://localhost:8080/purejson
	{"html":"<b>Hello World!</b>"}
	*/
}