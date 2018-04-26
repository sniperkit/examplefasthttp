# examplefasthttp

A sample of the fast http server that accepts the json, and returns the result from the database in the format json.

Run

    go run main.go db.go

    or if you have Linux

    ./bin/linux/app

Request

    curl -H "Content-Type: application/json" -X POST -d '{"id":1}' http://*:4000
    curl -H "Content-Type: application/json" -X POST -d '{"id":2}' http://*:4000
    curl -H "Content-Type: application/json" -X POST -d '{"id":3}' http://*:4000
