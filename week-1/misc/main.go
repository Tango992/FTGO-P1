package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Book struct {
	Title     string
	Author    string
	Available bool
}

var books = []Book{
	{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Available: true},
	{Title: "To Kill a Mockingbird", Author: "Harper Lee", Available: true},
	{Title: "1984", Author: "George Orwell", Available: false},
}

func main() {
	fmt.Println("\nWelcome to the Library Book Rental System!")
	showMenu()
}

func showMenu() {
	fmt.Println("\nMain Menu:")
	fmt.Println("1. Show available books")
	fmt.Println("2. Rent a book")
	fmt.Println("3. Exit Program")

	for {
		var choice uint8
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("\nPlease enter your choice (1/2/3): ")
		scanner.Scan()
		_, err := fmt.Sscanf(scanner.Text(), "%d\n", &choice)

		if err != nil {
			fmt.Println("Input is not an integer!")
			continue
		}

		switch choice {
		case 1:
			showAvailableBooks(books)
		case 2:
			rentBook(books)
		case 3:
			fmt.Println("Thank you for using the Library Book Rental System. Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Input is out of range (1-3)!")
			continue
		}
	}
}

func showAvailableBooks(books []Book) {
	fmt.Println("\nAvailable Books:")
	for i, v := range books {
		fmt.Printf("%v. Title: %v, Author: %v, Available: %v\n", i+1, v.Title, v.Author, v.Available)
	}
}

func rentBook(books []Book) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("\nEnter the title of the book you want to rent: ")
		scanner.Scan()
		input := strings.ToLower(scanner.Text())
		bookExist := false
	
		for _, v := range books {
			if input == strings.ToLower(v.Title) {
				bookExist = true
				if v.Available {
					fmt.Printf("\nBook %q has been successfully rented!\n", v.Title)
				} else {
					fmt.Printf("\nBook %q is not available for rent!\n", v.Title)
				}
			}
		}

		if !bookExist {
			fmt.Println("Book input does not match!")
			continue
		}

		fmt.Print("\nDo you want to rent another book? (yes/no): ")
		scanner.Scan()
		cont := scanner.Text()
		if cont == "yes" {
			continue
		} else {
			fmt.Println("Thank you for using the Library Book Rental System. Goodbye!")
			os.Exit(0)
		}
	}
}