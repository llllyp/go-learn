package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
gin.Default() — 带有 Logger 和 Recovery
- Logger -- 将请求日志写入标准输出(方法, 路径, 状态码, 延迟)
- Recovery -- 从处理函数中的任务 panic 恢复并返回 500 响应, 防止服务器崩溃

gin.New() — 一个空白引擎
创建一个完全空白的路由器, 不附加任何中间件

中间件三个级别
- 全局中间件
	应用于路由器中的每个路由. 使用 router.Use()注册, 适用于日志记录和panic恢复等普遍使用的关注点

- 分组中间件
	应用于路由组中的所有路由. 使用 group.Use() 注册. 适用于将认证或授权应用到路由子集

- 路由级中间件
	仅用于单个路由. 作为额外参数传递给 router.GET()等, 适用于路由特定逻辑, 如自定义限流或输入验证

执行顺序: cxt.Next() 将控制权传递给下一个中间件, 然后在 cxt.Next()返回后继续执行(LIFO). 
如果中间件不调用 cxt.Next() 后续的中间件和处理函数将被跳过
- next 之前 -> 此处的代码在请求到达主忽略此函数之前运行. 用于设置任务, 如记录开始时间, 验证令牌或使用cxt设置上下文
- next -> 调用链中的下一个处理函数, 执行在次数暂停, 直到所有下游处理函数完成
- next 之后 -> 次数的代码在主处理函数完后后运行, 用于清理, 记录响应状态或测量延迟
*/

func main() {
	router := gin.New()

	router.Use(MyLogger())
	router.Use(gin.Recovery())

	router.GET("/test", func(ctx *gin.Context) {
		example := ctx.MustGet("example").(string)
		log.Println(example)
	})

	log.Fatal(router.Run())
}

// 自定义中间件
func MyLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		now := time.Now()

		// Set example variable
		ctx.Set("example", "12345")

		// before request

		ctx.Next()

		// after request

		latency := time.Since(now)
		log.Print(latency)

		// access the status we are sending
		status := ctx.Writer.Status()
		log.Println(status)
	}
}