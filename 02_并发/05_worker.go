package main

import "sync"

func worker(tasks <- chan func(), wg *sync.WaitGroup) {
	defer wg.Done()
	for fn := range tasks {
		fn()
	}
}

func main() {
	const workerCount = 10
	tasks := make(chan func())

	var wg sync.WaitGroup
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go worker(tasks, &wg)
	}

	// 放入任务 10个worker 抢100个任务
	for i := 0; i < 100; i++ {
		tasks <- func() {
			// do something
		}
	}

	close(tasks)
	wg.Wait()
}