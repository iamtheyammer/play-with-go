package main

import (
  "database/sql"
  _ "github.com/lib/pq"
  "fmt"
  "strconv"
)

type Dbrow struct {
  ID int
  UserID string
  Timestamp int32
}

func GetRow (db *sql.DB, id int) string {
  row := db.QueryRow("SELECT * FROM already_messaged_users WHERE id=$1", id)
  var item Dbrow
  err := row.Scan(&item.ID, &item.UserID, &item.Timestamp)
  switch err {
  case sql.ErrNoRows:
    return "No rows returned."
  case nil:
    return "ID: " + strconv.Itoa(item.ID) + ", User ID: " + item.UserID + ", Timestamp: " + strconv.Itoa(int(item.Timestamp))
  default:
    panic(err)
  }
}

func ConnectToDb() *sql.DB {
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
