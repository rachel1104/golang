package admin

import (
	"fourthProject/models"
	"github.com/gin-gonic/gin"
)

type NavController struct {
	BaseController
}

func (con NavController) Index(c *gin.Context) {
	//查詢全部數據
	//navList := []models.Nav{}
	//models.DB.Find(&navList)
	//c.JSON(200, gin.H{
	//	"code":   200,
	//	"result": navList,
	//})
	//查詢一條數據
	navResult := models.Nav{Id: 1}
	models.DB.Find(&navResult)

	c.JSON(200, gin.H{
		"code":   200,
		"result": navResult,
	})
	////查詢id>3的數據
	//navList := []models.Nav{}
	//models.DB.Where("id>3").Find(&navList)
	//查詢id>3 and id<9的數據
	//navList := []models.Nav{}
	//var a := 3
	//var b := 9
	//models.DB.Where("id>? and id<>",a,b).Find(&navList)
	////使用in查詢id在3,5,6中的數據
	//navList := []models.Nav{}
	//models.DB.Where("id in (?)", []int{3, 5, 6}).Find(&navList)
	//使用like查詢
	//navList := []models.Nav{}
	//models.DB.Where("title like ?", "%回%").Find(&navList)
	////使用between
	//navList := []models.Nav{}
	//models.DB.Where("id between ? and ?", 3, 9).Find(&navList)
	////使用or
	//navList := []models.Nav{}
	//models.DB.Where("id=? or id=?", 3, 9).Find(&navList)
	//models.DB.Where("id=?", 3) or("id=?", 9).Find(&navList)
	////使用select 指定返回的字段
	//navList := []models.Nav{}
	//models.DB.Select("id,title").Find(&navList)
	//Order排序，Limit，Offset
	//navList := []models.Nav{}
	//models.DB.Order("id desc").Order("title desc").Limit(2).Find(&navList)
	//分頁
	//navList := []models.Nav{}
	//models.DB.Order("id desc").Offset(2).Limit(2).Find(&navList)
	//count統計數量
	//navList := []models.Nav{}
	//var num int64
	//models.DB.Find(&navList).Count(&num)
	//c.JSON(200, gin.H{
	//	"code":   200,
	//	"result": num,
	//})
	//使用原生sql刪除，修改
	//models.DB.Exec("DELETE FROM nav WHERE id=?", 1)
	//models.DB.Exec("update nav set title='asdfasf' where id=?", 1)
	//使用sql查詢uid=2數據
	//uerList := []models.User{}
	//models.DB.Raw("select * from user where id>2	").Scan(&uerList)
	//c.JSON(200, gin.H{
	//	"code":   200,
	//	"result": uerList,
	//})
}
