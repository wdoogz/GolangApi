package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Yo!")
	fmt.Println("Endpoint: homePage")
}

func handler() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func main() {
	handler()
}
