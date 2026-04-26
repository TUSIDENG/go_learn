package main

import (
	"sync"

	"go_learn/package/system_package/01/sync_example"
)

func main() {
	counter := sync_example.NewSafeCounter()

	// Increment the counter for "example" key
	counter.Inc("example")

	// Get the value for "example" key
	value := counter.Value("example")
	println("Value for 'example':", value)

	// Use goroutines to increment concurrently
	var wg sync.WaitGroup
	workers := 5
	incrementsPerWorker := 1000
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerWorker; j++ {
				counter.Inc("example")
				println("Incremented 'example' in goroutine")
				value := counter.Value("example")
				println("Current value for 'example':", value)
				// Simulate some work
				// time.Sleep(10 * time.Millisecond)
				// Note: Uncomment the above line to simulate work and see more interleaving in output
			}
		}()
	}

	wg.Wait()

	value = counter.Value("example")
	println("Value for 'example' after concurrent increments:", value)

	// Reset the counter
	counter.Reset()

	// Get the value again after reset
	value = counter.Value("example")
	println("Value for 'example' after reset:", value)
}
