package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
)

type Form struct {
	value map[string][]string
	File  map[string][]*multipart.FileHeader
}

func main() {
	r := gin.Default()
	r.POST("/user/save", func(ctx *gin.Context) {
		form, err := ctx.MultipartForm()
		if err != nil {
			log.Println(err)
		}
		value := form.Value
		files := form.File
		for _, fileArray := range files {
			for _, v := range fileArray {
				ctx.SaveUploadedFile(v, "./file/"+v.Filename)
			}
		}
		ctx.JSON(200, value)
	})
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
