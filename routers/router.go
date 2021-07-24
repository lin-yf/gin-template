package routers

import (
	"go-template/pkg/middlewares"
	"os"

	"go-template/routers/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middlewares.Session(string(os.Getenv("SECURE_KEY"))))
	r.Use(middlewares.CORSMiddleware())

	// api
	api := r.Group("/api")
	api.Use(middlewares.CORSMiddleware())
	api.POST("/auth/login", controllers.LoginPost)
	authApi := api.Use(middlewares.JWTAuth())
	{
		authApi.GET("/user/nav", controllers.GetCurrentNav)
	}
	return r
}
