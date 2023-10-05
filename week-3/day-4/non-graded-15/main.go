package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/non_graded_15")
	checkErr(err)
	defer db.Close()

	err = db.Ping()
	checkErr(err)

	// release2(db)
	// release3(db)
	// release4a(db)
	release4b(db)
}

func release2(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE Employees (employee_id INT PRIMARY KEY AUTO_INCREMENT, first_name VARCHAR(50), last_name VARCHAR(50), position VARCHAR(50))")
	checkErr(err)
	_, err = db.Exec("CREATE TABLE MenuItems (item_id INT PRIMARY KEY AUTO_INCREMENT, name VARCHAR(100), description VARCHAR(255), price DECIMAL(10, 2), category VARCHAR(50))")
	checkErr(err)
	_, err = db.Exec("CREATE TABLE Orders (order_id INT PRIMARY KEY AUTO_INCREMENT, table_number INT, employee_id INT, order_date DATE, status VARCHAR(50), FOREIGN KEY (employee_id) REFERENCES Employees(employee_id))")
	checkErr(err)
	_, err = db.Exec("CREATE TABLE OrderItems (order_item_id INT PRIMARY KEY AUTO_INCREMENT, order_id INT, item_id INT, quantity INT, subtotal DECIMAL(10, 2), FOREIGN KEY (order_id) REFERENCES Orders(order_id), FOREIGN KEY (item_id) REFERENCES MenuItems(item_id))")
	checkErr(err)
	_, err = db.Exec("CREATE TABLE Payments (payment_id INT PRIMARY KEY AUTO_INCREMENT, order_id INT, payment_date DATE, payment_method VARCHAR(50), total_amount DECIMAL(10, 2), FOREIGN KEY (order_id) REFERENCES Orders(order_id))")
	checkErr(err)

	fmt.Println("All tables created successfully!")
}

func release3(db *sql.DB) {
	employees := []map[string]string{
		{"first": "John", "last": "Doe", "position": "Manager"},
		{"first": "Jane", "last": "Smith", "position": "Waitress"},
		{"first": "John", "last": "Doe", "position": "Cook"},
	}
	for _, employee := range employees {
		_, err := db.Exec("INSERT INTO Employees (first_name, last_name, position) VALUES (?,?,?)", employee["first"], employee["last"], employee["position"])
		checkErr(err)
	}

	menus := []map[string]any{
		{"name": "Spaghetti Carbonara", "description": "Traditional Italian dish with eggs, cheese, pancetta, and pepper.", "price": 12.50, "category": "Main Course"},
		{"name": "Caesar Salad", "description": "Fresh lettuce with Caesar dressing, croutons and parmesan.", "price": 6.00, "category": "Starter"},
		{"name": "Tiramisu", "description": "Classic Italian dessert with coffee-soaked sponge and mascarpone.", "price": 5.50, "category": "Dessert"},
	}
	for _, menu := range menus {
		_, err := db.Exec("INSERT INTO MenuItems (name, description, price, category) VALUES (?,?,?,?)", menu["name"], menu["description"], menu["price"], menu["category"])
		checkErr(err)
	}

	orders := []map[string]any{
		{"table_number": 10, "employee_id": 1, "order_date": "2023-08-09", "status": "Placed"},
		{"table_number": 5, "employee_id": 2, "order_date": "2023-08-09", "status": "Completed"},
	}
	for _, order := range orders {
		_, err := db.Exec("INSERT INTO Orders (table_number, employee_id, order_date, status) VALUES (?,?,?,?)", order["table_number"], order["employee_id"], order["order_date"], order["status"])
		checkErr(err)
	}

	orderItems := []map[string]any{
		{"order_id": 1, "item_id": 1, "quantity": 2, "subtotal": 25.00},
		{"order_id": 2, "item_id": 2, "quantity": 1, "subtotal": 6.00},
		{"order_id": 2, "item_id": 3, "quantity": 1, "subtotal": 5.50},
	}
	for _, orderItem := range orderItems {
		_, err := db.Exec("INSERT INTO OrderItems (order_id, item_id, quantity, subtotal) VALUES (?,?,?,?)", orderItem["order_id"], orderItem["item_id"], orderItem["quantity"], orderItem["subtotal"])
		checkErr(err)
	}

	payments := []map[string]any{
		{"order_id": 1, "payment_date": "2023-08-09", "payment_method": "Credit Card", "total_amount": 25.00},
		{"order_id": 2, "payment_date": "2023-08-09", "payment_method": "Cash", "total_amount": 11.50},
	}
	for _, payment := range payments {
		_, err := db.Exec("INSERT INTO Payments (order_id, payment_date, payment_method, total_amount) VALUES (?,?,?,?)", payment["order_id"], payment["payment_date"], payment["payment_method"], payment["total_amount"])
		checkErr(err)
	}
	fmt.Println("All sample data created successfully!")
}

func release4a(db *sql.DB) {
	rows, err := db.Query("SELECT o.order_id, o.table_number, o.order_date, o.status, e.first_name, e.last_name, e.position FROM Orders o JOIN Employees e ON o.employee_id = e.employee_id;")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var (
			order_id int
			table_number int
			order_date string
			status string
			first_name string
			last_name string
			position string
		)
		if err = rows.Scan(&order_id, &table_number, &order_date, &status, &first_name, &last_name, &position); err != nil {
			panic(err.Error())
		}

		fmt.Printf("%v\t%v\t%v\t%v\t%v %v\t%v\n", order_id, table_number, order_date, status, first_name, last_name, position)
	}
}

func release4b(db *sql.DB) {
	rows, err := db.Query("SELECT o.order_id, o.table_number, o.order_date, o.status, e.first_name, e.last_name, e.position, GROUP_CONCAT(mi.name) AS items, p.total_amount FROM Orders o JOIN Employees e ON o.employee_id = e.employee_id JOIN Payments p ON o.order_id = p.order_id JOIN OrderItems oi ON o.order_id = oi.order_id JOIN MenuItems mi ON oi.item_id = mi.item_id GROUP BY oi.order_id;")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var (
			order_id int
			table_number int
			order_date string
			status string
			first_name string
			last_name string
			position string
			items string
			total_amount float32
		)
		if err = rows.Scan(&order_id, &table_number, &order_date, &status, &first_name, &last_name, &position, &items, &total_amount); err != nil {
			panic(err.Error())
		}

		fmt.Println(order_id, table_number, order_date, status, first_name, last_name, position, items, total_amount)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}