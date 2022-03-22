package routes

import (
	controller "zomato/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers)
	incomingRoutes.POST("/user", controller.CreateAUser)
	incomingRoutes.GET("/users/:user_id", controller.GetAUser)
	incomingRoutes.PATCH("/users/:user_id", controller.UpdateAUser)
	incomingRoutes.DELETE("/users/:user_id", controller.DeleteAUser)
}
