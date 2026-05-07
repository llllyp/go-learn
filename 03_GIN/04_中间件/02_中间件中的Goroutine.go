package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

/*
在中间件或处理函数中启动新的 Goroutine 时，不应该在其中使用原始上下文，必须使用只读副本。

一旦处理函数返回，gin.Context 就会被返回到池中，可能会被分配给一个完全不同的请求。
如果此时一个 goroutine 仍然持有对原始上下文的引用，它将读取或写入现在属于另一个请求的字段。
这会导致竞态条件、数据损坏或 panic。
调用 c.Copy() 会创建一个上下文快照，可以在处理函数返回后安全使用。
副本包含请求、URL、键和其他只读数据，但与池的生命周期分离。
*/

func main() {

	router := gin.Default()

	// 异步请求
	router.GET("/long_async", func(ctx *gin.Context) {
		// 创建一份副本，供 goroutine 内部使用
		cpCtx := ctx.Copy()
		go func ()  {
			time.Sleep(5 * time.Second)
			
			log.Println("Done! in path " + cpCtx.Request.URL.Path)
		}()
	})

	// 同步请求
	router.GET("/long_sync", func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)

		log.Println("Done! in path", ctx.Request.URL.Path)
	})

	router.Run()

	// 
}