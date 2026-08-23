package main

import (
	"context"
	"fmt"
	"sync"
)

// 1...5
type Result struct {
	Job      int
	Value    int
	WorkerID int
}

func submitJobs(ctx context.Context, jobsToProcess []int, jobs chan<- int) {
	defer close(jobs)

	for _, job := range jobsToProcess {
		fmt.Println(job)
		select {
		case <-ctx.Done():
			return
		case jobs <- job:
		}
	}
}

func worker(ctx context.Context, cancel context.CancelFunc, id int, jobs <-chan int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			if job == 8 {
				cancel()
				return
			}
			result := Result{
				Job:      job,
				Value:    job * job,
				WorkerID: id,
			}
			results <- result
		}
	}
}

func collectResults(results <-chan Result) {
	for result := range results {
		fmt.Printf("worker %d produced %d > %d\n", result.WorkerID, result.Job, result.Value)
	}
}

func main() {
	const (
		noOfWorkers = 5
	)
	jobsToProcess := []int{8, 8, 8, 5, 7}
	ctx, cancel := context.WithCancel(context.Background())

	jobs := make(chan int)
	results := make(chan Result)

	var wg sync.WaitGroup
	wg.Add(noOfWorkers)

	go submitJobs(ctx, jobsToProcess, jobs)

	for id := range noOfWorkers {
		go worker(ctx, cancel, id, jobs, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	collectResults(results)
}
