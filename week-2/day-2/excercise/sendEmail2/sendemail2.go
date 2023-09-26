package main

import (
	"fmt"
	"sync"
	"time"
)

type Notification struct{
	UserID int
	Message string
}

func sendEmailAsync(UserID int, message string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Printf("Email notification sent to user %v: %v\n", UserID, message)
}

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	notifications := []Notification{
		{UserID: 101, Message: "Your order has been confirmed."},
		{UserID: 202, Message: "Your account has been created."},
		{UserID: 303, Message: "Your payment was successful."},
	}

	for _, notification := range notifications {
		wg.Add(1)
		go sendEmailAsync(notification.UserID, notification.Message, &wg)
	}

	fmt.Println("Main application continues...")

	wg.Wait()
	
	fmt.Println("Main application finished.")
	fmt.Printf("Time elapsed: %v\n", time.Since(now))
}