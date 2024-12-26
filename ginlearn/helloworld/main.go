package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type User struct {
	Id   int64  `form:"id"`
	Name string `form:"name"`
}

func main() {
	r := gin.Default()
	//curl http://localhost:8080/hello  get获取json返回值 {“name”:"hello world"}
	r.GET("/user/save", func(ctx *gin.Context) {
		var user User
		err := ctx.Bind(&user)
		if err != nil {
			log.Println(err)
		}
		//address, ok := ctx.GetQuery("address")
		//address := ctx.DefaultQuery("address", "wuhan")
		ctx.JSON(200, user)
	})
	//r.GET("/hello", func(ctx *gin.Context) {
	//	//返回数组，map，list，结构体
	//	ctx.JSON(200, gin.H{
	//		"name": "hello world!",
	//	})
	//})
	//r.Any("/user", func(ctx *gin.Context) {
	//	//返回数组，map，list，结构体
	//	ctx.JSON(200, "any")
	//})
	//r.GET("/user/*path", func(ctx *gin.Context) {
	//	//返回数组，map，list，结构体
	//	ctx.JSON(200, ctx.Param("path"))
	//})
	//v1 := r.Group("/v1")
	//{
	//	v1.GET("find", func(ctx *gin.Context) {
	//		ctx.JSON(200, "v1 find")
	//	})
	//	v1.GET("save", func(ctx *gin.Context) {
	//		ctx.JSON(200, "v1 save")
	//	})
	//}
	//v2 := r.Group("/v2")
	//{
	//	v2.GET("find", func(ctx *gin.Context) {
	//		ctx.JSON(200, "v2 find")
	//	})
	//	v2.GET("save", func(ctx *gin.Context) {
	//		ctx.JSON(200, "v2 save")
	//	})
	//}
	//r.POST("/user", func(ctx *gin.Context) {
	//	//返回数组，map，list，结构体
	//	ctx.JSON(200, "post")
	//})
	//r.PUT("/user", func(ctx *gin.Context) {
	//	//返回数组，map，list，结构体
	//	ctx.JSON(200, "put")
	//})
	//r.DELETE("/user", func(ctx *gin.Context) {
	//	//返回数组，map，list，结构体
	//	ctx.JSON(200, "delete")
	//})
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
