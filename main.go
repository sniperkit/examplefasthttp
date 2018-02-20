package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)

// Data the structure of database
type Data struct {
	Id     int64   `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float32 `json:"price"`
}

var db *sql.DB

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Catch error: %v", r)
			os.Exit(1)
		}
	}()
	err := InitDB("./book.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	router := fasthttprouter.New()
	router.POST("/", Parse)
	log.Fatal(fasthttp.ListenAndServe(":4000", router.Handler))
}

// Parse handle json request
func Parse(ctx *fasthttp.RequestCtx) {
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
	err = db.QueryRow("SELECT title, author, price FROM books WHERE id = ?", r.Id).Scan(&r.Title, &r.Author,
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
