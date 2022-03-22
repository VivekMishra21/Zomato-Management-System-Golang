package models

import(
	"fmt"
	Config"zomato/config"

	_"github.com/go-sql-driver/mysql"
)




type OrderData struct {
	OrderId  int    `json:"orderid"`
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
}

func GetAllOrderData(orderdata *[]OrderData) (err error) {
	if err = Config.DB.Find(orderdata).Error; err != nil {
		return err
	}
	return nil
}

func CreateAOrderData(orderdata *OrderData) (err error) {
	if err = Config.DB.Create(orderdata).Error; err != nil {
		return err
	}
	return nil
}

func GetAOrderData(orderdata *OrderData, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(orderdata).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAOrderData(orderdata *OrderData, id string) (err error) {
	fmt.Println(orderdata)
	Config.DB.Save(orderdata)
	return nil
}

func DeleteAOrderData(orderdata *OrderData, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(orderdata)
	return nil
}



