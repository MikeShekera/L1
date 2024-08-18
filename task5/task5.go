package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

func main() {
	var duration uint
	fmt.Scan(&duration)
	ch := make(chan interface{})
	var wg sync.WaitGroup

	wg.Add(1)
	go startReader(ch, &wg)

	to := time.After(time.Duration(duration) * time.Second)
loop:
	for {
		select {
		case <-to:
			break loop
		default:
			ch <- generateRandomValue()
		}
	}
	close(ch)
	wg.Wait()
}

func startReader(ch chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		receivedData, ok := <-ch
		if ok {
			fmt.Println(
				" received: ", receivedData, " of type: ", reflect.TypeOf(receivedData),
			)
		} else {
			fmt.Println("Reader has stopped")
			return
		}
	}
}

func generateRandomValue() interface{} {
	switch rand.Intn(4) {
	case 0:
		return rand.Int()
	case 1:
		return rand.Float64()
	case 2:
		return rand.Intn(2) == 0
	case 3:
		return rune(rand.Intn(0x10FFFF))
	default:
		return nil
	}
}
