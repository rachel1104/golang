package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

type Article3 struct {
	Title   string `xml:"title"`
	Desc    string `xml:"desc"`
	Content string `xml:"content"`
}
type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Age      int    `json:"age" form:"age"`
}

// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func Text(str1, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + str2
}
func main() {
	//创建一个默认的路由引擎
	r := gin.Default()
	//自定义模板函数，注意要把这个函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		//注册模板函数
		"UnixToTime": UnixToTime,
		"Text":       Text,
	})
	//加载模板 放在配置路由器前面
	r.LoadHTMLGlob("templates/**/*")
	//配置静态web目录第一个参数表示路由，第二个参数表示映射的目录
	r.Static("/static", "./static")
	//GET 请求不传参数 http://localhost:8080/no    结果：{"username":"rachel"}
	r.GET("/no", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"username": "rachel",
		})
	})
	//GET 请求传值  http://localhost:8080/?username=gege&age=18&page=3  结果：{"age":"18","page":"3","username":"gege"}
	r.GET("/", func(c *gin.Context) {
		username := c.Query("username")
		age := c.Query("age")
		page := c.DefaultQuery("page", "1")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})
	//GET 请求传值id  http://localhost:8080/article?id=3  结果：{"id":"3","msg":"新闻详情页面"}
	r.GET("/article", func(c *gin.Context) {
		id := c.DefaultQuery("id", "1")
		c.JSON(http.StatusOK, gin.H{
			"id":  id,
			"msg": "新闻详情页面",
		})
	})
	//get post
	r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/user.html", gin.H{})
	})
	//获取表单传过来的数据
	r.POST("/doAddUser1", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "20")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})
	//获取get,post的数据绑定到结构体  http://localhost:8080/getUser?username=222&password=123455&age=3    结果：{"username":"222","password":"123455","age":3}
	r.GET("/getUser", func(c *gin.Context) {
		user := &UserInfo{}

		c.ShouldBind(&user)
		if err := c.ShouldBind(&user); err == nil {
			fmt.Println(user)
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
	})
	r.POST("/doAddUser2", func(c *gin.Context) {
		user := &UserInfo{}
		if err := c.ShouldBind(&user); err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
	})
	//获取post xml数据
	/*
		<?xml version="1.0" encoding="UTF-8"?>
		<article>
		  <content type="string">我是张三</content>
		  <title type="string">张三</title>
		</article>
	*/
	r.POST("/xml", func(c *gin.Context) {
		article := &Article3{}
		xmlSliceData, _ := c.GetRawData() //获取 c.Requests.Body 读取请求数据
		if err := xml.Unmarshal(xmlSliceData, article); err == nil {
			c.JSON(http.StatusOK, article)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
	})

	//动态路由传值  http://localhost:8080/list/123
	r.GET("/list/:cid", func(c *gin.Context) {
		cid := c.Param("cid")
		c.String(http.StatusOK, "%v", cid)
	})
	//启动web服务
	r.Run(":8080")
}
