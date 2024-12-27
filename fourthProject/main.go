package main

import (
	"fmt"
	"fourthProject/models"
	"fourthProject/routers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
)

type Article3 struct {
	Title   string `xml:"title"`
	Desc    string `xml:"desc"`
	Content string `xml:"content"`
}
type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Age      int    `json:"age" form:"age"`
}

// 全局中间件
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

//// 时间戳转换成日期  将此方法移至models
//func UnixToTime(timestamp int) string {
//	fmt.Println(timestamp)
//	t := time.Unix(int64(timestamp), 0)
//	return t.Format("2006-01-02 15:04:05")
//}

func Text(str1, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + str2
}
func main() {
	//创建一个默认的路由引擎
	r := gin.Default()
	//自定义模板函数，注意要把这个函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		//注册模板函数
		"UnixToTime": models.UnixToTime,
		"Text":       Text,
	})
	//加载模板 放在配置路由器前面
	r.LoadHTMLGlob("templates/**/*")
	//配置静态web目录第一个参数表示路由，第二个参数表示映射的目录
	r.Static("/static", "./static")

	//配置session中间件
	////创建基于cookie的存储引擎，secret1111参数是用于加密的密钥
	//store := cookie.NewStore([]byte("secret1111"))
	////配置session中间件，sotre是前面创建的存储引擎
	//r.Use(sessions.Sessions("mysession", store))

	//
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	//加载全局中间件
	r.Use(initMiddleware1, initMiddleware2)

	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.DefaultRoutersInit(r)
	//启动web服务
	r.Run(":80")
}
