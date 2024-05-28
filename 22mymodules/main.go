package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Modules")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Good afternoon everyone!")
}

func serveHome(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Creating routes from golang for the first time"))
}