# examplefasthttp

A sample of the fasthttp server with catching interrupt, that accepts the json, and returns the result from the 
database in the format 
json.

Run

    ./bin/linux/app
    
Docker

    docker build -t test .
    docker run --publish 4000:4000 --name test --rm test

Request

    curl -H "Content-Type: application/json" -X POST -d '{"id":1}' http://*:4000
    curl -H "Content-Type: application/json" -X POST -d '{"id":2}' http://*:4000
    curl -H "Content-Type: application/json" -X POST -d '{"id":3}' http://*:4000

Build

    make check - check errors
    make vendor - update vendor packages
    make build - build app

Size

    goupx ./bin/linux/app