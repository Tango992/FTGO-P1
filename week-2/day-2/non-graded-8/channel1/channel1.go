package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	ch := make(chan string)
	chSum := make(chan int)
	chIteration := make(chan int, 100)
	var even, odd int
	
	go fizzBuzz(100, ch, chSum, chIteration)

	for i := 0; i < 100; i++ {
		fmt.Println(<- ch)
	}

	for i := 0; i < 100; i++{
		temp := <- chIteration
		if temp % 2 == 0 {
			even++
		} else {
			odd++
		}
	}
	fmt.Printf("\nOdd numbers: %v\n", odd)
	fmt.Printf("Even numbers: %v\n", even)
	fmt.Println("Sum of all iteration:", <- chSum)
	fmt.Printf("Elapsed time: %v\n", time.Since(now))
}

func fizzBuzz(n int, ch chan string, chSum chan int, chIteration chan int) {
	var temp int
	for i := 0; i < n; i++ {
		if i == 0 {
			ch <- "0"
		} else if i%3 == 0 && i%5 == 0 {
			ch <- fmt.Sprintf("%vFizzBuzz", i)
		} else if i%3 == 0 {
			ch <- fmt.Sprintf("%vFizz", i)
		} else if i%5 == 0 {
			ch <- fmt.Sprintf("%vBuzz", i)
		} else {
			ch <- fmt.Sprint(i)
		}
		chIteration <- i
		temp += i
	}
	chSum <- temp
}