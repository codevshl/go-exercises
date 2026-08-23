package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func producer(ctx context.Context, limit int, jobs chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(jobs)

	for job := range limit {
		select {
		case <-ctx.Done():
			return
		case jobs <- job:
			fmt.Printf("produced: %d\n", job)
		}
	}
}

func consumer(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			fmt.Printf("consumer %d: %d\n", id, job)
		}
	}
}

func main() {
	const limit = 5
	jobs := make(chan int, 3)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	// One producer and two consumers call Done.
	wg.Add(3)

	go producer(ctx, limit, jobs, &wg)
	go consumer(ctx, 1, jobs, &wg)
	go consumer(ctx, 2, jobs, &wg)

	wg.Wait()
}
