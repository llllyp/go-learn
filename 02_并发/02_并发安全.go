package main

import (
	"sync"
	"sync/atomic"
)

/*
1. 互斥锁
	sync.Mutex 同一时间只允许一个goroutine持有锁
2. 读写锁
	sync.RWMutex 适合读多写少的场景, 读可以并发, 写必须独占
	写锁与任何锁都互斥, 读锁不能升级为写锁
3. 原子操作 atomic
	简单的数值操作(增减, 替换, 比较)不需要锁, 用原子指令更换
	无锁, 性能极高
	只支持基本类型: int32,int64,uint32,uint64,uintptr,pointer
4. 一次性执行 sync.Once
	保证某段代码全局只执行一次, 常用于单例, 初始化
	即使多个goroutine同时调用, 也只执行一次
5. 等待组 sync.WaitGroup
	等待一组goroutine全部执行完, 不是保护数据, 是同步流程
	适用: 主goroutine 等待一批协程结束
*/

var mu sync.Mutex
var count int

func add() {
	mu.Lock()   // 加锁
	count++     // 临界锁
	mu.Unlock() // 解锁
}
//--------------------------------------------------
var rmu sync.RWMutex
func method() {
	rmu.RLock()   // 读加锁
	rmu.RUnlock() // 读解锁

	rmu.Lock()    // 写加锁（互斥所有）
	rmu.Unlock()  // 写解锁
}
//--------------------------------------------------
var count3 int64
func method2() {
	atomic.AddInt64(&count3, 1) // 原子+1
	v := atomic.LoadInt64(&count3)
	atomic.StoreInt64(&count3, 0)
}
//--------------------------------------------------
var once sync.Once
var config *Config

func initConfig() {
	once.Do(func() {
		config = loadConfig()
	})
}
func loadConfig() *Config {
	return &Config{"cfg"}
}
type Config struct {
	cfg string
}
//--------------------------------------------------
var wg sync.WaitGroup
func method3() {
	wg.Add(2)
	go func ()  {
		defer wg.Done()
	}()

	go func(){
		defer wg.Done()
	}()
	wg.Wait() //阻塞直到计数为0
}