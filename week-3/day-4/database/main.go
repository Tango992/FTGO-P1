package main

import (
	"database/sql"
	"fmt"
	"context"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/sesi7")
	catch(err)
	defer db.Close()

	err = db.Ping()
	catch(err)
	
	// _, err = db.Exec("CREATE TABLE IF NOT EXISTS test (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(50));")
	// catchErr(err)

	// _, err = db.Exec("INSERT INTO test (name) VALUES (?), (?), (?);", "John", "Jane", "Baz")
	// catch(err)

	// rows, err := db.Query("SELECT * FROM test")
	// catch(err)
	// defer rows.Close()

	// for rows.Next() {
	// 	var (
	// 		id int
	// 		name string
	// 	)
	// 	err = rows.Scan(&id, &name)
	// 	catch(err)

	// 	fmt.Println(id, name)
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	_, err = db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS user (id INT AUTO_INCREMENT PRIMARY KEY NOT NULL, name VARCHAR(50));")
	catch(err)

	_, err = db.ExecContext(ctx, "INSERT INTO user (name) VALUES (?),(?)", "Jane", "John")
	catch(err)

	rows, err := db.QueryContext(ctx, "SELECT * FROM user")
	catch(err)
	defer rows.Close()

	for rows.Next() {
		var (
			id int
			name string
		)
		err = rows.Scan(&id, &name)
		catch(err)

		fmt.Println(id, name)
	}
}

func catch(err error) {
	if err != nil {
		panic(err.Error())
	}
}
