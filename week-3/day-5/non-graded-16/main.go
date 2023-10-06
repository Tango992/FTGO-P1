package main

import (
	"fmt"
	"non-graded-16/config"
	"non-graded-16/entity"
	"non-graded-16/handler"
	"os"
	"regexp"
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
			var email, firstName, lastName, birth, password string
			emailRegex, _ := regexp.Compile(`^[\w-\.]+@(?:[\w-]+\.)+[\w-]{2,4}$`)
			birthRegex, _ := regexp.Compile(`^\d{4}\-(?:0[1-9]|1[012])\-(?:0[1-9]|[12][0-9]|3[01])$`)

			fmt.Println("Registrasi")
			for {
				fmt.Print("Masukkan email\t\t\t\t: ")
				fmt.Scanln(&email)
	
				if !emailRegex.MatchString(email) {
					fmt.Printf("Invalid email address!\n\n")
					continue
				}
				break
			}

			fmt.Print("Masukkan nama pertama\t\t\t: ")
			fmt.Scanln(&firstName)
			fmt.Print("Masukkan nama belakang (jika ada)\t: ")
			fmt.Scanln(&lastName)

			for {
				fmt.Print("Masukkan tanggal lahir (YYYY-MM-DD)\t: ")
				fmt.Scanln(&birth)
	
				if !birthRegex.MatchString(birth) {
					fmt.Printf("Invalid date of birth!\n\n")
					continue
				}
				break
			}

			fmt.Print("Masukkan password\t\t\t: ")
			fmt.Scanln(&password)

			user := entity.User {
				Email: email,
				FirstName: firstName,
				LastName: lastName,
				Birth: birth,
				Password: password,
			}

			if err := handler.Register(user); err != nil {
				fmt.Println("Kesalahan saat registrasi", err)
			} else {
				fmt.Println("Registrasi berhasil")
			}

		case 2:
			var email, password string
			fmt.Println("Login")
			fmt.Print("Masukkan email\t\t: ")
			fmt.Scanln(&email)
			fmt.Print("Masukkan password\t: ")
			fmt.Scanln(&password)

			user, authenticated := handler.Login(email, password)
			if authenticated {
				fmt.Printf("Login berhasil. Selamat datang %v!\n", user.FirstName)
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