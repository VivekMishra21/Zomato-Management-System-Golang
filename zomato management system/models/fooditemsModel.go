package models

import (
	"fmt"
	Config "zomato/config"

	_ "github.com/go-sql-driver/mysql"
)

type FoodItems struct {
	ItemsId      int    `json:"itemsid"`
	RestaurantId int    `json:"restaurantid"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
}

func GetAllFoodItems(fooditems *[]FoodItems) (err error) {
	if err = Config.DB.Find(fooditems).Error; err != nil {
		return err
	}
	return nil
}

func CreateAFoodItems(fooditems *FoodItems) (err error) {
	if err = Config.DB.Create(fooditems).Error; err != nil {
		return err
	}
	return nil
}

func GetAFoodItems(fooditems *FoodItems, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(fooditems).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAFoodItems(fooditems *FoodItems, id string) (err error) {
	fmt.Println(fooditems)
	Config.DB.Save(fooditems)
	return nil
}

func DeleteAFoodItems(fooditems *FoodItems, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(fooditems)
	return nil
}
