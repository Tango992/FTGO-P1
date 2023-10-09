package handler

import (
	"database/sql"
	"fmt"
)

func CustomerOrders(db *sql.DB) error {
	rows, err := db.Query(`
		WITH all_transactions AS 
			(SELECT c.name, COUNT(c.id) AS Transactions
			FROM orders o
			JOIN customers c ON c.id = o.customer_id
			GROUP BY c.id)
		SELECT *
		FROM all_transactions
		WHERE Transactions > 1
	`)
	if err != nil {
		return err
	}

	fmt.Println("Customers who ordered more than one book:")
	for rows.Next() {
		var (
			customerName string
			transaction float32
		)

		if err := rows.Scan(&customerName, &transaction); err != nil {
			return err
		}
		fmt.Printf("%v : %v orders\n", customerName, transaction)
	}
	return nil
}