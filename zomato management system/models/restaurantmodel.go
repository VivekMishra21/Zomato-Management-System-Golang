package models

import (
	"fmt"
	Config "zomato/config"

	_ "github.com/go-sql-driver/mysql"
)

type Restaurant struct {
	RestaurantId int    `json:"restaurantid"`
	CityId       int    `json:"cityid"`
	Name         string `json:"name"`
	Landmark     string `json:"landmark"`
	Address      string `json:"address"`
	Rating       string `josn:"rating"`
}

func GetAllRestaurant(restaurant *[]Restaurant) (err error) {
	if err = Config.DB.Find(restaurant).Error; err != nil {
		return err
	}
	return nil
}

func CreateAResaturant(restaurant *Restaurant) (err error) {
	if err = Config.DB.Create(restaurant).Error; err != nil {
		return err
	}
	return nil
}

func GetARestaurant(restaurant *Restaurant, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(restaurant).Error; err != nil {
		return err
	}
	return nil
}

func UpdateARestaurant(restaurant *Restaurant, id string) (err error) {
	fmt.Println(restaurant)
	Config.DB.Save(restaurant)
	return nil
}

func DeleteARestaurant(restaurant *Restaurant, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(restaurant)
	return nil
}
