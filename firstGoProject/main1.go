package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article1 struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	//配置路由
	//可以返回字符串，json，结构体，jsonp，xml，
	//返回字符串
	r.GET("/", func(c *gin.Context) {
		c.String(200, "值：%v", "Hello World")
	})
	//返回json
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"sucess": true,
			"msg":    "你好gin",
		})
	})
	//返回结构体
	r.GET("/struct", func(c *gin.Context) {
		a := &Article{
			Title:   "标题1",
			Desc:    "描述",
			Content: "测试内容",
		}
		c.JSON(200, a)
	})
	//返回JSONP 主要用于跨域请求 ,http://localhost:8080/jsonp?callback=aaa  参数放入回调函数中
	//结果：aaa({"title":"标题1","desc":"描述","content":"测试内容"});
	r.GET("/jsonp", func(c *gin.Context) {
		a := &Article1{
			Title:   "标题1",
			Desc:    "描述",
			Content: "测试内容",
		}
		c.JSONP(200, a)
	})
	//返回一个xml
	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"success": true,
			"msg":     "你好",
		})
	})
	//返回模板
	r.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是后台的商品数据",
		})
	})
	r.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"title": "我是后台的数据",
		})
	})
	r.POST("/add", func(c *gin.Context) {
		c.String(200, "我是POST请求返回的数据")
	})
	r.PUT("/edit", func(c *gin.Context) {
		c.String(200, "我是put请求返回的数据")
	})
	r.DELETE("/delete", func(c *gin.Context) {
		c.String(http.StatusOK, "我是delete请求返回的数据")
	})
	//启动web服务
	r.Run(":8080")
}
