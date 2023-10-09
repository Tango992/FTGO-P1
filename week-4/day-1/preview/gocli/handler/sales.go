package handler

import (
	"database/sql"
	"fmt"
)

func TotalSales(db *sql.DB) error {
	rows, err := db.Query(`
		SELECT t.name AS Type, ROUND(SUM(o.grand_total),2) As GrandTotal
		FROM orders o 
		JOIN books_price bp ON o.books_price_id = bp.id
		JOIN types t ON t.id = bp.type_id
		GROUP BY t.id
	`)
	if err != nil {
		return err
	}

	fmt.Println("Total sales for each book type:")
	for rows.Next() {
		var (
			bookType string
			total float32
		)

		if err := rows.Scan(&bookType, &total); err != nil {
			return err
		}
		fmt.Printf("%v : $%v\n", bookType, total)
	}
	return nil
}