package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	store, err := redis.NewStoreWithDB(10, "tcp", "localhost:6379", "", "123456", "3", []byte("secret"))
	if err != nil {
		panic(err)
	}

	router.Use(sessions.Sessions("mysession", store))

	router.GET("/set", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		session.Set("count", 1)
		session.Save()
		ctx.JSON(http.StatusOK, gin.H{"count": 1})
	})

	router.GET("/get", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		count := session.Get("count")
		// 获取到的 count 为 null 可能是第二次访问的session变了
		ctx.JSON(http.StatusOK, gin.H{"count": count})
	})

	router.Run()
}
