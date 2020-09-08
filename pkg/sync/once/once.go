package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(sayHello)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-done)
	}
}

func sayHello() {
	fmt.Println("Hello")
}
