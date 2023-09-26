package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Without goroutine: +- 30µs
// With goroutine	: +- 100µs

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	words := []string{"exampleword", "john", "hacktiv8"}

	for _, word := range words {
		wg.Add(1)
		go scrabble(word, &wg)
	}
	wg.Wait()
	fmt.Printf("Elapsed time: %v\n", time.Since(now))
}

func scrabble(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	var score int
	for _, c := range s {
		switch strings.ToUpper(string(c)) {
		case "A","E","I","O","U","L","N","R","S","T":
			score += 1
		case "D","G":
			score += 2
		case "B","C","M","P":
			score += 3
		case "F","H","V","W","Y":
			score += 4
		case "K":
			score += 5
		case "J","X":
			score += 8
		case "Q","Z":
			score += 10
		}
	}
	fmt.Printf("%v score: %v\n", s, score)
}