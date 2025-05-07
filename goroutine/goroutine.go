package main

import "sync"

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	println("doing task", id)
}

func main() {

	// We can use waitgroup to wait for go routine function to complete

	var wg sync.WaitGroup

	// The function will be scheduled by goroutine and will concurrently like running in parallel

	for i := range 10 {
		wg.Add(2)
		go task(i+1, &wg)

		go func(i int) {
			defer wg.Done()
			println("doing work inside loop", i)
		}(i)
	}

	wg.Wait()
}
