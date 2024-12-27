package admin

import (
	"fourthProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleController struct {
}

func (con ArticleController) Index(c *gin.Context) {
	////查詢article表數據
	//articleList := []models.Article{}
	//models.DB.Find(&articleList)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": articleList,
	//})

	//獲取文章對應的分類(文章對分類 1對1)
	//articleList := []models.Article{}
	//models.DB.Preload("ArticleCate").Find(&articleList)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": articleList,
	//})

	////獲取所有的分類的文章
	articleCateList := []models.ArticleCate{}
	//models.DB.Preload("Article").Find(&articleCateList)
	models.DB.Preload("Article").Find(&articleCateList)
	c.JSON(http.StatusOK, gin.H{
		"result": articleCateList,
	})
}
func (con ArticleController) Add(c *gin.Context) {
	c.String(200, "文章新增")
}
func (con ArticleController) Edit(c *gin.Context) {
	c.String(200, "编辑")
}
