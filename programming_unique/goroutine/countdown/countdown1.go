package countdown

import (
	"fmt"
	"time"
)

// countdown implements the countdown for a rocket launch.

func main() {
	fmt.Println("commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown>0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
}

func launch() {
	fmt.Println("lift off!")
}
