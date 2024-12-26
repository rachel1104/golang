package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 返回字符串
func main() {
	r := gin.Default()
	r.GET("/user/save", func(ctx *gin.Context) {
		//ctx.Header("content-type", "text/html; charset=utf-8")
		ctx.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}

}
