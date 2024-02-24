package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

// Thread 1
func main() {
	waiGroup := sync.WaitGroup{}
	waiGroup.Add(25)

	// Thread 2
	go task("A", &waiGroup)

	// Thread 3
	go task("B", &waiGroup)

	// Thread 4
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
			waiGroup.Done()
		}
	}()
	waiGroup.Wait()
}
