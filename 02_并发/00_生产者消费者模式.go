package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func producer(ch chan<- int) {
	for i := range 10 {
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch <-chan int, id int) {
	for value := range ch {
		fmt.Printf("Consumer %d received: %d\n", id, value)
		time.Sleep(50 * time.Millisecond)
	}
}

// func main() {
// 	ch := make(chan int, 5)

// 	go producer(ch)

// 	// 多个消费者
// 	for i := 0; i < 3; i++ {
// 		go consumer(ch, i)
// 	}
// 	time.Sleep(2 * time.Second)
// }

// 使用 goroutine + select 管理停止信号
func producer2(ctx context.Context, id int, queue chan<- int) {
	defer log.Printf("producer %d exit", id)

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		item := rand.Intn(1000)
		select {
		case queue <- item:
			log.Printf("producer %d -> %d\n", id, item)
		case <-ctx.Done():
			return
		}

		time.Sleep(300 * time.Millisecond)
	}
}

func consumer2(ctx context.Context, id int, queue <-chan int) {
	defer log.Printf("consumer %d exit", id)

	for {
		select {
		case item := <-queue:
			log.Printf("consumer %d <- %d\n", id, item)
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	queue := make(chan int, 5)
	ctx, cancel := context.WithCancel(context.Background())

	for i := range 2 {
		go producer2(ctx, i+1, queue)
	}

	for i := range 3 {
		go consumer2(ctx, i+1, queue)
	}

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(500 * time.Microsecond)

	fmt.Println("done")
}
