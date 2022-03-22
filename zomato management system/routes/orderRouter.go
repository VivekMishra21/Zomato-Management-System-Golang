package routes

import (
	controller "zomato/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controller.GetOrders)
	incomingRoutes.POST("/order", controller.CreateAOrder)
}
