package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
在 *gin.Context 上调用一个方法，传入 HTTP 状态码和要序列化的数据。Gin 会自动处理 Content-Type 头、序列化和写入响应

// All rendering methods share this pattern:
c.JSON(http.StatusOK, data)   // application/json
c.XML(http.StatusOK, data)    // application/xml
c.YAML(http.StatusOK, data)   // application/x-yaml
c.TOML(http.StatusOK, data)   // application/toml
c.ProtoBuf(http.StatusOK, data) // application/x-protobuf

*/

func main() {
	router := gin.Default()

	router.GET("/user", func(ctx *gin.Context) {

		user := gin.H{"name": "Lena", "role": "admin"}

		switch ctx.Query("format") {
		case "xml":
			ctx.XML(http.StatusOK, user)
		case "yaml":
			ctx.YAML(http.StatusOK, user)
		default:
			ctx.JSON(http.StatusOK, user)
		}
	})

	router.Run()

}