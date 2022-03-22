package routes

import (
	controller"zomato/controllers"

	"github.com/gin-gonic/gin"
)

func RestaurantRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/restaurants", controller. GetRestaurant)
	incomingRoutes.POST("/restaurant", controller.CreateAResaturant)
	incomingRoutes.GET("/restaurants/:restaurant_id", controller. GetARestaurant)

	incomingRoutes.PATCH("/restaurants/:restaurant_id", controller.UpdateARestaurant)
	incomingRoutes.DELETE("/restaurants/:restaurant_id", controller.DeleteARestaurant)
}
