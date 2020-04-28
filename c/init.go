package c

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
	"server/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	gin.SetMode(setting.RunMode)

	router := r.Group("")
	{
		router.POST("/signup", Signup) //注册
		router.POST("/login", Login)
	}

	router_jwted := r.Group("")
	router_jwted.Use(middleware.Jwt())
	{
		router_jwted.POST("/logout", Logout)
		router_jwted.GET("/getuserinfo", GetUserInfo)
		router_jwted.POST("/post", Post)
		router_jwted.GET("/task/:category", GetTaskList)
		router_jwted.GET("/verifyCodeImage", GetverifyCodeImage)
	}

	return r
}
