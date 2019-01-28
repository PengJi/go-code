package countdown

import (
	"fmt"
	"os"
	"time"
)

// countdown implements the countdown for a rocket launch.

// the ticker goroutine never terminates if the launch is aborted.
// this is a "goroutine leak".

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))  //read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("commencing countdown. press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown>0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			fmt.Println("test")
		case <-abort:
			fmt.Println("lauch aborted!")
			return
		}
	}
	launch()
}

func launch(){
	fmt.Println("lift off!")
}
