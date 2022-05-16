package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Modules in golang!")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	//If any error occurs use log module, which will print the errors directly or simply use comma ok syntax
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hello mod users!")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang!</h1>"))
}
