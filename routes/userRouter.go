package routes

import (
	controller "github.com/abitiGG/Golang-Jwt-Project/controllers"

	"github.com/abitiGG/Golang-Jwt-Project/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
