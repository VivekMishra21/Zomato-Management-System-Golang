package models

import (
	"fmt"
	Config "zomato/config"

	_ "github.com/go-sql-driver/mysql"
)

type City struct {
	CityId int    `json:"cityid"`
	Name   string `json:"name"`
}

func GetAllCity(city *[]City) (err error) {
	if err = Config.DB.Find(city).Error; err != nil {
		return err
	}
	return nil
}

func CreateACity(city *City) (err error) {
	if err = Config.DB.Create(city).Error; err != nil {
		return err
	}
	return nil
}

func GetACity(city *City, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(city).Error; err != nil {
		return err
	}
	return nil
}

func UpdateACity(city *City, id string) (err error) {
	fmt.Println(city)
	Config.DB.Save(city)
	return nil
}

func DeleteACity(city *City, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(city)
	return nil
}
