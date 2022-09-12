package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Model fo courses-goes to another file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"wesite"`
}

//fake DB

var courses []Course

//middleware - goes to another file

func (c *Course) isEmpty() bool {
	return c.CourseName == ""

}

func main() {
	fmt.Println("API")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Anila", Website: "go.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Angular", CoursePrice: 199, Author: &Author{Fullname: "Hitesh", Website: "lco.com"}})
	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controller - file

//serve home route - just showing something on localhost

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to our website"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Couse")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request - params is key value pair
	params := mux.Vars(r)
	fmt.Printf("type of params is %T", params)

	//loop through the courses ,find matching id and return the response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given course id")
	return

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("content-Type", "application/json")

	//what if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return

	}

	//generate unique id and convert that id into string because in struct we mentioned id as string
	//append our new course into existing courses (fake DB)
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa((rand.Intn(100)))

	//adding out new value into fake DB
	courses = append(courses, course)
	//show all value
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update course")
	w.Header().Set("Content-Type", "application/json")

	//first - grab id from request

	params := mux.Vars(r)

	//loop through the value once we hit the id then remove that course from fake DB
	//add with my ID which means add my data into the fake DB (using the same removed id)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			//decode json value and add data
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			//update operation so id shoud be same as removed one
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return

		}
	}
	//what if : body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send updated data")
	}
	//what if : {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty() {
		json.NewEncoder(w).Encode("send updated data inside json")
		return
	}

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop ,id, remove

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("course is deleted")
			break

		}

	}

}
