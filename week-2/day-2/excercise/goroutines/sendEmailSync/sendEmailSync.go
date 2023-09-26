package main

import (
	"fmt"
	"time"
)

type Notification struct{
	UserID int
	Message string
}

func sendEmailSync(UserID int, message string) {
	time.Sleep(2 * time.Second)
	fmt.Printf("Email notification sent to user %v: %v\n", UserID, message)
}

func main() {
	now := time.Now()
	notifications := []Notification{
		{UserID: 101, Message: "Your order has been confirmed."},
		{UserID: 202, Message: "Your account has been created."},
		{UserID: 303, Message: "Your payment was successful."},
	}

	for _, notification := range notifications {
		sendEmailSync(notification.UserID, notification.Message)
	}
	fmt.Println("Main application finished.")
	fmt.Printf("Time elapsed: %v\n", time.Since(now))
}