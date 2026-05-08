package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

/*
默认情况下，Gin 将所有日志输出写入 os.Stdout。可以在创建路由器之前通过设置 gin.DefaultWriter 来重定向。

*/

func main() {
	gin.DisableConsoleColor()

	logFile, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(logFile)
	// 同时写入文件和控制台
	// gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	// 生产环境中的日志轮转 lumberjack 日志轮转库
	gin.DefaultWriter = &lumberjack.Logger{
		Filename:   "gin.log",
		MaxSize:    100, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	}

	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	router.Run()
}
