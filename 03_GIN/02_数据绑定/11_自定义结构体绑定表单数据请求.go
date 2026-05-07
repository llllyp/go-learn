package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
// 参数嵌套
type StructA struct {
  FieldA string `form:"field_a"`
}

type StructB struct {
  NestedStruct StructA
  FieldB       string `form:"field_b"`
}

type StructC struct {
  NestedStructPointer *StructA
  FieldC              string `form:"field_c"`
}

type StructD struct {
  NestedAnonyStruct struct {
    FieldX string `form:"field_x"`
  }
  FieldD string `form:"field_d"`
}

func main() {
	router := gin.Default()

	router.GET("/getb", func(ctx *gin.Context) {
		var b StructB
		ctx.Bind(&b)
		ctx.JSON(http.StatusOK, gin.H{
			"a": b.NestedStruct,
			"b": b.FieldB,
		})
	})

	router.GET("/getc", func(ctx *gin.Context) {
    var b StructC
    ctx.Bind(&b)
    ctx.JSON(http.StatusOK, gin.H{
      "a": b.NestedStructPointer,
      "c": b.FieldC,
    })
  })

  router.GET("/getd", func(ctx *gin.Context) {
    var b StructD
    ctx.Bind(&b)
    ctx.JSON(http.StatusOK, gin.H{
      "x": b.NestedAnonyStruct,
      "d": b.FieldD,
    })
  })

  router.Run()

  /*
  curl "http://localhost:8080/getb?field_a=hello&field_b=world"
 
  curl "http://localhost:8080/getc?field_a=hello&field_c=world"

  curl "http://localhost:8080/getd?field_x=hello&field_d=world"
  */
}