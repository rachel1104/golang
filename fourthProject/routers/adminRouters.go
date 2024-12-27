package routers

import (
	"fourthProject/controllers/admin"
	"fourthProject/middlewares"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	//配置中间件 添加middlewares.InitMiddleware
	adminRouters := r.Group("/admin", middlewares.InitMiddleware)
	{
		//配置路由
		adminRouters.GET("/", admin.IndexController{}.Index)

		adminRouters.GET("/user", admin.UserController{}.Index)
		adminRouters.GET("/user/add", admin.UserController{}.Add)
		//adminRouters.POST("/user/doUpload", admin.UserController{}.DoUpload)
		adminRouters.GET("/user/edit", admin.UserController{}.Edit)
		adminRouters.GET("/user/delete", admin.UserController{}.Delete)
		//adminRouters.POST("/user/doEdit", admin.UserController{}.DoEdit)
		//adminRouters.GET("/user/add1", admin.UserController{}.Add1)
		//adminRouters.POST("/user/doUpload1", admin.UserController{}.DoUpload1)
		adminRouters.GET("/nav", admin.NavController{}.Index)
		adminRouters.GET("/student", admin.StudentController{}.Index)
		adminRouters.GET("/bank", admin.BankController{}.Index)

		adminRouters.GET("/article", admin.ArticleController{}.Index)
		adminRouters.GET("/article/add", admin.ArticleController{}.Add)
		adminRouters.GET("/article/edit", admin.ArticleController{}.Edit)
	}
}
