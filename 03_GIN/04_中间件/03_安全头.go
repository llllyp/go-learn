package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  expectedHost := "localhost:8080"

  // Setup Security Headers
  r.Use(func(c *gin.Context) {
    if c.Request.Host != expectedHost {
      c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
      return
    }
    c.Header("X-Frame-Options", "DENY")
	// 控制浏览器允许加载哪些资源（脚本、样式、图片、字体等）以及从哪些来源。这是防御跨站脚本（XSS）最有效的方式之一，因为它可以阻止内联脚本并限制脚本来源。
    c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
    // 激活浏览器内置的 XSS 过滤器。此头在现代浏览器中已基本弃用（Chrome 在 2019 年移除了其 XSS Auditor），但它仍为使用旧浏览器的用户提供纵深防御。
	c.Header("X-XSS-Protection", "1; mode=block")
	// 强制浏览器在指定的 max-age 期间对所有未来请求使用 HTTPS。这可以防止协议降级攻击和通过不安全 HTTP 连接的 cookie 劫持。includeSubDomains 指令将此保护扩展到所有子域。
    c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	// 控制传出请求中发送多少引用者信息。没有此头，完整的 URL（包括可能包含令牌或敏感数据的查询参数）可能会泄露给第三方站点。strict-origin 仅发送来源（域名）且仅通过 HTTPS。
    c.Header("Referrer-Policy", "strict-origin")
	// 防止 MIME 类型嗅探攻击。没有此头，浏览器可能将文件解释为与声明不同的内容类型，允许攻击者执行伪装成无害文件类型的恶意脚本（例如上传一个实际上是 JavaScript 的 .jpg）。
    c.Header("X-Content-Type-Options", "nosniff")
	// 限制页面可以使用哪些浏览器功能（地理位置、摄像头、麦克风等）。如果攻击者成功注入脚本，这可以限制损害，因为这些脚本无法访问敏感的设备 API。
    c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
    c.Next()
  })

  r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })

  r.Run() // listen and serve on 0.0.0.0:8080
}