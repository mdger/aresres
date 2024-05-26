package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// TODO: structured logging

	http.HandleFunc("GET /", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "sadfkjsdf")
	})

	port := ":8485"
	log.Println("starting http server and listening on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
