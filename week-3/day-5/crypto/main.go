package main

import (
	"fmt"
	"os"
	"projectsDB/config"
	"projectsDB/entity"
	"projectsDB/handler"
)

func main() {
	config.Connect()
	defer config.DB.Close()

	for {
		fmt.Printf("\n1. Registrasi\n")
		fmt.Println("2. Login")
		fmt.Println("3. Keluar")
		fmt.Printf("Pilih opsi: ")

		var choice int
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println("Invalid input")
			continue
		}

		switch choice {
		case 1:
			var username, password string
			fmt.Println("Registrasi")
			fmt.Print("Masukkan username: ")
			fmt.Scanln(&username)
			fmt.Print("Masukkan password: ")
			fmt.Scanln(&password)

			user := entity.User {
				Username: username,
				Password: password,
			}

			if err := handler.Register(user); err != nil {
				fmt.Println("Kesalahan saat registrasi", err)
			} else {
				fmt.Println("Registrasi berhasil")
			}

		case 2:
			var username, password string
			fmt.Println("Login")
			fmt.Print("Masukkan username: ")
			fmt.Scanln(&username)
			fmt.Print("Masukkan password: ")
			fmt.Scanln(&password)

			user, authenticated, err := handler.Login(username, password)
			if err != nil {
				panic(err.Error())
			} else if authenticated {
				fmt.Printf("Login berhasil. Selamat datang %v!\n", user.Username)
			} else {
				fmt.Println("Login gagal")
			}
			
		case 3:
			fmt.Println("Exiting program")
			os.Exit(0)

		default:
			fmt.Println("Opsi di luar range!")
		}
	}
}