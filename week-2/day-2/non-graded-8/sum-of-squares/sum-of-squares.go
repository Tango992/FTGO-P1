package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := calculateSumOfSquares(numbers)
	fmt.Println("Sum of squares:", result)
	fmt.Println("Elapsed time:", time.Since(now))
}

func calculateSumOfSquares(numbers []int) int {
	var temp int
	var wg sync.WaitGroup

	for _, number := range numbers {
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			temp += int(math.Pow(float64(i), 2))
		}(number)
	}
	wg.Wait()
	return temp
}
