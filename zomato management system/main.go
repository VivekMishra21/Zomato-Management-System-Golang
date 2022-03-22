package main

import (
	"fmt"

	Config "zomato/config"
	Models "zomato/models"
	Routes "zomato/routes"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})

	Config.DB.AutoMigrate(&Models.Restaurant{})

	Config.DB.AutoMigrate(&Models.Order{})

	Config.DB.AutoMigrate(&Models.OrderData{})

	Config.DB.AutoMigrate(&Models.FoodItems{})

	Config.DB.AutoMigrate(&Models.City{})

	router := gin.New()
	router.Use(gin.Logger())

	Routes.UserRoutes(router)
	Routes.RestaurantRoutes(router)
	Routes.OrderRoutes(router)
	Routes.OrderDataRoutes(router)
	Routes.FoodItemsRoutes(router)
	Routes.CityRoutes(router)
	router.Run("localhost:8000")

}
