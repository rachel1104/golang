package api

import "github.com/gin-gonic/gin"

type ApiController struct{}

func (con ApiController) Index(c *gin.Context) {
	c.String(200, "API接口")
}

func (con ApiController) Userlist(c *gin.Context) {
	c.String(200, "API接口--userlist")
}

func (con ApiController) Plist(c *gin.Context) {
	c.String(200, "API接口--plist")
}
