package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
ShouldBind 会根据 HTTP 方法和 Content-Type 请求头自动选择绑定引擎：

- 对于 GET 请求，使用查询字符串绑定（form 标签）。
- 对于 POST/PUT 请求，它会检查 Content-Type——对 application/json 使用 JSON 绑定，
对 application/xml 使用 XML 绑定，对 application/x-www-form-urlencoded 或 multipart/form-data 使用表单绑定。

*/

type Person struct {
  Name     string    `form:"name"`
  Address  string    `form:"address"`
  Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {

	route := gin.Default()
	route.GET("/testing", startPage)
	route.POST("/testing", startPage)
	route.Run()
}

func startPage(ctx *gin.Context) {
	var person Person
	if err := ctx.ShouldBind(&person); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Name: %s, Address: %s, Birthday: %s\n", person.Name, person.Address, person.Birthday)
  	ctx.JSON(http.StatusOK, gin.H{
		"name":     person.Name,
		"address":  person.Address,
		"birthday": person.Birthday,
	})

	/*
	curl "http://localhost:8080/testing?name=appleboy&address=xyz&birthday=1999-03-15"

	curl -X POST http://localhost:8080/testing \
  	-d "name=appleboy&address=xyz&birthday=1999-03-15"

	curl -X POST http://localhost:8080/testing \
	-H "Content-Type: application/json" \
	-d '{"name":"appleboy","address":"xyz","birthday":"1999-03-15"}'
	*/
 }