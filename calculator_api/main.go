package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Route1)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
