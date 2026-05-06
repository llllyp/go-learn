package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
具有相同 name 属性的 HTML 复选框在被选中时会提交多个值。
Gin 可以通过使用带有 [] 后缀的 form 结构体标签（匹配 HTML 的 name 属性）
将这些值直接绑定到结构体的 []string 切片中。

colors[] 中的 [] 后缀是 HTML 的约定，不是 Go 的要求。结构体标签必须与 HTML 的 name 属性完全匹配。
如果 HTML 使用 name="colors"（不带方括号），结构体标签应该是 form:"colors"。
*/


type myForm struct {
	// Colors []string `form:"colors[]"`
	Colors []string `form:"colors"`
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func (ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "form.html", nil)
	})

	router.POST("/", func(ctx *gin.Context) {
		var fakeForm myForm
		if err := ctx.ShouldBind(&fakeForm); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"color": fakeForm.Colors})
	})

	router.Run()

	/*
	curl -X POST http://localhost:8080/ -d "colors[]=red&colors[]=green&colors[]=blue"

	curl -X POST http://localhost:8080/ -d "colors=red&colors=green&colors=blue"
		

	curl -X POST http://localhost:8080/ \
	-d "colors[]=green"

	curl -X POST http://localhost:8080/
	*/
}