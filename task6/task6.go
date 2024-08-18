package task6

import (
	"context"
	"fmt"
	"time"
)

// Stop chan
func worker1(stopChan chan struct{}) {
	for {
		if _, ok := <-stopChan; !ok {
			fmt.Println("Goroutine stopped")
			return
		} else {
			//Work
		}
	}
}

// Bool stopper
func worker2(closeBool *bool) {
	for {
		if *closeBool {
			fmt.Println("Goroutine stopped")
			return
		} else {
			//Work
		}
	}
}

// Context
func worker3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine stopped")
			return
		default:
			//Work
		}
	}
}

// Timer
func worker4(lifeDuration time.Duration) {
	timer := time.NewTimer(lifeDuration * time.Second)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			fmt.Println("Goroutine stopped")
			return
		default:
			//Work
		}
	}
}

// Callback Func
func worker5(callbackFunc func() bool) {
	for {
		if callbackFunc() {
			fmt.Println("Goroutine stopped")
			return
		}
		//Work
	}
}
