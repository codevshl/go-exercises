package main

import (
	"context"
	"fmt"
	"time"
)

func newRateLimiter(
	ctx context.Context,
	interval time.Duration,
	burst int,
) <-chan struct{} {
	tokens := make(chan struct{}, burst)

	for range burst {
		tokens <- struct{}{}
	}

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker.C:

				select {
				case tokens <- struct{}{}:
				default:
				}
			}
		}
	}()

	return tokens
}

func main() {
	const (
		requestCount = 8
		burst        = 5
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	limiter := newRateLimiter(
		ctx,
		500*time.Millisecond,
		burst,
	)

	for requestID := 1; requestID <= requestCount; requestID++ {
		select {
		case <-ctx.Done():
			fmt.Println("rate limiter stopped")
			return

		case <-limiter:
			fmt.Printf(
				"request %d allowed at %s\n",
				requestID,
				time.Now().Format("15:04:05.000"),
			)
		}
	}

	cancel()
}
