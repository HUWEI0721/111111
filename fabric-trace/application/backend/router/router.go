package router

// 路由文件
import (
	con "backend/controller"
	"backend/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	// 解决跨域问题
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // 允许的来源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // 允许的请求方法
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},                          // 暴露的响应头
		AllowCredentials: true,                                                // 允许传递凭据（例如 Cookie）
		MaxAge:           12 * time.Hour,                                      // 预检请求的有效期
	}))
	// 设置静态文件目录
	r.Static("/static", "./dist/static")
	r.LoadHTMLGlob("dist/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	// 测试GET请求
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//注册
	r.POST("/register", con.Register)
	//登录
	r.POST("/login", con.Login)
	//登出
	r.POST("/logout", con.Logout)
	//查询用户的类型
	r.POST("/getInfo", middleware.JWTAuthMiddleware(), con.GetInfo)
	//茅台酒上链
	r.POST("/uplink", middleware.JWTAuthMiddleware(), con.Uplink)
	// 获取茅台酒的上链信息
	r.POST("/getMoutaiInfo", con.GetMoutaiInfo)
	// 获取用户的茅台酒ID列表
	r.POST("/getMoutaiList", middleware.JWTAuthMiddleware(), con.GetMoutaiList)
	// 获取所有的keyword2
	r.POST("/getAllMoutaiInfo", middleware.JWTAuthMiddleware(), con.GetAllMoutaiInfo)
	// 获取茅台酒上链历史(溯源)
	r.POST("/getMoutaiHistory", middleware.JWTAuthMiddleware(), con.GetMoutaiHistory)
	return r
}
