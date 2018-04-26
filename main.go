package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Catch error: %q \n", r)
		}
	}()

	var d Handler
	db, err := d.initDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := fasthttprouter.New()
	router.POST("/", d.parse)
	log.Fatal(fasthttp.ListenAndServe(":4000", router.Handler))
}
