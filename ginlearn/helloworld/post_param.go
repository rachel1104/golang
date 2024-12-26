package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	//curl http://localhost:8080/user/save

	r.POST("/user/save", func(ctx *gin.Context) {
		//post参数是form格式
		id := ctx.PostForm("id")
		name := ctx.PostForm("name")
		address := ctx.PostFormArray("address")
		addressMap := ctx.PostFormMap("addressMap")
		ctx.JSON(200, gin.H{
			"id":         id,
			"name":       name,
			"address":    address,
			"addressMap": addressMap,
		})
	})
	//	address := ctx.QueryArray("address")
	//	ctx.JSON(200, address)
	//})
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
