package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	//设置cookie
	r.GET("/cookie", func(ctx *gin.Context) {
		//设置cookie
		ctx.SetCookie("site_cookie", "cookievalue", 3600, "/", "localhost", false, true)
	})
	//读取cookie
	r.GET("/cookie1", func(ctx *gin.Context) {
		data, err := ctx.Cookie("site_cookie")
		if err != nil {
			log.Println(err)
			ctx.JSON(200, err.Error())
			return
		}
		ctx.JSON(200, data)
	})
	//删除cookie
	r.GET("/del", func(ctx *gin.Context) {
		//设置cookie，maxAge设置为-1，表示删除cookie
		ctx.SetCookie("site_cookie", "cookievalue", -1, "/", "localhost", false, true)
		ctx.String(200, "删除cookie")
	})
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}

}
