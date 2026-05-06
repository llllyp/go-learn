package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Gin 的表单绑定通过 form 结构体标签中的 default 选项支持默认值。这适用于标量值，
从 Gin v1.11 开始，也适用于具有显式集合格式的集合（切片/数组）。

- 将默认值放在表单键之后：form:"name,default=William"。

multi（默认）：重复键或逗号分隔的值
csv：逗号分隔的值
ssv：空格分隔的值
tsv：制表符分隔的值
pipes：管道符分隔的值

*/

type Filters struct {
  Tags      []string `form:"tags" collection_format:"csv"`     // /search?tags=go,web,api
  Labels    []string `form:"labels" collection_format:"multi"` // /search?labels=bug&labels=helpwanted
  IdsSSV    []int    `form:"ids_ssv" collection_format:"ssv"`  // /search?ids_ssv=1 2 3
  IdsTSV    []int    `form:"ids_tsv" collection_format:"tsv"`  // /search?ids_tsv=1\t2\t3
  Levels    []int    `form:"levels" collection_format:"pipes"` // /search?levels=1|2|3
}


type Person struct {
  Name      string    `form:"name,default=William"`
  Age       int       `form:"age,default=10"`
  Friends   []string  `form:"friends,default=Will;Bill"`  // multi/csv: use ; in defaults
  Addresses [2]string `form:"addresses,default=foo bar" collection_format:"ssv"`
  LapTimes  []int     `form:"lap_times,default=1;2;3" collection_format:"csv"`
}

func main() {
	route := gin.Default()

	route.POST("/person", func(ctx *gin.Context) {
		var req Person
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
 		ctx.JSON(http.StatusOK, req)
	})

	route.Run()

	/*
	curl -X POST http://localhost:8080/person
	{
		"Name": "William",
		"Age": 10,
		"Friends": ["Will", "Bill"],
		"Addresses": ["foo", "bar"],
		"LapTimes": [1, 2, 3]
	}
	*/
}