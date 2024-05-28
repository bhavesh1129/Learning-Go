package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bhavesh1129/mongodbapi/database"
	"github.com/bhavesh1129/mongodbapi/router"
)

func main() {
	database.DBConnection()
	r := router.Router()
	fmt.Println("Server is getting started!")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Print("Listening is at port 4000 port")
}
