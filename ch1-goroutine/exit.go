package ch1_goroutine

import (
	"fmt"
	"sync"
)

// 通过close来关闭cancel管道向多个Goroutine广播退出的指令。
// 不过这个程序依然不够稳健：当每个Goroutine收到退出指令退出时一般会进行一定的清理工作，
// 但是退出的清理工作并不能保证被完成，因为main线程并没有等待各个工作Goroutine退出工作完成的机制。
// 我们可以结合sync.WaitGroup来改进:

func worker(wg *sync.WaitGroup, cancel chan bool) {
	defer wg.Done()

	for {
		select {
		case <-cancel:
			return
		default:
			fmt.Println("hello")
		}
	}
}
