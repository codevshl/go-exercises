package main

import (
	"fmt"
	"sync"
)

func printPattern(char string, current <-chan struct{}, next chan<- struct{}, times int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < times; i++ {
		<-current
		fmt.Print(char)
		if i == times-1 && char == "C" {
			fmt.Print("X")
			continue
		}

		next <- struct{}{}
	}
}

func main() {
	const limit = 3
	chanA := make(chan struct{})
	chanB := make(chan struct{})
	chanC := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(3)

	go printPattern("A", chanA, chanB, limit, &wg)
	go printPattern("B", chanB, chanC, limit, &wg)
	go printPattern("C", chanC, chanA, limit, &wg)

	chanA <- struct{}{}

	wg.Wait()
	fmt.Println()
}
