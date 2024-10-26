package main

import (
	"fmt"
	"net/http"
	"time"
)

func hora(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "Data e hora: %s", s)
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	http.HandleFunc("/hora", hora)

	fmt.Println("Listening on http://localhost:4444")
	http.ListenAndServe(":4444", nil)
}
