package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, App2")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":6666", nil))
}
