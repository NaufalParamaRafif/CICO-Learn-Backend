package main

import (
	"github.com/NaufalParamaRafif/CICO-Learn-Backend/controllers"
	"github.com/NaufalParamaRafif/CICO-Learn-Backend/initializers"
	"github.com/NaufalParamaRafif/CICO-Learn-Backend/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/user/signup", controllers.Signup)
	r.POST("/user/login", controllers.Login)
	r.POST("/user/logout", middleware.RequireAuth, controllers.Logout)
	r.GET("/user/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
