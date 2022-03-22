package Controllers

import (
	"net/http"
	Models"zomato/models"
	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	var order []Models.Order
	err := Models.GetAllOrder(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}


func CreateAOrder(c *gin.Context) {
	var order Models.Order
	c.BindJSON(&order)
	err := Models.CreateAOrder(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}


func GetAOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var order Models.Order
	err := Models.GetAOrder(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}


func UpdateAOrder(c *gin.Context) {
	var order Models.Order
	id := c.Params.ByName("id")
	err := Models.GetAOrder(&order, id)
	if err != nil {
		c.JSON(http.StatusNotFound, order)
	}
	c.BindJSON(&order)
	err = Models.UpdateAOrder(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func DeleteAOrder(c *gin.Context) {
	var order Models.Order
	id := c.Params.ByName("id")
	err := Models.DeleteAOrder(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}