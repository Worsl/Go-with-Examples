package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	// Simulate some work
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
	wg.Done()

}

func main() {
	numOfWorkers := 5

	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1)
		go worker(i)

	}

	wg.Wait()
}
