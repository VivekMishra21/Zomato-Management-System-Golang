package routes

import (
	controller "zomato/controllers"

	"github.com/gin-gonic/gin"
)

func CityRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/citys", controller.GetCitys)
	incomingRoutes.POST("/city", controller.CreateACity)
	incomingRoutes.GET("/citys/:city_id", controller.GetACity)

	incomingRoutes.PATCH("/citys/:city_id", controller.UpdateACity)
	incomingRoutes.DELETE("/citys/:city_id", controller.DeleteACity)
}
