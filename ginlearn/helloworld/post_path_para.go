package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	//curl http://localhost:8080/user/save/332/uwe

	r.POST("/user/save/:id/:name", func(ctx *gin.Context) {
		id := ctx.Param("id")
		name := ctx.Param("name")
		//addressMap := ctx.PostFormMap("addressMap")
		ctx.JSON(200, gin.H{
			"id":   id,
			"name": name,
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
