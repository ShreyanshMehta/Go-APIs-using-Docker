package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Hello, Shreyansh Mehta")
	greeter()
	courses = append(courses, Course{CourseId: "33224", CourseName: "GO Lang", CoursePrice: 332, Author: &Author{Fullname: "Shreyansh Mehta", Website: "www.shreyanshmehta.com"}})
	courses = append(courses, Course{CourseId: "34224", CourseName: "Python", CoursePrice: 1000, Author: &Author{Fullname: "Shreyansh Mehta", Website: "www.shreyanshmehta.com"}})

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/insert", insertOneCourse).Methods("POST")
	port := os.Getenv("PORT")
	port = ":" + port
	fmt.Println(port, os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(port, r))
}

func greeter() {
	fmt.Println("hey, this is a greeter")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the GoLang World</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func insertOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
	}
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}
