package handler

import (
	"database/sql"
	"fmt"
	"p1-latihan-lc3/entity"
)

func DisplayBooks(db *sql.DB) error {
	rows, err := db.Query(`
		SELECT buku_id, judul, pengarang, tahun_terbit, stok 
		FROM buku
	`)
	if err != nil {
		return err
	}
	defer rows.Close()

	var books []entity.Buku
	for rows.Next() {
		var book entity.Buku

		if err := rows.Scan(&book.Buku_id, &book.Judul, &book.Pengarang, &book.Tahun_terbit, &book.Stok); err != nil {
			return err
		}
		books = append(books, book)
	}
	BooksParser(books)

	return nil
}

func BooksParser(books []entity.Buku) {
	fmt.Println("Daftar Buku:")
	for _, book := range books {
		fmt.Printf("%v. %v by %v (%v)\n", book.Buku_id, book.Judul, book.Pengarang, book.Pengarang)
	}
}