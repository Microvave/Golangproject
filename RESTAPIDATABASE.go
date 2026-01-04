package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Student struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Score float64
}

func (Student) TableName() string {
	return "student"
}

func connectDB() {
	var err error

	dsn := "root:root@tcp(127.0.0.1:3306)/student_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(" connect database failed")
	}

	fmt.Println(" connected to database")
}

func addStudent(c *gin.Context) {
	var student Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.Create(&student).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, student)
}

func updateStudent(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var student Student
	if err := db.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}
	var updateStudent Student
	if err := c.ShouldBindJSON(&updateStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&student).Updates(updateStudent).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updateStudent)
}
func deleteStudent(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var student Student
	if err := db.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}

	// 3. ลบข้อมูลออกจาก slice
	if err := db.Delete(&student).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "student deleted",
	})
}

func main() {
	connectDB()
	db.AutoMigrate(&Student{})

	r := gin.Default()

	r.GET("/students", func(c *gin.Context) {
		var students []Student

		if err := db.Find(&students).Error; err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, students)
	})
	r.POST("/students", addStudent)
	r.PUT("/students/:id", updateStudent)
	r.DELETE("/students/:id", deleteStudent)
	r.Run(":8080")
}
