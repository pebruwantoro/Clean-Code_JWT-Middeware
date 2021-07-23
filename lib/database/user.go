package database

import (
	"pebruwantoro/middleware/config"
	"pebruwantoro/middleware/middlewares"
	"pebruwantoro/middleware/models"
)

// Access all of user data without authentication
func GetUsers() (interface{}, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Access user data by id without authentication
func GetOneUser(id int) (models.User, error) {
	var user models.User
	if err := config.DB.Find(&user, "id=?", id).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Creating new user
func CreateUser(user models.User) (interface{}, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Deleting user data without authentication
func DeleteUser(id int) (interface{}, error) {
	var user models.User
	if err := config.DB.Where("id=?", id).First(&user).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Delete(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

//Updating user data without authentication
func UpdateUser(user models.User) (models.User, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

//Login user use authentication JWT Middlewares
func LoginUsers(email, password string) (interface{}, error) {
	var user models.User
	var err error
	if err = config.DB.Where("email = ? AND password =?", email, password).First(&user).Error; err != nil {
		return nil, err
	}
	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, err
}

// Getting user data by id with authentication
func GetDetailUsers(userId int) (models.User, error) {
	var user models.User
	if err := config.DB.Find(&user, userId).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Deleting user data with authentication
func DeleteOneUser(userId int) (interface{}, error) {
	var user models.User
	if err := config.DB.Where("id=?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Delete(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Updating user data with authentication
func UpdateOneUser(user models.User) (models.User, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
