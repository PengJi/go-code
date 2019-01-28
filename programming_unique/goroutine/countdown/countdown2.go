package countdown

import (
	"fmt"
	"os"
	"time"
)

// countdown implements the countdown for a rocket launch.

func main() {
	// abort
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))  //read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("commencing countdown. press return to abort.")
	select {
	case <-time.After(10 * time.Second):
		// do nothing.
		fmt.Println("time after")
	case <-abort:
		fmt.Println("launch aborted!")
		return
	}
	launch()
}

func launch() {
	fmt.Println("lift off!")
}