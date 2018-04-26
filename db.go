package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"log"
	"github.com/valyala/fasthttp"
	"fmt"
	"encoding/json"
)

// Handler for mydb
type Handler struct {
	db *sql.DB
}

// initDB initial mydb
func (p *Handler) initDB() (*sql.DB, error) {
	var err error
	p.db, err = sql.Open("sqlite3", "./book.mydb")
	if err != nil {
		log.Fatal(err)
	}
	if err = p.db.Ping(); err != nil {
		log.Fatal(err)
	}
	tx, err1 := p.db.Begin()
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
	_, err = p.db.Exec(sql)
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
	return p.db, err
}

// Parse handle json request
func (p *Handler) parse(ctx *fasthttp.RequestCtx) {
	req := ctx.PostBody()
	if req == nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, "Error, please send a request body, 400\n")
		return
	}
	var r Data
	err := json.Unmarshal(req, &r)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, "Bad request, 400\n")
		return
	}
	err = p.db.QueryRow("SELECT title, author, price FROM books WHERE id = ?", r.ID).Scan(&r.Title, &r.Author,
		&r.Price)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		fmt.Fprintf(ctx, "Error, not found value, 404\n")
		return
	}
	b, err := json.Marshal(r)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "Error, internal server error, 500\n")
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	ctx.Write(b)
}

