package admin

import (
	"fourthProject/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type StudentController struct {
	BaseController
}

func (con StudentController) Index(c *gin.Context) {
	//1、獲取學生信息
	//studentList := []models.Student{}
	//models.DB.Find(&studentList)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": studentList,
	//})

	//2.獲取課程信息
	//lessonList := []models.Lesson{}
	//models.DB.Find(&lessonList)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": lessonList,
	//})

	//3.獲取所有學生的課程信息 多對多
	//studentList := []models.Student{}
	//models.DB.Preload("Lesson").Find(&studentList)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": studentList,
	//})

	//4.查詢一個學生選修了哪些課程
	//studentList := []models.Student{}
	//models.DB.Preload("Lesson").Where("studentname=?", "sophia").Find(&studentList)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": studentList,
	//})

	//5.查詢課程被哪些學生選修了  在lesson.go中添加Student []Student `gorm:"many2many:lesson_students;"`
	//lessonList := []models.Lesson{}
	//models.DB.Preload("Student").Find(&lessonList)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": lessonList,
	//})

	////6.查詢數學課被哪些學生選修
	//lessonList := []models.Lesson{}
	//models.DB.Preload("Student").Where("id=?", 1).Find(&lessonList)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": lessonList,
	//})

	////7.查詢數據指定條件
	//lessonList := []models.Lesson{}
	//models.DB.Preload("Student").Offset(1).Order("id desc").Limit(2).Find(&lessonList)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": lessonList,
	//})

	//8.查詢課程被哪些同學選項，除去某些人
	//lessonList := []models.Lesson{}
	////models.DB.Preload("Student", "id!=2").Find(&lessonList)
	//models.DB.Preload("Student", "id not in (1,2)").Find(&lessonList)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": lessonList,
	//})

	//9.查詢課程被哪些同學選項，按id倒序排列 自定義預加載
	lessonList := []models.Lesson{}
	//models.DB.Preload("Student", "id!=2").Find(&lessonList)
	models.DB.Preload("Student", func(db *gorm.DB) *gorm.DB {
		return models.DB.Order("student.id desc")
	}).Find(&lessonList)
	c.JSON(http.StatusOK, gin.H{
		"result": lessonList,
	})
}
