package main

import (
	"fmt"
	"sync"
)

func printOdd(start int, end int, oddTurn <-chan struct{}, evenTurn chan<- struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := start; i <= end; i += 2 {
		<-oddTurn
		fmt.Println(i)
		if i < end {
			evenTurn <- struct{}{}
		}
	}
}

func printEven(start int, end int, evenTurn <-chan struct{}, oddTurn chan<- struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := start; i <= end; i += 2 {
		<-evenTurn
		fmt.Println(i)
		if i < end {
			oddTurn <- struct{}{}
		}
	}
}

// assuming start will always be <= end
func main() {
	const (
		start = 6
		end   = 6
	)

	oddTurn := make(chan struct{})
	evenTurn := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	go printOdd(start+1, end, oddTurn, evenTurn, &wg)
	go printEven(start, end, evenTurn, oddTurn, &wg)

	evenTurn <- struct{}{}

	wg.Wait()
}
