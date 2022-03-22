package routes

import (
	controller"zomato/controllers"

	"github.com/gin-gonic/gin"
)

func FoodItemsRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/fooditems", controller.GetFoodItems)
	incomingRoutes.POST("/fooditem", controller.CreateAFoodItems)
	incomingRoutes.GET("/fooditems/:fooditem_id", controller.GetAFoodItems)

	incomingRoutes.PATCH("/fooditems/:fooditem_id", controller.UpdateAFoodItems)
	incomingRoutes.DELETE("/fooditems/:fooditem_id", controller.DeleteAFoodItems)
}