package main

/*
math/big 任意进度数学包, 和math(标准浮点数学函数包)是两个不同的包
1. math包
	适用于普通科学计算, 入参和返回值基本都是float64
1.1 常用常量与特殊值
	math.Pi              // 3.141592653589793
	math.E               // 2.718281828459045
	math.MaxFloat64      // float64 最大值
	math.Inf(1)          // 正无穷；参数 -1 表示负无穷
	math.NaN()           // 非数字
	math.IsNaN(x)        // 是否为 NaN
	math.IsInf(x, 1)     // 是否为无穷（1 正、-1 负、0 任意）
1.2 取整类
	math.Floor(3.7)   // 3   向下取整
	math.Ceil(3.2)    // 4   向上取整
	math.Round(3.5)   // 4   四舍五入（Go 1.10+）
	math.Trunc(3.9)   // 3   截断小数部分
1.3 幂 开方 对数
	math.Pow(2, 10)      // 1024    x 的 y 次方
	math.Pow10(3)        // 1000    10 的 n 次方
	math.Sqrt(16)        // 4       平方根
	math.Cbrt(27)        // 3       立方根
	math.Exp(1)          // e^1
	math.Log(math.E)     // 1       自然对数
	math.Log2(8)         // 3
	math.Log10(1000)     // 3
1.4 三角函数
	math.Sin(math.Pi / 2)   // 1
	math.Cos(0)             // 1
	math.Tan(0)             // 0
	math.Asin(1)            // π/2，反三角同理还有 Acos、Atan
1.5 其他常用
	math.Abs(-5)      // 5
	math.Max(1.2, 3)  // 3（Go 1.21 起内置 min/max，math 版本仅限 float64）
	math.Min(1.2, 3)  // 1.2
	math.Mod(7, 3)    // 1  取余
	math.Hypot(3, 4)  // 5  直角三角形斜边

2 `math/big` 包：任意精度计算
| 类型      | 含义                            |
| --------- | ------------------------------- |
| big.Int   | 任意精度整数(余额, 区块号用它)  |
| big.Float | 任意精度(可配置)浮点数          |
| big.Rat   | 任意精度 分数(有理数)           |

2.1 big.Int常用API
	// 创建
	x := big.NewInt(100)
	y, _ := new(big.Int).SetString("12345678901234567890", 10) // 字符串转大整数，10 进制

	// 四则运算（注意：结果存入接收者 z，返回 z，支持链式）
	z := new(big.Int)
	z.Add(x, y)     // 加
	z.Sub(x, y)     // 减
	z.Mul(x, y)     // 乘
	z.Div(x, y)     // 整除
	z.Mod(x, y)     // 取余
	z.DivMod(x, y, m) // 同时得商和余数
	z.Neg(x)        // 取负

	// 比较：返回 -1 / 0 / 1
	z.Cmp(y)        // z < y / == / >

	// 转换
	z.String()               // 转十进制字符串
	z.Int64()                // 转 int64（超范围结果不正确，需先 IsInt64() 判断）
	z.SetBytes(bytes)        // 字节转大整数（解析链上数据常用）
	z.Bytes()                // 大整数转字节
2.2 big.Float 常用API
	f := big.NewFloat(3.14)
	f.SetInt(x)      // 从 big.Int 转
	f.SetFloat64(1.5)

	// 运算（接收者存结果，和 Int 一致）
	f.Add(a, b)
	f.Sub(a, b)
	f.Mul(a, b)
	f.Quo(a, b)      // 除法
	f.Cmp(b)         // 比较

	// 输出/转换
	f.Text('f', 6)   // 保留 6 位小数字符串
	v, _ := f.Float64()

2.3 big.Rat 分数, 需要精确小数时很有用
	r := big.NewRat(1, 3)        // 1/3
	r.SetFrac(a, b)              // a/b
	r.FloatString(6)             // "0.333333"，精确无舍入误差
*/
