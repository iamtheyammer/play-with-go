package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

/*
Dbrow - a struct for rows from the database; not used at the moment
*/
type Dbrow struct {
	ID        int    `json:"id"`
	UserID    string `json:"user_id"`
	Timestamp int32  `json:"timestamp"`
}

/*
Error - a struct we'll use to output errors as json
*/
type Error struct {
	Err bool   `json:"error"`
	Msg string `json:"msg"`
}

var globalDB *sql.DB = connectToDb()

/*
GetRow - gets a row by passing in the database from ConnectToDb and the Row ID
*/
func GetRow(id int) string {
	row := globalDB.QueryRow("SELECT * FROM already_messaged_users WHERE id=$1", id)
	fmt.Println("GetRow ID", id)
	var item Dbrow
	err := row.Scan(&item.ID, &item.UserID, &item.Timestamp)
	switch err {
	case sql.ErrNoRows:
		err := &Error{true, "No rows for that ID."}
		jRow, jError := json.Marshal(err)
		checkError(jError)
		return string(jRow)
	case nil:
		res := &Dbrow{item.ID, item.UserID, item.Timestamp}
		fmt.Println(item.UserID)
		jRes, jError := json.Marshal(res)
		checkError(jError)
		fmt.Println(string(jRes))
		return string(jRes)
	default:
		panic(err)
	}
}

/*
ConnectToDb - Gives you a database object for using GetRow
*/
func connectToDb() *sql.DB {
	connStr := "user=sammendelson dbname=sammendelson sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	checkError(err)
	return db
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("DB ERROR: %s", err)
	}
}
