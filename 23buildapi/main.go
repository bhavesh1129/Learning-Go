package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for author
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Model for course
type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice float32 `json:"price"`
	Author      *Author `json:"author"`
}

// Fake db
var courses []Course

// Middleware/helper
func (c *Course) isEmpty() bool {
	return c.CourseName == "" && c.CoursePrice == 0.0 && c.Author == nil
}

func main() {
	fmt.Println("Building APIs")
	r := mux.NewRouter()

	//seeding mock data to fake db
	courses = append(courses, Course{
		CourseID: "1", 
		CourseName: "Get Ready for Golang", 
		CoursePrice: 29.99, 
		Author: &Author{
			Fullname: "Bhavesh Garg", 
			Website: "www.go.dev",
		},
	})

	courses = append(courses, Course{
		CourseID:    "2",
		CourseName:  "Introduction to Python Programming",
		CoursePrice: 39.99,
		Author: &Author{
			Fullname: "Alice Smith",
			Website:  "www.python101.com",
		},
	})
	
	courses = append(courses, Course{
		CourseID:    "3",
		CourseName:  "Web Development with JavaScript",
		CoursePrice: 49.99,
		Author: &Author{
			Fullname: "Bob Johnson",
			Website:  "www.jsmasterclass.com",
		},
	})
	
	courses = append(courses, Course{
		CourseID:    "4",
		CourseName:  "Mastering React Framework",
		CoursePrice: 59.99,
		Author: &Author{
			Fullname: "Charlie Brown",
			Website:  "www.reactpros.com",
		},
	})
	
	//Routing
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourseById).Methods("GET")
	r.HandleFunc("/add-course", createACourse).Methods("POST")
	r.HandleFunc("/update-course/{id}", updateACourse).Methods("PUT")
	r.HandleFunc("/delete-course/{id}", deleteACourse).Methods("DELETE")
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers

// homepage
func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to our homepage!"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, course := range courses {
		if course.CourseID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("No course found with the given id")
}

func createACourse(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Please provide some data")
		return
	}

	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid JSON format")
		return
	}
	if course.isEmpty() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Please provide some data in your JSON object")
		return
	}

	for _, c := range courses {
		if course.CourseName == c.CourseName{
			json.NewEncoder(w).Encode("A course with the same name already exists")
            return
		}
	}

	// Generate a unique string for CourseID
	rand.Uint32()
	course.CourseID = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(course)
}

func updateACourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("The object which you wants to update don't exists!")
	return
}

func deleteACourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("The course has been successfully deleted now!")
			return
		}
	}
}
