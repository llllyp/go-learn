package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
DataFromReader 允许你将任何 io.Reader 的数据直接流式传输到 HTTP 响应，而无需先将整个内容缓冲到内存中。
这对于构建代理端点或高效地从远程源提供大文件至关重要。

常见用例：
- 代理远程资源 — 从外部服务（如云存储 API 或 CDN）获取文件并转发给客户端。数据通过服务器流过，而不会完全加载到内存中。
- 提供生成的内容 — 在生产动态生成的数据（如 CSV 导出或报告文件）时进行流式传输。
- 大文件下载 — 提供太大而无法保存在内存中的文件，从磁盘或远程源分块读取。
*/

func main() {
	router := gin.Default()

	router.GET("/someDataFormReader", func(ctx *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			log.Printf("获取图片异常 %s", err.Error())
			ctx.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string {
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		ctx.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})

	router.Run()
}
