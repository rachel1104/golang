package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	//curl http://localhost:8080/user/save?addressMap[home]=Beijing&addressMap[company]=公司

	r.GET("/user/save", func(ctx *gin.Context) {
		addressMap := ctx.QueryMap("addressMap")
		ctx.JSON(200, addressMap)
	})
	//	address := ctx.QueryArray("address")
	//	ctx.JSON(200, address)
	//})
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
