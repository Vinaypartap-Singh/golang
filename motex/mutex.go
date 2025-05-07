package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.Lock() // Lock before accessing the shared variable
	counter++
	mutex.Unlock() // Unlock after done
}

func main() {
	var wg sync.WaitGroup

	for range 1000 {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
