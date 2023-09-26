package main

import (
	"fmt"
	"sync"
	"time"
)

type Log struct {
	Type string
	Message string
}

func displayLog(logType string, logMessage string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Printf("[%v] - %v\n", logType, logMessage)
}

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	logs := []Log{
		{Type: "WARN", Message: "Memory usage is high"},
		{Type: "INFO", Message: "User 123 logged in"},
		{Type: "ERROR", Message: "Failed to fetch data from API"},
	}

	for _, log := range logs {
		wg.Add(1)
		go displayLog(log.Type, log.Message, &wg)
	}

	fmt.Println("Main application continues...")

	wg.Wait()
	
	fmt.Println("Main application finished.")
	fmt.Printf("Time elapsed: %v\n", time.Since(now))

}