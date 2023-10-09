package handler

import (
	"database/sql"
	"fmt"
)

func JaneSmithBooks(db *sql.DB) error {
	rows, err := db.Query(`
		SELECT b.name
		FROM books b
		JOIN authors a ON b.author_id = a.id
		WHERE a.name = "Jane Smith"
	`)
	if err != nil {
		return err
	}

	fmt.Println("Books by Jane Smith:")
	for rows.Next() {
		var book string

		if err := rows.Scan(&book); err != nil {
			return err
		}
		fmt.Println(book)
	}
	return nil
}