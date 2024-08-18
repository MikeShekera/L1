package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"reflect"
	"sync"
	"time"
)

func main() {
	ch := make(chan interface{})
	terminateCh := make(chan os.Signal, 1)
	signal.Notify(terminateCh, os.Interrupt)
	var wg sync.WaitGroup
	defer close(ch)
	defer close(terminateCh)

	var numOfWorkers int
	_, err := fmt.Scan(&numOfWorkers)
	if err != nil {
		log.Fatal(err)
	}
	wg.Add(numOfWorkers)
	for index := range numOfWorkers {
		go startWorker(ch, index+1, &wg)
	}

	go func() {
		<-terminateCh
		close(terminateCh)
		close(ch)
		wg.Wait()
		time.Sleep(1 * time.Second)
		os.Exit(1)
	}()

	for {
		time.Sleep(300 * time.Microsecond)
		ch <- generateRandomValue()
	}
}

func startWorker(ch chan interface{}, workerIndex int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		receivedData, ok := <-ch
		if ok {
			fmt.Println(
				"Worker №", workerIndex, " received: ", receivedData, " of type: ", reflect.TypeOf(receivedData),
			)
		} else {
			fmt.Println(fmt.Sprintf("Worker №%d has stopped", workerIndex))
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
