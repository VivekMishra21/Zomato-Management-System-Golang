package models

import(
	"fmt"
	Config"zomato/config"

	_"github.com/go-sql-driver/mysql"
)


type User struct {
	UserId      int    `json:"userid"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber int    `json:"phonenumber"`
	Address     string `json:"address"`


}

func GetAllUsers(user *[]User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateAUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetAUser(user *User, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAUser(user *User, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}

func DeleteAUser(user *User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}




