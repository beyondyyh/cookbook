package ch1_goroutine

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// run: go test -v -run Test_Prime
func Test_Prime(t *testing.T) {
	ch := GenerateNatural()
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime)
	}
}

// run: go test -v -run Test_work
func Test_work(t *testing.T) {
	cancel := make(chan bool)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, cancel)
	}

	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
}

// run: go test -v -run Test_worker1
func Test_work1(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker1(ctx, &wg)
	}

	time.Sleep(time.Second)
	cancel()

	wg.Wait()
}
