package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func createTable(db *sql.DB) {
	exec(db, "CREATE DATABASE IF NOT EXISTS portugaldocs")
	exec(db, "USE portugaldocs")
	// exec(db, "DROP TABLE IF EXISTS usuarios")
	exec(db, `
		CREATE TABLE IF NOT EXISTS usuarios (
			id int(11) NOT NULL AUTO_INCREMENT,
			nome varchar(50) NOT NULL,
			PRIMARY KEY (id)
		)
	`)
}

func create(db *sql.DB) {
	exec(db, "USE portugaldocs")

	// Com transactions, sempre pensar em ajustar os fluxos SCID no banco de dados
	tx, _ := db.Begin()

	stmt, err := db.Prepare("insert into usuarios(nome) values(?)")
	res, _ := stmt.Exec("Raziel")

	if err != nil {
		tx.Rollback()
		panic(err)
	}

	tx.Commit()

	id, _ := res.LastInsertId()
	rows, _ := res.RowsAffected()
	fmt.Println(
		id,
		rows,
	)
}

type usuario struct {
	id   int
	nome string
}

func read(db *sql.DB) {
	exec(db, "USE portugaldocs")

	rows, _ := db.Query("select id, nome from usuarios")
	defer rows.Close()

	data := []usuario{}
	for rows.Next() {
		u := usuario{}
		rows.Scan(&u.id, &u.nome)
		data = append(data, u)
	}
	fmt.Println(data)
}

func update(db *sql.DB) {
	exec(db, "USE portugaldocs")

	// Com transactions, sempre pensar em ajustar os fluxos SCID no banco de dados
	tx, _ := db.Begin()

	stmt, err := db.Prepare("update usuarios set nome = ? where id = ?")
	res, _ := stmt.Exec("Jaiara", 1)

	if err != nil {
		tx.Rollback()
		panic(err)
	}

	tx.Commit()

	id, _ := res.LastInsertId()
	rows, _ := res.RowsAffected()
	fmt.Println(
		id,
		rows,
	)
}

func delete(db *sql.DB) {
	exec(db, "USE portugaldocs")

	// Com transactions, sempre pensar em ajustar os fluxos SCID no banco de dados
	tx, _ := db.Begin()

	stmt, err := db.Prepare("delete from usuarios where id = ?")
	res, _ := stmt.Exec("7")

	if err != nil {
		tx.Rollback()
		panic(err)
	}

	tx.Commit()

	id, _ := res.LastInsertId()
	rows, _ := res.RowsAffected()
	fmt.Println(
		id,
		rows,
	)
}

func main() {
	db, err := sql.Open("mysql", "portugaldocs:portugaldocs@/portugaldocs")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	createTable(db)

	// C.R.U.D
	create(db)
	read(db)
	update(db)
	delete(db)
}
