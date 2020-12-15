package ch1_goroutine

import (
	"context"
	"fmt"
	"sync"
)

// 在Go1.7发布时，标准库增加了一个context包，
// 用来简化对于处理单个请求的多个Goroutine之间与请求域的数据、超时和退出等操作，
// 官方有博文对此做了专门介绍。我们可以用context包来重新实现前面的线程安全退出或超时的控制:

func worker1(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			fmt.Println("hello")
		}
	}
}
