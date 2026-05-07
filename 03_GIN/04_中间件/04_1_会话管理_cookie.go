package main


// go get github.com/gin-contrib/sessions

import (
  "net/http"

  "github.com/gin-contrib/sessions"
  "github.com/gin-contrib/sessions/cookie"
  "github.com/gin-gonic/gin"
)

// 将会话数据存储在加密的 cookie 中;
func main() {
	router := gin.Default()

	store := cookie.NewStore([]byte("your-secret-key"))
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/login", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		session.Set("user", "john")
		session.Save()
		ctx.JSON(http.StatusOK, gin.H{
			"message": "logged in",
		})
	})

	router.GET("/profile", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		user := session.Get("user")
		if user == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "not logged in",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"user": user})
	})

	router.GET("/logout", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		session.Clear()
		session.Save()
	})

	router.Run()
}

