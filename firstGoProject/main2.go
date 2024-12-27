package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

type Article2 struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// 时间戳转换成日期
func UnixToTime2(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func Text2(str1, str2 string) string {
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

	//配置路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "首页",
			"msg":   "成功",
			"score": 89,
			"hobby": []string{"吃饭", "上课", "学习", "睡觉"},
			"newsList": []interface{}{
				&Article2{
					Title:   "新闻标题1",
					Desc:    "描述1",
					Content: "新闻详情1",
				},
				&Article2{
					Title:   "新闻标题2",
					Desc:    "描述2",
					Content: "新闻详情2",
				},
			},
			"testslice": []string{},
			"news": &Article2{
				Title:   "新闻标题3",
				Desc:    "描述3",
				Content: "新闻详情3",
			},
			"date": 1629423555,
		})
	})
	r.GET("/news", func(c *gin.Context) {
		news := &Article{
			Title:   "新闻标题",
			Desc:    "描述",
			Content: "新闻详情",
		}
		c.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "新闻首页",
			"news":  news,
		})
	})

	//启动web服务
	r.Run(":8080")
}
