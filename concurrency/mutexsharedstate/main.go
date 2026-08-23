package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.value
}

func (c *Counter) Add(delta int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value += delta
}

func main() {
	counter := &Counter{}

	const (
		goRoutineCount = 5
		incrementsEach = 1000
	)

	var wg sync.WaitGroup
	wg.Add(goRoutineCount)

	for range goRoutineCount {
		go func() {
			defer wg.Done()
			for range incrementsEach {
				counter.Increment()
			}
		}()
	}

	wg.Wait()

	fmt.Printf("count: %d\n", counter.Value())

	counter1 := &Counter{}
	const (
		positiveWorkers     = 3
		negativeWorkers     = 2
		positiveDelta       = 10
		negativeDelta       = -5
		operationsPerWorker = 100
	)

	wg.Add(positiveWorkers + negativeWorkers)

	for range positiveWorkers {
		go func() {
			defer wg.Done()
			for range operationsPerWorker {
				counter1.Add(positiveDelta)
			}
		}()
	}

	for range negativeWorkers {
		go func() {
			defer wg.Done()
			for range operationsPerWorker {
				counter1.Add(negativeDelta)
			}
		}()
	}

	wg.Wait()

	fmt.Println("Value after delta add operation:", counter1.Value())
}
