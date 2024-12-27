package routers

import (
	"fmt"
	"fourthProject/controllers/itying"
	"fourthProject/middlewares"
	"github.com/gin-gonic/gin"
	"time"
)

// 中间件
func initMiddleware1(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("1-中间件")
	//调用该请求的剩余处理程序
	c.Next()
	fmt.Println("2-中间件")
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}
func initMiddleware2(c *gin.Context) {
	fmt.Println("3-中间件")
	//调用该请求的剩余处理程序
	c.Next()
	fmt.Println("4-中间件")
}
func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/", middlewares.InitMiddleware)
	//defaultRouters := r.Group("/",initMiddleware1)  //配置中间件
	{
		defaultRouters.GET("/", initMiddleware1, initMiddleware2, itying.DefaultController{}.Index)
		defaultRouters.GET("/news", initMiddleware1, initMiddleware2, itying.DefaultController{}.News)
		defaultRouters.GET("/shop", initMiddleware1, initMiddleware2, itying.DefaultController{}.Shop)
		defaultRouters.GET("/deletecookie", initMiddleware1, initMiddleware2, itying.DefaultController{}.DeleteCookie)

	}
}
