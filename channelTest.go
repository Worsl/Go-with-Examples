package main

import (
	"fmt"
	"time"
)

func main() {
    timer := time.NewTimer(2 * time.Second)

    // Wait for the timer to expire and receive the time when it does.
    time_ := <-timer.C

	fmt.Println(time_)
    fmt.Println("Timer expired.")
}