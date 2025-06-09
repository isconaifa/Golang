package main

import (
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}
func main() {

	http.HandleFunc("/home", homeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
