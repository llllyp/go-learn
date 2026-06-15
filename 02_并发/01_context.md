# Context上下文
context.Context 是Go语言标准库中用于跨 goroutine传递取消信号, 耗时, 截止时间和请求范围值的标准方式. 它是并发控制的核心工具

核心概念
- Context在多个goroutine之间传播控制信号
- Context是不可变的, 每次派生都会创建新的Context
- Context形成树形结构, 父Context取消时, 所有子Context也会被取消

主要用途
- 取消控制: 取消长时间运行的操作
- 超时控制: 为操作设置超时时间
- 截止时间: 设置操作必须完成的最后期限
- 传递值: 在请求范围内传递元数据(如traceID,用户ID等)

## context.Backgroud()
根 Context, 通常用于main函数, 初始化或测试中, 它永远不会被取消, 没有值, 没有截止时间

## context.TODO()
当不确定使用哪个Context时使用, 同产格式占位符, 表示"稍后会替换为真正的Context"

## context.WithCancel(parent)
创建可取消的Context, 返回 Context和cancel函数. 调用 cancel函数会取消该 Context及其所有子Context

## context.WithDeadline(parent, deadline)
创建有截止时间的Context, 到达截止时间后自动取消
```go
deadline := time.Now().Add(10 * time.Second)
ctx, cancel := context.WithDeadline(context.Backbroud(), deadline)
```
### Context基本使用
取消控制示例
```go
func cancellableDemo() {
    ctx, cancel := context.Withcancel(context.Backgroud())

    // 启动工作 goroutine
    go func() {
        for {
            select {
                case <- ctx.Done(): fmt.Println("Goroutine收到取消消息: ", ctx.Err())
                return
                default:
                fmt.Println("工作中......")
                time.Sleep(500 * time.Millisecond)
            }
        }
    }()

    // 工作两秒后取消
    time.Sleep(2 * time.Second)
    fmt.Println("发送取消信号")
    cancel()

    // 等待goroutine退出
    time.Sleep(500 * time.Millisecond)
}
```
超时控制
```go
func timeoutContextDemo() {
	fmt.Println("=== 超时Context ===")
	
	// 设置1秒超时
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	ch := make(chan string)
	
	// 模拟一个耗时2秒的操作
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "result"
	}()
	
	select {
	case result := <-ch:
		fmt.Println("收到结果:", result)
	case <-ctx.Done():
		fmt.Println("操作超时:", ctx.Err())  // 输出: context deadline exceeded
	}
}
```
截止时间示例
```go
func deadlineContextDemo() {
	fmt.Println("=== 截止时间Context ===")
	
	// 设置3秒后的截止时间
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	
	// 检查剩余时间
	if d, ok := ctx.Deadline(); ok {
		fmt.Printf("截止时间: %v, 剩余: %v\n", d, time.Until(d))
	}
	
	// 等待超过截止时间
	time.Sleep(4 * time.Second)
	
	select {
	case <-ctx.Done():
		fmt.Println("已超过截止时间:", ctx.Err())
	default:
		fmt.Println("未超时")
	}
}
```
值传递示例
```go
type contextKey string

const (
	requestIDKey contextKey = "requestID"
	userIDKey    contextKey = "userID"
)

func valueContextDemo() {
	fmt.Println("=== Context传递值 ===")
	
	// 创建携带值的context
	ctx := context.Background()
	ctx = context.WithValue(ctx, requestIDKey, "req-123")
	ctx = context.WithValue(ctx, userIDKey, "user-456")
	
	// 在函数中读取值
	processRequest(ctx)
}

func processRequest(ctx context.Context) {
	if reqID := ctx.Value(requestIDKey); reqID != nil {
		fmt.Printf("Request ID: %v\n", reqID)
	}
	
	if userID := ctx.Value(userIDKey); userID != nil {
		fmt.Printf("User ID: %v\n", userID)
	}
}
```