package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type User3 struct {
	Id         int64             `form:"id"`
	Name       string            `form:"name"`
	Address    []string          `form:"address" binding:"required"`
	AddressMap map[string]string `form:"addressMap"`
}

func main() {
	r := gin.Default()
	//curl http://localhost:8080/user/save

	r.POST("/user/save", func(ctx *gin.Context) {
		var user User3
		//post参数是form格式
		ctx.ShouldBind(&user)
		addressMap := ctx.PostFormMap("addressMap")
		user.AddressMap = addressMap
		ctx.JSON(200, user)
	})
	//	address := ctx.QueryArray("address")
	//	ctx.JSON(200, address)
	//})
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
