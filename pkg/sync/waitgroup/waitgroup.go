package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	// var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i <= 5; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println("Work done for ", i)
		}(i)
	}

	wg.Wait()
}
