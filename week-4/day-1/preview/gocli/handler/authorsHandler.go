package handler

import (
	"database/sql"
	"fmt"
)

func TopAuthor(db *sql.DB) error {
	rows, err := db.Query(`
		SELECT a.name, ROUND(SUM(o.grand_total), 2) AS Earnings
		FROM orders o
		JOIN books_price bp ON o.books_price_id = bp.id
		JOIN books b ON b.id = bp.book_id
		JOIN authors a ON a.id = b.author_id
		GROUP BY a.name
		ORDER BY Earnings DESC LIMIT 1
	`)
	if err != nil {
		return err
	}

	fmt.Println("Top earning author:")
	for rows.Next() {
		var (
			author string
			earnings float32
		)

		if err := rows.Scan(&author, &earnings); err != nil {
			return err
		}
		fmt.Printf("%v with earnings: $%v\n", author, earnings)
	}
	return nil
}