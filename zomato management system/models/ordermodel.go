package models
import(
	"fmt"
	Config"zomato/config"

	_"github.com/go-sql-driver/mysql"
)




type Order struct {
	OrderId      int `json:"orderid"`
	UserId       int `json:"userid"`
	RestaurantId int `json:"restaurantid"`
	CityId       int `json:"cityid"`
	TotalPrice   int `json:"totalprice"`
}

func GetAllOrder(order *[]Order) (err error) {
	if err = Config.DB.Find(order).Error; err != nil {
		return err
	}
	return nil
}

func CreateAOrder(order *Order) (err error) {
	if err = Config.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func GetAOrder(order *Order, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAOrder(order *Order, id string) (err error) {
	fmt.Println(order)
	Config.DB.Save(order)
	return nil
}

func DeleteAOrder(order *Order, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(order)
	return nil
}




