package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hello)
	mux.HandleFunc("/db", ShowDbResult)
	http.ListenAndServe(":8000", mux)
}
