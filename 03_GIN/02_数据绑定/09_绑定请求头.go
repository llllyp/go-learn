package main

/*
ShouldBindHeader 使用 header 结构体标签将 HTTP 请求头直接绑定到结构体中。
这对于从传入请求中提取元数据（如 API 速率限制、认证令牌或自定义域头信息）非常有用。

根据 HTTP 规范，请求头名称不区分大小写。header 结构体标签值会进行不区分大小写的匹配，
因此 header:"Rate" 将匹配以 Rate、rate 或 RATE 发送的请求头。
*/

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

type testHeader struct {
  Rate   int    `header:"Rate"`
  Domain string `header:"Domain"`
}

func main() {
  r := gin.Default()

  r.GET("/", func(c *gin.Context) {
    h := testHeader{}

    if err := c.ShouldBindHeader(&h); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    c.JSON(http.StatusOK, gin.H{"Rate": h.Rate, "Domain": h.Domain})
  })

  r.Run()

  /*
	curl -H "Rate:300" -H "Domain:music" http://localhost:8080/
	{"Domain":"music","Rate":300}
 
	curl http://localhost:8080/
	{"Domain":"","Rate":0}
  */
}