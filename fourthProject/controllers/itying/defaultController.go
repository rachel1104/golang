package itying

import (
	"fmt"
	"fourthProject/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	//设置cookie
	//3600表示秒，时段内cooki有效
	//第六个参数secure，当为true时，在http中无效，在https中有效，所以一般设置为false
	//第七个参数httpOnly，是微软对cookie的扩展，若设置了属性，将无法读取cookie信息，防止xss攻击
	c.SetCookie("gin_cookie", "test", 3600, "/", ".itying.com", false, true)
	//过期时间延时
	c.SetCookie("username", "zhangsan", 5, "/", ".itying.com", false, true)

	//设置sessions
	session := sessions.Default(c)
	//配置session的过期时间
	session.Options(sessions.Options{
		MaxAge: 3600 * 6, //6小时
	})
	session.Set("username", "zhangsan11111")
	session.Save() //设置session的时候必须调用
	fmt.Println("请求--------")
	print(models.UnixToTime(1629788418))
	time.Sleep(time.Second)
	c.HTML(http.StatusOK, "default/index.html", gin.H{
		"title": "hello world",
		"t":     1629788418,
	})
}
func (con DefaultController) News(c *gin.Context) {
	//获取cookie
	name, _ := c.Cookie("gin_cookie")
	ho, _ := c.Cookie("username")

	//获取sessions
	session := sessions.Default(c)
	username := session.Get("username")
	c.String(200, "name=%v----ho=%v, username=%v", name, ho, username)
}
func (con DefaultController) Shop(c *gin.Context) {
	//获取cookie
	name, _ := c.Cookie("gin_cookie")
	ho, _ := c.Cookie("username")
	c.String(200, "name=%v----ho=%v", name, ho)
}

func (con DefaultController) DeleteCookie(c *gin.Context) {
	//删除cookie,一种方式是value设置为空，或者maxAge设置为负数-1
	c.SetCookie("gin_cookie", "test", -1, "/", "localhost", false, true)
	c.String(200, "删除成功")
}
