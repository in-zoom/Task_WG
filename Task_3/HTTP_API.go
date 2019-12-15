package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	fmt.Fprintf(w, "Cats Service. Version 0.1")
}
