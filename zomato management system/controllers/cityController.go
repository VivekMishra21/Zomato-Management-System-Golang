package Controllers

import (
	"net/http"
	Models "zomato/models"

	"github.com/gin-gonic/gin"
)

func GetCitys(c *gin.Context) {
	var citys []Models.City
	err := Models.GetAllCity(&citys)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, citys)
	}
}

func CreateACity(c *gin.Context) {
	var city Models.City
	c.BindJSON(&city)
	err := Models.CreateACity(&city)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, city)
	}
}

func GetACity(c *gin.Context) {
	id := c.Params.ByName("id")
	var city Models.Restaurant
	err := Models.GetARestaurant(&city, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, city)
	}
}

func UpdateACity(c *gin.Context) {
	var city Models.City
	id := c.Params.ByName("id")
	err := Models.GetACity(&city, id)
	if err != nil {
		c.JSON(http.StatusNotFound, city)
	}
	c.BindJSON(&city)
	err = Models.UpdateACity(&city, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, city)
	}
}

func DeleteACity(c *gin.Context) {
	var city Models.City
	id := c.Params.ByName("id")
	err := Models.DeleteACity(&city, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
