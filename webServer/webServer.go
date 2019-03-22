/*
Playing with an HTTP server
*/
package webServer

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This changes the response")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	http.ListenAndServe(":8000", mux)
}
