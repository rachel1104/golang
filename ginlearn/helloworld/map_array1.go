package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type User2 struct {
	Id         int64             `form:"id"`
	Name       string            `form:"name"`
	Address    []string          `form:"address" binding:"required"`
	AddressMap map[string]string `form:"addressMap"`
}

func main() {
	r := gin.Default()
	//curl http://localhost:8080/user/save?id=22&name=user&addressMap[home]=Beijing&addressMap[company]=公司

	r.GET("/user/save", func(ctx *gin.Context) {
		var user2 User2
		ctx.ShouldBindQuery(&user2)
		user2.AddressMap = ctx.QueryMap("addressMap")
		ctx.JSON(200, user2)
	})
	//	address := ctx.QueryArray("address")
	//	ctx.JSON(200, address)
	//})
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
