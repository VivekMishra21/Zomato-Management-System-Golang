package routes

import (
	controller "zomato/controllers"

	"github.com/gin-gonic/gin"
)

func OrderDataRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderdata", controller.GetOrderData)
	incomingRoutes.POST("/orderdata", controller.CreateAOrderData)
	incomingRoutes.GET("/orderdata/:orderdata_id", controller.GetAOrderData)

	incomingRoutes.PATCH("/orderdata/:orderdata_id", controller.UpdateAOrderData)
	incomingRoutes.DELETE("/orderdata/:orderdata_id", controller.DeleteAOrderData)
}