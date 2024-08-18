package task2

//Варианты реализации: WaitGroup, вычитывание данных из канала

//WaitGroup
/*func main() {
	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	for _, val := range arr {
		wg.Add(1)
		go getSquareOfNum(val, &wg)
	}

	wg.Wait()
}

func getSquareOfNum(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(os.Stdout, num*num)
}*/

// Вычитывание данных из канала
/*func main() {
	arr := []int{2, 4, 6, 8, 10}
	ch := make(chan bool, len(arr))
	defer close(ch)
	for range len(arr) {
		ch <- true
	}

	for _, val := range arr {
		go getSquareOfNum(val, ch)
	}

	for len(ch) != 0 {

	}
}

func getSquareOfNum(num int, ch chan bool) {
	fmt.Fprintln(os.Stdout, num*num)
	<-ch
}*/
