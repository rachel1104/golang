package admin

import (
	"fmt"
	"fourthProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strconv"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	//c.String(200, "用户列表")
	con.success(c)
}
func (con UserController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/useradd.html", gin.H{})
}

/*
按照日期生成目录图像
1.获取上传的文件
2.获取后缀名，判断类型是否正确 .jpg,.png,gif,.jpeg
3.创建图片保存目录 static/upload/20241226
4.生成文件名和文件保存的目录
5.执行上传
*/
func (con UserController) DoUpload(c *gin.Context) {
	username := c.PostForm("username")
	//1.获取上传的文件
	file, err := c.FormFile("faceimage")

	if err == nil {
		//2.获取后缀名，判断类型是否正确 .jpg,.png,gif,.jpeg
		extName := path.Ext(file.Filename)
		allowExtMap := map[string]bool{
			".jpg":  true,
			".jpeg": true,
			".png":  true,
			".gif":  true,
		}
		if _, ok := allowExtMap[extName]; !ok {
			c.String(200, "上传的文件类型不合法")
			return
		}
		//3.创建图片保存目录 static/upload/20241226
		day := models.GetDay()
		dir := "./static/upload" + day
		err := os.MkdirAll(dir, 0666)
		if err != nil {
			fmt.Println(err)
			c.String(200, err.Error())
			return
		}
		//4.生成文件名和文件保存的目录
		fileName := strconv.FormatInt(models.GetUnix(), 10) + extName
		//5.执行上传
		dst := path.Join(dir, fileName)
		c.SaveUploadedFile(file, dst)
	}
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
	}) //{"dst":"static/upload/11.png","success":true,"username":"3222"}
}

func (con UserController) Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/useredit.html", gin.H{})
}
func (con UserController) DoEdit(c *gin.Context) {
	//c.String(200, "执行修改")
	username := c.PostForm("username")
	file1, err1 := c.FormFile("faceimage1")
	dst1 := path.Join("./static/upload", file1.Filename)
	if err1 == nil {
		c.SaveUploadedFile(file1, dst1)
	}
	file2, err2 := c.FormFile("faceimage2")
	dst2 := path.Join("./static/upload", file2.Filename)
	if err2 == nil {
		c.SaveUploadedFile(file2, dst2)
	}
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
		"dst1":     dst1,
		"dst2":     dst2,
	}) //{"dst1":"static/upload/11.png","dst2":"static/upload/72dba1a02d836c983fd60930cb49df11.jpg","success":true,"username":"edit"}

}

func (con UserController) Add1(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/useradd1.html", gin.H{})
}

func (con UserController) DoUpload1(c *gin.Context) {
	username := c.PostForm("username")
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["faceimage[]"]

	for _, file := range files {
		// 上传文件至指定目录
		dst := path.Join("./static/upload", file.Filename)
		c.SaveUploadedFile(file, dst)
	}
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
	})
}
