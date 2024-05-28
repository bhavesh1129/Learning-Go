package controller

import (
	"encoding/json"
	"net/http"

	"github.com/bhavesh1129/mongodbapi/helpers"
	"github.com/bhavesh1129/mongodbapi/models"
	"github.com/gorilla/mux"
)

// insert one record
func CreateAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	helpers.InsertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

// update one record
func MarkedAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	helpers.UpdateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// delete one record
func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	helpers.DeleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// delete all record
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	helpers.DeleteAllMovies()
	json.NewEncoder(w).Encode("All movies has been deleted succesfully!")
}

// get all movies
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allMovies := helpers.GetAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}
