package main

import(
  "net/http"

  "github.com/gin-gonic/gin"
)
/*
类型 - Must bind
	方法 - Bind、BindJSON、BindXML、BindQuery、BindYAML
	行为 - 这些方法底层使用 MustBindWith。如果存在绑定错误，请求将使用 c.AbortWithError(400, err).SetType(ErrorTypeBind) 中止。
		会将响应状态码设置为 400，并将 Content-Type 头设置为 text/plain; charset=utf-8。注意，如果你在此之后尝试设置响应码，
		将会出现警告 [GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422。
		希望更好地控制行为，请考虑使用 ShouldBind 等效方法。
类型 - Should bind
	方法 - ShouldBind、ShouldBindJSON、ShouldBindXML、ShouldBindQuery、ShouldBindYAML
	行为 - 这些方法底层使用 ShouldBindWith。如果存在绑定错误，错误会被返回，由开发者负责适当地处理请求和错误。
*/

// Binding from JSON
type Login struct {
  User     string `form:"user" json:"user" xml:"user"  binding:"required"`
  Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
  router := gin.Default()

  // Example for binding JSON ({"user": "manu", "password": "123"})
  router.POST("/loginJSON", func(c *gin.Context) {
    var json Login
    if err := c.ShouldBindJSON(&json); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    if json.User != "manu" || json.Password != "123" {
      c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
      return
    }

    c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
  })

  // Example for binding XML (
  //  <?xml version="1.0" encoding="UTF-8"?>
  //  <root>
  //    <user>manu</user>
  //    <password>123</password>
  //  </root>)
  router.POST("/loginXML", func(c *gin.Context) {
    var xml Login
    if err := c.ShouldBindXML(&xml); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    if xml.User != "manu" || xml.Password != "123" {
      c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
      return
    }

    c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
  })

  // Example for binding a HTML form (user=manu&password=123)
  router.POST("/loginForm", func(c *gin.Context) {
    var form Login
    // This will infer what binder to use depending on the content-type header.
    if err := c.ShouldBind(&form); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    if form.User != "manu" || form.Password != "123" {
      c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
      return
    }

    c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
  })

  // Listen and serve on 0.0.0.0:8080
  router.Run(":8080")
}