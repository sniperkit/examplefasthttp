package mydb

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"log"
)

// Handler for mydb
type Handler struct {
	db *sql.DB
}

// InitDB initial mydb
func (d *Handler) InitDB() (*sql.DB, error) {
	var err error
	d.db, err = sql.Open("sqlite3", "./book.mydb")
	if err != nil {
		log.Fatal(err)
	}
	if err = d.db.Ping(); err != nil {
		log.Fatal(err)
	}
	tx, err1 := d.db.Begin()
	if err1 != nil {
		log.Fatal(err)
	}
	sql := `
	drop table if exists books;
	create table if not exists books (
    id integer primary key autoincrement not null,
    title varchar(255) not null,
    author varchar(255) not null,
    price real not null
	);
	create unique index books_id_uindex on books (id);
	delete from books;
`
	_, err = d.db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}

	stmt, err2 := tx.Prepare("INSERT INTO books(title, author, price) values(?, ?, ?)")
	if err2 != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec("A cute love story", "Nidhi Agrawal", 1.32)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec("Ultimate Pleasure", "Rachel G", 1.54)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec("A Howl In The Night", "Lorelei Sutton", 2.02)
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return d.db, err
}
