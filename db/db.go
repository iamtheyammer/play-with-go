/*
Playing with Postgres
https://www.calhoun.io/querying-for-a-single-record-using-gos-database-sql-package/
*/

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	type Result struct {
		ID        int
		UserID    string
		Timestamp int32
	}

	if err != nil {
		fmt.Println("ERROR:")
		fmt.Println(err)
	}

	fmt.Println("DB result:")
	rows, err := db.Query("SELECT * FROM already_messaged_users WHERE id=$1", 1)

	if err != nil {
		fmt.Println("DB ERROR:")
		fmt.Println(err)
	}
	rows.Next()

	var row Result

	switch err := rows.Scan(&row.ID, &row.UserID, &row.Timestamp); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(row.ID, row.UserID, row.Timestamp)
	default:
		panic(err)
	}
}
