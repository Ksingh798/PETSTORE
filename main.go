package main

import (
	"log"
	"net/http"
	"strconv"

	"petstore/internal"
	"petstore/pkg"
)

func main() {
	r := pkg.NewRouter()
	log.Printf("Listening at %d port...", internal.DefaultServicePort)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(internal.DefaultServicePort), r))
}
