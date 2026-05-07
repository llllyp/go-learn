package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
AsciiJSON 将数据序列化为 JSON，但会将所有非 ASCII 字符转义为 \uXXXX Unicode 转义序列。
HTML 特殊字符如 < 和 > 也会被转义。结果是一个仅包含 7 位 ASCII 字符的响应体

- 你的 API 消费者需要严格的 ASCII 安全响应（例如，无法处理 UTF-8 编码字节的系统）。
- 你需要将 JSON 嵌入到仅支持 ASCII 的上下文中，例如某些日志系统或遗留传输层。
- 你希望确保 <、> 和 & 等字符被转义，以避免将 JSON 嵌入 HTML 时的注入问题。
*/

func main() {

	router := gin.Default()

	router.GET("/someJSON", func(ctx *gin.Context) {
		data := map[string]interface{}{
			"lang": "go语言",
			"tag": "<br>",
		}

		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		ctx.AsciiJSON(http.StatusOK, data)
	})

	router.Run()

}