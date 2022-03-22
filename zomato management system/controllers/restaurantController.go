package Controllers

import (
	"net/http"
	Models"zomato/models"
	"github.com/gin-gonic/gin"
)

func GetRestaurant(c *gin.Context) {
	var restaurant []Models.Restaurant
	err := Models.GetAllRestaurant(&restaurant)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, restaurant)
	}
}


func CreateAResaturant(c *gin.Context) {
	var restaurant Models.Restaurant
	c.BindJSON(&restaurant)
	err := Models.CreateAResaturant(&restaurant)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, restaurant)
	}
}


func GetARestaurant(c *gin.Context) {
	id := c.Params.ByName("id")
	var restaurant Models.Restaurant
	err := Models.GetARestaurant(&restaurant, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, restaurant)
	}
}


func UpdateARestaurant(c *gin.Context) {
	var restaurant Models.Restaurant
	id := c.Params.ByName("id")
	err := Models.GetARestaurant(&restaurant, id)
	if err != nil {
		c.JSON(http.StatusNotFound, restaurant)
	}
	c.BindJSON(&restaurant)
	err = Models.UpdateARestaurant(&restaurant, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, restaurant)
	}
}

func DeleteARestaurant(c *gin.Context) {
	var restaurant Models.Restaurant
	id := c.Params.ByName("id")
	err := Models.DeleteARestaurant(&restaurant, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}