package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	})
	r.Static("/css", "./static/css")
	//加载模板
	//r.LoadHTMLFiles("./templates/index.tmpl", "./templates/user.tmpl")
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "<a href='https://www.baidu.com'>hello world</a>",
		})
	})
	r.GET("/user", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "user.tmpl", gin.H{
			"title": "hello user",
		})
	})
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
