package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

/*
{
"Id": 123,
"Name": "user",
"Address":[
"wuhan",
"shanghai"
],
"AddressMap":{
"home": "wuh"
}
}
*/
type User4 struct {
	Id         int64             `json:"id"`
	Name       string            `json:"name"`
	Address    []string          `json:"address" binding:"required"`
	AddressMap map[string]string `json:"addressMap"`
}

func main() {
	r := gin.Default()
	//curl http://localhost:8080/user/save

	r.POST("/user/save", func(ctx *gin.Context) {
		var user User4
		//post参数是form格式
		ctx.ShouldBindJSON(&user)
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
