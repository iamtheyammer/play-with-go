package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	//"fmt"
)

/*
Hello - /route
*/
func Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This changes the response")
}

/*
ShowDbResult - /db route
*/
func ShowDbResult(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query()["id"] == nil {
		io.WriteString(w, "{\"error\":true,\"msg\":\"No ID.\"}")
		return
	}
	id, err := strconv.Atoi(r.URL.Query()["id"][0])
	fmt.Println(id, "ID")
	if err != nil {
		io.WriteString(w, "{\"error\":true,\"msg\":\"Invalid ID.\"}")
		return
	}
	// fmt.Println(r.URL.Query()["id"][0])
	result := GetRow(id)
	io.WriteString(w, result)
}
