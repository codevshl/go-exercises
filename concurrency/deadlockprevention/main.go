package main

import (
	"fmt"
	"sync"
	"time"
)

func firstTask(lockA, lockB *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	lockA.Lock()
	defer lockA.Unlock()

	fmt.Println("First task acquired lockA")

	time.Sleep(100 * time.Millisecond)

	lockB.Lock()
	defer lockB.Unlock()

	fmt.Println("First task acquired lockB")
}

func secondTask(lockA, lockB *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	lockA.Lock()
	defer lockA.Unlock()

	fmt.Println("Second task acquired lockA")

	time.Sleep(100 * time.Millisecond)

	lockB.Lock()
	defer lockB.Unlock()

	fmt.Println("Second task acquired lockB")
}

func main() {
	var (
		lockA sync.Mutex
		lockB sync.Mutex
		wg    sync.WaitGroup
	)

	wg.Add(2)

	go firstTask(&lockA, &lockB, &wg)
	go secondTask(&lockA, &lockB, &wg)

	wg.Wait()
}
