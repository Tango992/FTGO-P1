package handler

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"time"
)

func AddBorrow(db *sql.DB) error {
	DisplayBooks(db)

	scanner := bufio.NewScanner(os.Stdin)
	var bukuId, customerId int

	for {
		fmt.Print("Masukkan ID Buku yang ingin dipinjam: ")
		scanner.Scan()
		if _, err := fmt.Sscanf(scanner.Text(), "%d", &bukuId); err != nil {
			fmt.Println("Invalid input")
			continue
		}
		break
	}

	for {
		fmt.Print("Masukkan ID pelanggan: ")
		scanner.Scan()
		if _, err := fmt.Sscanf(scanner.Text(), "%d", &customerId); err != nil {
			fmt.Println("Invalid input")
			continue
		}
		break
	}

	returnDate := time.Now().AddDate(0,0,7)

	result, err := db.Exec(`INSERT INTO peminjaman (anggota_id, tanggal_kembali) VALUES (?,?)`, customerId, returnDate)
	if err != nil {
		return err
	}

	peminjamanId, _ := result.LastInsertId()
	_, err = db.Exec(`INSERT INTO detail_peminjaman (peminjaman_id, buku_id, jumlah_buku) VALUES (?,?,?)`, peminjamanId, bukuId, 1)
	if err != nil {
		return err
	}
	fmt.Println("Berhasil meminjam!")
	return nil
}