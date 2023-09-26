package main

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"
)

// Without goroutine : +-100µs
// With goroutine : +-220µs

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	strings := []string{
		"As Soon As Possible",
		"Liquid-crystal Display",
		"Thank George It's Friday",
		"Loh Gak Bahaya Ta",
	}
	for _, string := range strings {
		wg.Add(1)
		go acronym(string, &wg)
	}
	wg.Wait()
	fmt.Printf("Elapsed time: %v\n", time.Since(now))
}

func acronym(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	var acronym string
	words := regexp.MustCompile("[-| ]").Split(s, -1)
	for _, word := range words {
		acronym += strings.ToUpper(string(word[0]))
	}
	fmt.Println(acronym)
}