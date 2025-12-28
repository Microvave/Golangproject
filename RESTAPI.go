package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Name  string
	Score float64
}

var students []Student

func addStudent(c *gin.Context) {
	var student Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	students = append(students, student)

	c.JSON(201, student)
}

func updateStudent(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if id < 0 || id >= len(students) {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}
	var updateStudent Student
	if err := c.ShouldBindJSON(&updateStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	students[id] = updateStudent
	c.JSON(http.StatusOK, updateStudent)
}
func deleteStudent(c *gin.Context) {
	// 1. ดึง id จาก URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// 2. ตรวจสอบ index
	if id < 0 || id >= len(students) {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}

	// 3. ลบข้อมูลออกจาก slice
	deletedStudent := students[id]
	students = append(students[:id], students[id+1:]...)
	c.JSON(http.StatusOK, gin.H{
		"message": "student deleted",
		"data":    deletedStudent,
	})
}

func main() {

	r := gin.Default()

	r.GET("/students", func(c *gin.Context) {
		c.JSON(200, students)
	})
	r.POST("/students", addStudent)

	r.PUT("/students/:id", updateStudent)
	r.DELETE("/students/:id", deleteStudent)
	r.Run(":8080")
}
