package main

import (
	"jwtEx/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	controllers.Connect()
	api := r.Group("/v1")
	{
		api.POST("/user/register", controllers.RegisterUser)
		api.POST("/user/login", controllers.GenerateToken)
		api.POST("user/login1", controllers.GenerateTokenByXML)
		secured := api.Group("/secured")
		{
			secured.GET("/ping", controllers.GetUser)
			secured.GET("/getuser", controllers.GetUserInXML)
		}
	}
	r.Run()
}
