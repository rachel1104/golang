package admin

import (
	"fmt"
	"fourthProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	//查询数据库
	userList := []models.User{}
	//查詢所有用戶
	//models.DB.Find(&userList)
	//查詢age大于8的用戶
	models.DB.Where("age>8").Find(&userList)
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": userList,
	})
}
func (con UserController) Add(c *gin.Context) {
	//c.String(http.StatusOK, "增加用户")
	user := models.User{
		Username: "andy",
		Pwd:      "123456",
		Age:      3,
		Email:    "andy@gmail.com",
		AddTime:  1000001111,
	}
	models.DB.Create(&user)
	fmt.Println(user)
	c.String(http.StatusOK, "增加用户成功")
}

func (con UserController) Edit(c *gin.Context) {
	////查詢id等于1的數據,保存所有字段
	//方法1
	//user := models.User{Id: 1}
	//models.DB.First(&user)
	////更新數據
	//user.Username = "曼曼"
	//user.Pwd = "123456"
	//user.Email = "sophia@gmail.com"
	//user.AddTime = int(models.GetUnix())
	//models.DB.Save(&user)
	//c.String(http.StatusOK, "修改用户成功")
	//方法2
	//user := models.User{}
	//models.DB.Model(&user).Where("id=?", 1).Update("username", "曼曼")
	//方法3
	user := models.User{}
	models.DB.Where("id=1").First(&user)
	user.Username = "luo曼菲"
	user.Pwd = "123456"
	user.Age = 10
	user.Email = "sophia@gmail.com"
	user.AddTime = int(models.GetUnix())
	models.DB.Save(&user)
	c.String(http.StatusOK, "修改用户成功")
}
func (con UserController) Delete(c *gin.Context) {
	//user := models.User{Id: 2}
	//models.DB.Delete(&user)
	user := models.User{}
	models.DB.Where("username=?", "andy").Delete(&user)
	c.String(http.StatusOK, "删除用户成功")
}
