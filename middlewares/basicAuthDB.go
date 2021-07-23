package middlewares

import (
	"pebruwantoro/middleware/config"
	"pebruwantoro/middleware/models"

	"github.com/labstack/echo"
)

// Basic Authentication Database
func BasicAuthDb(username, password string, c echo.Context) (bool, error) {
	db := config.DB
	var user models.User
	if err := db.Where("email=? AND password=?", username, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, nil
}
