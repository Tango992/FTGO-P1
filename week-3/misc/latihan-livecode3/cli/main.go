package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"p1-latihan-lc3/config"
	"p1-latihan-lc3/handler"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		displayMenu()

		var userOption int
		scanner.Scan()
		if _, err := fmt.Sscanf(scanner.Text(), "%d", &userOption); err != nil {
			fmt.Println("Input harus berupa angka!")
			continue
		}

		fmt.Println()
		switch userOption {
		case 1:
			if err := handler.DisplayBooks(db); err != nil {
				log.Fatal(err)
			}

		case 2:
			if err := handler.DisplayBorrow(db); err != nil {
				log.Fatal(err)
			}

		case 3:
			if err := handler.AddBorrow(db); err != nil {
				log.Fatal(err)
			}

		case 4:
			fmt.Println()

		case 5:
			fmt.Println("Sampai jumpa!")
			os.Exit(0)

		default:
			fmt.Println("Input diluar jangkauan 1-5")
		}
	}
}

func displayMenu() {
	fmt.Printf("\nMENU\n")
	fmt.Println("1. Tampilkan Daftar Buku")
	fmt.Println("2. Tampilkan Daftar Peminjaman")
	fmt.Println("3. Pinjam Buku")
	fmt.Println("4. Kembali Buku")
	fmt.Println("5. Keluar")
	fmt.Print("Pilih menu (1/2/3/4/5): ")
}