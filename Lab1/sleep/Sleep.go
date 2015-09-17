//package main


package mySleep

import (
	"fmt"
	"time"
)

func subroutine(n int) {
	for i := 0; i < 3; i++ {
		fmt.Println(n, ":", i)
		minutes := 1
		amt := time.Duration(minutes)
		CustomSleep(time.Minute * amt)
	}
}
func main() {
	go subroutine(0)
	var inputtoend string
	fmt.Scanln(&inputtoend)
}

func CustomSleep(timeduration time.Duration) {
	<-time.After(time.Duration(timeduration))
}
