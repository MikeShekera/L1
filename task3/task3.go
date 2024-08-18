package task3

// Mutex and WaitGroup
/*func main() {
	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	var mu sync.Mutex
	var sumOfSquares int

	for _, val := range arr {
		wg.Add(1)
		go getSquareOfNum(val, &wg, &mu, &sumOfSquares)
	}

	wg.Wait()
	fmt.Println(sumOfSquares)
}

func getSquareOfNum(val int, wg *sync.WaitGroup, mu *sync.Mutex, sum *int) {
	defer wg.Done()
	mu.Lock()
	*sum += val * val
	mu.Unlock()
}*/

//Реализация через Wg и Chan
/*func main() {
	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	ch := make(chan int, len(arr))

	for _, val := range arr {
		wg.Add(1)
		go getSquareOfNum(val, &wg, ch)
	}

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	var sum int
	for v := range ch {
		sum += v
	}
	fmt.Println(sum)
}

func getSquareOfNum(num int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	ch <- num * num
}*/
