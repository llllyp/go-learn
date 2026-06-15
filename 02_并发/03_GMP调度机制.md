# 1. Goroutine背后的执行模型
Go 运行时采用 G-M-P（Goroutine、Machine、Processor）调度模型 支撑高并发。三者职责分别是：

G（Goroutine）：用户级协程，包含栈、状态等上下文，是被调度的基本单元。
M（Machine）：映射到操作系统线程，用于真正执行 G。
P（Processor）：逻辑处理器，持有可运行 G 的本地队列，并维护内存分配缓存。
只有当 G 绑定到 P，再由拥有该 P 的某个 M 执行时，协程才会真正运行，形成 G → P → M 的执行链路。参考

# 2. 调度关键机制
- 本地队列优先：M 首先从自身绑定的 P 的本地队列中取 G，避免全局锁。
- Work Stealing：本地队列为空时，M 会从其他 P 窃取一半就绪 G，保证均衡。
- 全局队列兜底：全局可运行队列确保没有 P 被饿死。
- 自旋复用线程：没有可运行 G 时，M 会短暂自旋等待，减少频繁创建/销毁线程。
- Hand Off 机制：当 G 因系统调用阻塞，P 会解绑并交给其他空闲 M，维持整体吞吐。

# 3. 实用调试建议
- 通过 GODEBUG=schedtrace=1000 观察 gomaxprocs、threads、runqueue 等调度指标。
- 使用 go tool trace 获取更详细的时间线视图，分析 G 的生命周期。
- 结合 WaitGroup、context 管理 G 的退出和取消，可以更好地配合调度器。参考
- 配套示例： lesson-01/examples/advanced/06-gmp.go 会创建一批 CPU 密集型 goroutine，并调整 GOMAXPROCS。可搭配命令
- GODEBUG=schedtrace=1000,scheddetail=1 go run lesson-01/examples/advanced/06-gmp.go 现场观察调度日志。

