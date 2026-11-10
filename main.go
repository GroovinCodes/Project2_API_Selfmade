package main

// Here, we'll create a student grade sorting system
// This will use the gin API to conduct retrieval & edit operations.
// For editing, you would change the grades of the student within a slice

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// struct to contain student values and to be serialized
type Student struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
	Id    string `json:"id"`
}

var students = []Student{
	{Name: "Jason", Grade: 73, Id: "1"},
	{Name: "Terry", Grade: 80, Id: "2"},
	{Name: "Trevor", Grade: 92, Id: "3"},
}

// func to return json of all grades
func getStudents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, students)
}

// Now functions to edit grade and adding new students
// create body.json file and conduct edit operation
func addStudent(c *gin.Context) {
	var newStudent Student
	if err := c.BindJSON(&newStudent); err != nil {
		return //bindjson sends an error back
	}

	students = append(students, newStudent)
	c.IndentedJSON(http.StatusCreated, newStudent)
} // create a post in main

// Helper func to retrieve students by id
func StudentById(id string) (*Student, error) {
	for i, s := range students {
		if s.Id == id {
			return &students[i], nil
		}
	}
	return nil, errors.New("Student not found.")
}

// func to retrieve students via id
func getStudentById(c *gin.Context) {
	id := c.Param("id")
	student, err := StudentById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Student not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, student)
}

// To test, use the command: curl localhost:8080/students/<idnum>
// Now, the ultimate test - editing the grades via a query parameter
func editGrade(c *gin.Context) {
	id := c.Param("id")
	newGrade := c.Param("grade")

	gradeVal := 0
	fmt.Sscanf(newGrade, "%d", &gradeVal) // converts into a string into int

	student, err := StudentById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Student not found"})
		return
	}

	student.Grade = gradeVal
	c.IndentedJSON(http.StatusOK, students)
}

// Use the command curl -X PUT localhost:8080/students/<idnum>/<gradevalue> to update the student grades

// Now time to implement changes to the function
func main() {
	router := gin.Default()
	router.GET("/students", getStudents)
	router.GET("/students/:id", getStudentById)
	router.POST("/students", addStudent)
	router.PUT("/students/:id/:grade", editGrade) // changes the grade value via a query parameter
	router.Run("localhost:8080")
}

// Done! You've completed this project.