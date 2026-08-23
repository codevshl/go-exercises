package main

import (
	"fmt"
	"sync"
)

type commandType int

const (
	addCommand commandType = iota
	getCommand
)

type Command struct {
	kind  commandType
	delta int
	reply chan int
}

func runCounter(commands <-chan Command) {
	value := 0

	for cmd := range commands {
		switch cmd.kind {
		case addCommand:
			value += cmd.delta
		case getCommand:
			cmd.reply <- value
		}
	}
}

func main() {
	const (
		senderCount      = 5
		updatesPerSender = 100
		delta            = 10
	)

	commands := make(chan Command)

	go runCounter(commands)

	var wg sync.WaitGroup
	wg.Add(senderCount)

	for range senderCount {
		go func() {
			defer wg.Done()
			for range updatesPerSender {
				commands <- Command{
					kind:  addCommand,
					delta: delta,
				}
			}
		}()
	}

	wg.Wait()

	results := make(chan int)

	commands <- Command{
		kind:  getCommand,
		reply: results,
	}

	finalValue := <-results

	close(commands)
	fmt.Println("final value: ", finalValue)
}
