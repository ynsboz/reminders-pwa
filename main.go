package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.FileServer(http.Dir(""))
	http.Handle("/", handler)
	log.Fatal(http.ListenAndServe("", handler))
}
