package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func runWorker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(300 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker stopped")
			return
		case <-ticker.C:
			fmt.Println("Worker processing")
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go runWorker(ctx, &wg)

	wg.Wait()

	fmt.Println("main finished")
}
