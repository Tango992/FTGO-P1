package main

import (
	"fmt"
	"strings"
	"time"
)

func scheduleTask(task func(), i int) {
	task()
}

func main() {
	go scheduleTask(fibonacci, 1)
	go scheduleTask(star, 2)
	go scheduleTask(box, 3)
	fmt.Println("Main application continues...")
	time.Sleep(1 * time.Second)
	fmt.Println("All task completed.")
}

func fibonacci() {
	n := 100
	var a uint64 = 0
	var b uint64 = 1
	var c uint64
	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	fmt.Printf("%vth fibbonaci number: %v\n", n, b)
}

func star() {
	n := 5
	mid := n / 2
	for i := 0; i < n; i++ {
		space := mid - i
		if space < 0 {
			space *= -1
		}
		star := 2 * (mid - space) + 1
		fmt.Print(strings.Repeat(" ", space))
		fmt.Print(strings.Repeat("*", star))
		fmt.Println()
	}
}

func box() {
	N := 10
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == 0 || i == N-1 || j == 0 || j == N-1 {
				fmt.Print("*")
			} else if j > 0 && j < N-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}