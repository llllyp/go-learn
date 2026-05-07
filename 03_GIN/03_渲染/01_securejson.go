package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
SecureJSON 可以防御一类称为 JSON 劫持的漏洞。
在旧版浏览器（主要是 Internet Explorer 9 及更早版本）中，
恶意页面可以包含一个 <script> 标签指向受害者的 JSON API 端点。
如果该端点返回顶层 JSON 数组（例如 ["secret","data"]），浏览器会将其作为 JavaScript 执行。
通过覆盖 Array 构造函数，攻击者可以拦截解析后的值并将敏感数据泄露到第三方服务器。
*/

func main() {

	router := gin.Default()

	// You can also use your own secure json prefix
  	// router.SecureJsonPrefix(")]}',\n")

	router.GET("/someJson", func(ctx *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		ctx.SecureJSON(http.StatusOK, names)
	})

	router.Run()

	// 现代浏览器已经修复了这个漏洞，因此 SecureJSON 主要在你需要支持旧版浏览器或安全策略要求纵深防御时才有意义。
	// 对于大多数新 API，标准的 c.JSON() 就足够了。
}