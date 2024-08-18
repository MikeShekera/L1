package main

import (
	"time"
)

func main() {
	mySleep(1 * time.Second)
}

func mySleep(duration time.Duration) {
	if duration <= 0 {
		return
	}
	to := time.After(duration)
	<-to
}
