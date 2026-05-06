package main

/*
在要绑定的字段的 uri/form 标签中指定 parser=encoding.TextUnmarshaler

*/
import (
	"encoding"
	"strings"

	"github.com/gin-gonic/gin"
)

type Birthday string

func (b *Birthday) UnmarshalText(text []byte) error {
  *b = Birthday(strings.Replace(string(text), "-", "/", -1))
  return nil
}
var _ encoding.TextUnmarshaler = (*Birthday)(nil) 
/*
go 经典技巧
- 编译期就检查, *Birthday 是否实现 encoding.TextUnmarshaler 接口
- 如果没有实现正确, 编译直接报错, 而不是运行时才发现
- nil 指针不占用内存, 纯静态检查
*/

func main() {

	route := gin.Default()

	var request struct {
    Birthday         Birthday   `form:"birthday,parser=encoding.TextUnmarshaler"`
    Birthdays        []Birthday `form:"birthdays,parser=encoding.TextUnmarshaler" collection_format:"csv"`
    BirthdaysDefault []Birthday `form:"birthdaysDef,default=2020-09-01;2020-09-02,parser=encoding.TextUnmarshaler" collection_format:"csv"`
  }

  route.GET("/test", func(ctx *gin.Context) {
    _ = ctx.BindQuery(&request)
    ctx.JSON(200, request)
  })
  route.Run(":8088")
  /*
 curl 'localhost:8088/test?birthday=2000-01-01&birthdays=2000-01-01,2000-01-02'
 
 {
  "Birthday": "2000/01/01",
  "Birthdays": [
    "2000/01/01",
    "2000/01/02"
  ],
  "BirthdaysDefault": [
    "2020/09/01",
    "2020/09/02"
  ]
}
  */
}