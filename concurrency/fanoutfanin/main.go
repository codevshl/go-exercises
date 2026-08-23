package main

import (
	"context"
	"fmt"
	"sync"
)

type Result struct {
	ID    int
	Value int
	Err   error
}

func generate(ctx context.Context, values []int) <-chan int {
	input := make(chan int)
	go func() {
		defer close(input)
		for _, value := range values {
			select {
			case <-ctx.Done():
				return
			case input <- value:
			}
		}
	}()

	return input
}

func square(ctx context.Context, id int, input <-chan int) <-chan Result {
	output := make(chan Result)

	go func() {
		defer close(output)
		for {
			select {
			case <-ctx.Done():
				return
			case value, ok := <-input:
				if !ok {
					return
				}
				result := Result{
					ID:    id,
					Value: value * value,
				}
				if value == 3 {
					result.Err = fmt.Errorf("invalid value: %d", value)
				}
				select {
				case <-ctx.Done():
					return
				case output <- result:
				}

				if result.Err != nil {
					return
				}
			}
		}
	}()

	return output
}

func fanIn(ctx context.Context, channels ...<-chan Result) <-chan Result {
	merged := make(chan Result)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, channel := range channels {
		go func(ch <-chan Result) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case ch, ok := <-ch:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case merged <- ch:
					}
				}
			}
		}(channel)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {
	values := []int{4, 5, 8, 0, 3}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	input := generate(ctx, values)

	worker1 := square(ctx, 1, input)
	worker2 := square(ctx, 2, input)
	worker3 := square(ctx, 3, input)

	results := fanIn(ctx, worker1, worker2, worker3)

	for result := range results {
		if result.Err != nil {
			fmt.Printf("Worker: %d produced an error: %v\n", result.ID, result.Err)
			cancel()
			continue
		}
		fmt.Printf("Worker: %d produced value: %d\n", result.ID, result.Value)
	}
}
