package services

import (
	"go-twitter/database"
	"go-twitter/entities"
	"go-twitter/utils"
)

func CreateUser(user *entities.User) (*entities.User, error) {
	var err error

	res := database.Conn.Create(user)

	if res.Error != nil {
		err = res.Error
		utils.WarningLog.Println(err.Error())
	}

	return user, err
}

func GetUsers() (*[]entities.User, error) {
	var err error
	var users *[]entities.User

	res := database.Conn.Order("user_name").Find(&users)

	if res.Error != nil {
		err = res.Error
		utils.WarningLog.Println(err.Error())
	}
	return users, err
}

func GetUser(userId string) (*entities.User, error) {
	var err error
	var user *entities.User

	res := database.Conn.Find(&user, userId)

	if res.Error != nil {
		err = res.Error
		utils.WarningLog.Println(err.Error())
	}

	return user, err
}

func UpdateUser(newUser *entities.User, userId string) (*entities.User, error) {
	user, err := GetUser(userId)
	if err == nil {
		res := database.Conn.Model(user).Updates(newUser)

		if res.Error != nil {
			err = res.Error
			utils.WarningLog.Println(err.Error())
		}
	}
	return user, err
}

func DeleteUser(userId string) (*entities.User, error) {
	user, err := GetUser(userId)

	if err == nil {
		res := database.Conn.Delete(user)
		if res.Error != nil {
			err = res.Error
			utils.WarningLog.Println(err.Error())
		}
	}
	return user, err
}
