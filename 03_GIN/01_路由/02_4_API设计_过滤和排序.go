package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 通过查询参数接受过滤和排序. 使用一致的参数名称保持接口可预测性
func main() {
	router := gin.Default()

	router.GET("/api/products", func(ctx *gin.Context) {
		category := ctx.Query("category")
		minPrice := ctx.Query("min_price")
		maxPrice := ctx.Query("max_price")

		sortBy := ctx.DefaultQuery("sort", "created_at")
		order := ctx.DefaultQuery("order", "desc")

		allowed := map[string]bool {"create_at": true, "price": true, "name": true}
		if !allowed[sortBy] {
			sortBy = "created_at"
		}

		if order != "asc" && order != "desc" {
			order = "desc"
		}

		// Build and execute your query using these filters...
		_ = category
		_ = minPrice
		_ = maxPrice

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    []gin.H{},
			"filters": gin.H{
				"category":  category,
				"min_price": minPrice,
				"max_price": maxPrice,
				"sort":      sortBy,
				"order":     order,
			},
		})
	})
}