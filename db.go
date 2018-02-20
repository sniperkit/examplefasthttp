package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

// InitDB initial database
func InitDB(s string) (err error) {
	db, err = sql.Open("sqlite3", s)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	if err != nil {
		return errors.Wrap(err, "Not open DB in InitDB")
	}
	if err = db.Ping(); err != nil {
		return errors.Wrap(err, "Not ping DB in InitDB")
	}
	tx, err1 := db.Begin()
	if err1 != nil {
		return errors.Wrap(err, "Not starts a transaction in InitDB")
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
	_, err = db.Exec(sql)
	if err != nil {
		return errors.Wrap(err, "Error fo exec executes in InitDB")
	}

	stmt, err2 := tx.Prepare("INSERT INTO books(title, author, price) values(?, ?, ?)")
	if err2 != nil {
		return errors.Wrap(err, "Not prepared statement on transaction in InitDB")
	}
	defer stmt.Close()
	_, err = stmt.Exec("A cute love story", "Nidhi Agrawal", 1.32)
	_, err = stmt.Exec("Ultimate Pleasure", "Rachel G", 1.54)
	_, err = stmt.Exec("A Howl In The Night", "Lorelei Sutton", 2.02)
	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "Not commit the transaction")
	}
	return err
}

// ReturnBooks get field from table books
func ReturnBooks(db *sql.DB, Id int) (s *Data, err error) {
	s = new(Data)
	err = db.QueryRow("SELECT title, author, price FROM books WHERE id = ?", Id).Scan(&s.Title, &s.Author, &s.Price)
	return s, err
}
