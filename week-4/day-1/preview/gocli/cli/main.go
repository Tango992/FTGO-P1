package main

import (
	"fmt"
	"log"
	"os"
	"p1-preview-w3/config"
	"p1-preview-w3/handler"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		log.Fatal("Missing command line argument!")
	}
	userOption := os.Args[1]

	switch userOption {
	case "books":
		if err := handler.JaneSmithBooks(db); err != nil {
			log.Fatal(err)
		}

	case "sales":
		if err := handler.TotalSales(db); err != nil {
			log.Fatal(err)
		}
		
	case "customers":
		if err := handler.CustomerOrders(db); err != nil {
			log.Fatal(err)
		}

	case "topauthor":
		if err := handler.TopAuthor(db); err != nil {
			log.Fatal(err)
		}

	default:
		fmt.Println("unknown command", userOption)
	}
}