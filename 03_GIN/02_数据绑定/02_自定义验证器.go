package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

/*
除了内置的验证器(required, email, min, max等), 还可以自定义验证函数
*/
type Booking struct {
	// form:"check_in" 指定从查询参数中绑定的参数名
	// binding:"..." 验证规则列表，逗号分隔
	// time_format:"2006-01-02" 指定日期解析格式
	CheckIn time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	// gtfield=CheckIn 内置验证器,CheckOut必须晚于CheckIn
	// bookabledate 自定义验证器 ：自己实现的日期规则
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn,bookabledate" time_format:"2006-01-02"`
}

// 自定义验证函数 bookableDate, 验证日期是否在当前时间之后
var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	// validator.Func  验证器库定义的函数类型
	date, ok := fl.Field().Interface().(time.Time) // 将反射值转换为具体的 time.Time 类型
	if ok {                                        // ok 表示转换成功
		today := time.Now()
		if today.After(date) {
			return false // 验证失败, 日期在今天之前
		}
	}
	return true
}

func main() {
	route := gin.Default()

	// binding.Validator.Engine() 获取Gin正在使用的验证器引擎
	// 类型断言转为 *validator.Validate 具体类型
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// v.RegisterValidation("名称", 函数) 注册自定义验证器
		v.RegisterValidation("bookabledate", bookableDate)
	}

	route.GET("/bookable", func(ctx *gin.Context) {
		var b Booking
		if err := ctx.ShouldBindWith(&b, binding.Query); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	route.Run()
}
