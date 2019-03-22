package main

import (
	"encoding/json"
	"fmt"
)

type row struct {
	ID        int
	UserID    string
	Timestamp int64
}

func main() {
	fmt.Println("Hello!")
	row := &row{1, "eigjwoigjo", 1553272944853}
	jRow, _ := json.Marshal(row)
	fmt.Println(string(jRow))
}
