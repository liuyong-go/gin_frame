package routers

import (
	"gin_frame/config"
	"gin_frame/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	gin.SetMode(config.LoadConfig().RunMode)
	r.GET("hello", controllers.HelloWorld)
	return r
}
