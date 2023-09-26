package main

import (
	"fmt"
	"sync"
)

func factorial(n int, wg *sync.WaitGroup, resultChan chan string) {
	defer wg.Done()

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	resultChan <- fmt.Sprintf("%v factorial: %v\n", n, result)
}

func main() {
	numbers := []int{5, 7, 10, 3, 8}

	var wg sync.WaitGroup
	resultChan := make(chan string, len(numbers))
	for _, n := range numbers {
		wg.Add(1)
		go factorial(n, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for r := range resultChan{
		fmt.Printf(r)
	}
}