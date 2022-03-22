package Controllers
import (
	"net/http"
	Models"zomato/models"
	"github.com/gin-gonic/gin"
)

func GetFoodItems(c *gin.Context) {
	var foodItems []Models.FoodItems
	err := Models.GetAllFoodItems(&foodItems)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, foodItems)
	}
}


func CreateAFoodItems(c *gin.Context) {
	var foodItems Models.FoodItems
	c.BindJSON(&foodItems)
	err := Models.CreateAFoodItems(&foodItems)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, foodItems)
	}
}


func GetAFoodItems(c *gin.Context) {
	id := c.Params.ByName("id")
	var foodItems Models.FoodItems
	err := Models.GetAFoodItems(&foodItems, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, foodItems)
	}
}


func UpdateAFoodItems(c *gin.Context) {
	var foodItems Models.FoodItems
	id := c.Params.ByName("id")
	err := Models.GetAFoodItems(&foodItems, id)
	if err != nil {
		c.JSON(http.StatusNotFound, foodItems)
	}
	c.BindJSON(&foodItems)
	err = Models.UpdateAFoodItems(&foodItems, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, foodItems)
	}
}

func DeleteAFoodItems(c *gin.Context) {
	var fooditem Models.FoodItems
	id := c.Params.ByName("id")
	err := Models.DeleteAFoodItems(&fooditem, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}