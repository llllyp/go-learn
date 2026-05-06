package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
ShouldBindQuery 仅将 URL 查询字符串参数绑定到结构体，完全忽略请求体。
当需要确保 POST 请求体数据不会意外覆盖查询参数时 —— 例如在同时接受查询过滤器和 JSON 请求体的端点中——这将非常有用。
*/

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	route := gin.Default()
	route.Any("/testing", func(ctx *gin.Context) {
		var person Person
		if err := ctx.ShouldBindQuery(&person); err != nil {
			ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
		"name":    person.Name,
		"address": person.Address,
	})

	/*
	curl "http://localhost:8080/testing?name=appleboy&address=xyz"
	{"address":"xyz","name":"appleboy"}
	
	curl -X POST "http://localhost:8080/testing?name=appleboy&address=xyz" \
  		-d "name=ignored&address=ignored"

	{"address":"xyz","name":"appleboy"}  ShouldBindQuery 忽略了请求体 -d, 只读查询参数
	*/
	
  })

  route.Run()
}
