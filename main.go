package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var students = []Student{
	{ID: 1, Name: "Abdullah", Class: "1-b", Teacher: "Osman"},
	{ID: 2, Name: "Ahmet", Class: "2-b", Teacher: "Ömer"},
}

type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Class   string `json:"class"`
	Teacher string `json:"teacher"`
}

func main() {
	router := gin.Default()
	router.GET("/students", listStudents)
	router.POST("/students", createStudent)
	router.GET("/students/:id", getStudent)
	router.Run("localhost:9090")
}

func listStudents(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, students)
}

func createStudent(context *gin.Context) {
	var studentByUser Student
	err := context.BindJSON(&studentByUser)

	if err == nil && studentByUser.ID != 0 && studentByUser.Class != "" && studentByUser.Name != "" && studentByUser.Teacher != "" {
		students = append(students, studentByUser)
		context.IndentedJSON(http.StatusCreated, gin.H{"message": "Student has been created", "student_id": studentByUser.ID})
		return
	} else {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Student cannot be created"})
		return
	}
}

func getStudent(context *gin.Context) {
	str_id := context.Param("id")
	int_id, _ := strconv.Atoi(str_id)

	for _, a := range students {
		if a.ID == int_id {
			context.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})

}
