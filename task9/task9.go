package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	multiplyChan := make(chan int)
	outChan := make(chan int)
	wg := sync.WaitGroup{}

	arr := []int{1, 2, 3, 4}
	wg.Add(2)
	go multiplier(multiplyChan, outChan, &wg)
	go receiver(outChan, &wg)

	for _, v := range arr {
		multiplyChan <- v
	}
	close(multiplyChan)
	wg.Wait()
}

func multiplier(multiplyChan chan int, outChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range multiplyChan {
		outChan <- val * 2
	}
	fmt.Println("Writer has stopped")
	close(outChan)
}

func receiver(outChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range outChan {
		fmt.Fprintln(os.Stdout, val)
	}
	fmt.Println("Reader has stopped")
}
