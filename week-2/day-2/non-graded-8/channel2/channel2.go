package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	now := time.Now()
	resultSumSquare := make(chan int)
	resultSquareSum := make(chan int)

	go SumSquare(100, resultSumSquare)
	go SquareSum(100, resultSquareSum)

	fmt.Printf("Semua angka dijumlahkan lalu dikuadratkan: %v\n", <- resultSumSquare)
	fmt.Printf("Semua angka dijumlahkan lalu dikuadratkan: %v\n", <- resultSquareSum)
	fmt.Printf("Elapsed time: %v\n", time.Since(now))
}

func SumSquare(num int, result chan int) {
	var temp int
	for i := 1; i <= 100; i++ {
		temp += i
	}
	temp = int(math.Pow(float64(temp), 2))
	result <- temp
}

func SquareSum(num int, result chan int) {
	var temp int
	for i := 1; i <= 100; i++ {
		temp += int(math.Pow(float64(i), 2))
	}
	result <- temp
}