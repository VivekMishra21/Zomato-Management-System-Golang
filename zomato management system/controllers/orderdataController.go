package Controllers

import (
	"net/http"
	Models "zomato/models"

	"github.com/gin-gonic/gin"
)

func GetOrderData(c *gin.Context) {
	var orderdata []Models.OrderData
	err := Models.GetAllOrderData(&orderdata)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, orderdata)
	}
}

func CreateAOrderData(c *gin.Context) {
	var orderdata Models.OrderData
	c.BindJSON(&orderdata)
	err := Models.CreateAOrderData(&orderdata)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, orderdata)
	}
}

func GetAOrderData(c *gin.Context) {
	id := c.Params.ByName("id")
	var orderdata Models.OrderData
	err := Models.GetAOrderData(&orderdata, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, orderdata)
	}
}

func UpdateAOrderData(c *gin.Context) {
	var orderdata Models.OrderData
	id := c.Params.ByName("id")
	err := Models.GetAOrderData(&orderdata, id)
	if err != nil {
		c.JSON(http.StatusNotFound, orderdata)
	}
	c.BindJSON(&orderdata)
	err = Models.UpdateAOrderData(&orderdata, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, orderdata)
	}
}

func DeleteAOrderData(c *gin.Context) {
	var orderdata Models.OrderData
	id := c.Params.ByName("id")
	err := Models.DeleteAOrderData(&orderdata, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
