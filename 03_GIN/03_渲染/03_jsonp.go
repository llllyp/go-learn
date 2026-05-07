package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
JSONP（JSON with Padding）是一种在不支持 CORS 的旧版浏览器中进行跨域请求的技术。
它通过将 JSON 响应包装在 JavaScript 函数调用中来工作。浏览器通过 <script> 标签加载响应（不受同源策略限制），
包装函数会以数据作为参数执行。当调用 c.JSONP() 时，Gin 会检查 callback 查询参数。
如果存在，响应体将被包装为 callbackName({"foo":"bar"})，Content-Type 为 application/javascript。
如果没有提供 callback，响应的行为与标准的 c.JSON() 调用相同。

JSONP 是一种遗留技术。对于现代应用，改用 CORS。CORS 更安全，支持所有 HTTP 方法（不仅仅是 GET），
并且不需要将响应包装在回调中。仅在需要支持非常旧的浏览器或与需要它的第三方系统集成时使用 JSONP。
*/

func main() {
	router := gin.Default()

	router.GET("/jsonp", func(ctx *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		ctx.JSONP(http.StatusOK, data)
	})

	router.Run()
	/*
	curl "http://localhost:8080/jsonp?callback=handleData"
	handleData({"foo":"bar"})
	
	curl "http://localhost:8080/jsonp"
	{"foo":"bar"}
	*/
}