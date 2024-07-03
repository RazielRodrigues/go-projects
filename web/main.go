package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/bcrypt"
)

func withoutRouter() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL %s", r.URL.Path)
	})

	fs := http.FileServer(http.Dir("static/1"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}

func withRouter() (router *mux.Router) {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	r.HandleFunc("/about", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "URL %s", req.URL.Path)

	}).Methods("POST")
	/* .Host("www.mybookstore.com") .Schemes("https")
	 */router = r

	test := r.PathPrefix("/test").Subrouter()
	test.HandleFunc("/ok", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "URL %s", req.URL.Path)

	}).Methods("GET")

	return router
}

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "secret"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)

	/* 	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
	   		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

	   		for {
	   			// Read message from browser
	   			msgType, msg, err := conn.ReadMessage()
	   			if err != nil {
	   				return
	   			}

	   			// Print the message to the console
	   			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

	   			// Write message back to browser
	   			if err = conn.WriteMessage(msgType, msg); err != nil {
	   				return
	   			}
	   		}
	   	})

	   	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	   		http.ServeFile(w, r, "websockets.html")
	   	}) */

	/* 	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
	   		var user User
	   		json.NewDecoder(r.Body).Decode(&user)

	   		fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
	   	})

	   	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
	   		peter := User{
	   			Firstname: "John",
	   			Lastname:  "Doe",
	   			Age:       25,
	   		}

	   		json.NewEncoder(w).Encode(peter)
	   	}) */

	/* 	http.HandleFunc("/secret", secret)
	   	http.HandleFunc("/login", login)
	   	http.HandleFunc("/logout", logout) */

	http.ListenAndServe(":80", nil)

	/* 	tmpl := template.Must(template.ParseFiles("forms.html"))

	   	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	   		if r.Method != http.MethodPost {
	   			tmpl.Execute(w, nil)
	   			return
	   		}

	   		details := ContactDetails{
	   			Email:   r.FormValue("email"),
	   			Subject: r.FormValue("subject"),
	   			Message: r.FormValue("message"),
	   		}

	   		// do something with details
	   		_ = details

	   		fmt.Println(details)

	   		tmpl.Execute(w, struct{ Success bool }{true})
	   	})

	   	http.ListenAndServe(":80", nil) */

	/* 	tmpl := template.Must(template.ParseFiles("layout.html"))
	   	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	   		data := TodoPageData{
	   			PageTitle: "My TODO list",
	   			Todos: []Todo{
	   				{Title: "Task 1", Done: false},
	   				{Title: "Task 2", Done: true},
	   				{Title: "Task 3", Done: true},
	   			},
	   		}
	   		tmpl.Execute(w, data)
	   	})
	   	http.ListenAndServe(":80", nil) */

	//x := withRouter()

	/* 	db, err := sql.Open("mysql", "root:root@(172.18.0.2:3306)/golang?parseTime=true")

		if err != nil {
			log.Fatal(err)
		}
		if err := db.Ping(); err != nil {
			log.Fatal(err)
		}

		{ // Create a new table
			query := `
	            CREATE TABLE users (
	                id INT AUTO_INCREMENT,
	                username TEXT NOT NULL,
	                password TEXT NOT NULL,
	                created_at DATETIME,
	                PRIMARY KEY (id)
	            );`

			if _, err := db.Exec(query); err != nil {
				log.Fatal(err)
			}
		}

		{ // Insert a new user
			username := "johndoe"
			password := "secret"
			createdAt := time.Now()

			result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
			if err != nil {
				log.Fatal(err)
			}

			id, err := result.LastInsertId()
			fmt.Println(id)
		}

		{ // Query a single user
			var (
				id        int
				username  string
				password  string
				createdAt time.Time
			)

			query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
			if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
				log.Fatal(err)
			}

			fmt.Println(id, username, password, createdAt)
		}

		{ // Query all users
			type user struct {
				id        int
				username  string
				password  string
				createdAt time.Time
			}

			rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()

			var users []user
			for rows.Next() {
				var u user

				err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
				if err != nil {
					log.Fatal(err)
				}
				users = append(users, u)
			}
			if err := rows.Err(); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%#v", users)
		}

		{
			_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
			if err != nil {
				log.Fatal(err)
			}
		} */
}
