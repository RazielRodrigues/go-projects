package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

func hora(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "Data e hora: %s", s)
}

type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "portugaldocs:portugaldocs@/portugaldocs")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var u Usuario
	row := db.QueryRow("select id, nome from usuarios where id = ?", id)

	row.Scan(&u.ID, &u.Nome)

	jsonResponse, _ := json.Marshal(u)

	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "portugaldocs:portugaldocs@/portugaldocs")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var u []Usuario
	rows, _ := db.Query("select id, nome from usuarios")

	for rows.Next() {
		var u1 Usuario
		rows.Scan(&u1.ID, &u1.Nome)
		u = append(u, u1)
	}

	jsonResponse, _ := json.Marshal(u)

	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func HandlerUsuario(w http.ResponseWriter, r *http.Request) {

	sid := strings.TrimPrefix(r.URL.Path, "/usuario/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)
	case r.Method == "GET":
		usuarioTodos(w, r)
	default:
		fmt.Fprintf(w, "not found")
	}

}

func main() {

	http.Handle("/", http.FileServer(http.Dir("/public")))
	http.HandleFunc("/hora", hora)
	http.HandleFunc("/usuario/", HandlerUsuario)

	fmt.Println("Listening on http://localhost:4444")
	http.ListenAndServe(":4444", nil)
}
