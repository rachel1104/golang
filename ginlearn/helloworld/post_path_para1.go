package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type User5 struct {
	Id         int64             `json:"id" uri:"id"`
	Name       string            `json:"name" uri:"name"`
	Address    []string          `json:"address" binding:"required"`
	AddressMap map[string]string `json:"addressMap"`
}

func main() {
	r := gin.Default()
	//curl http://localhost:8080/user/save/332/uwe
	//返回json
	r.POST("/user/save/:id/:name", func(ctx *gin.Context) {
		var user User5
		ctx.ShouldBindUri(&user)
		//addressMap := ctx.PostFormMap("addressMap")
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
