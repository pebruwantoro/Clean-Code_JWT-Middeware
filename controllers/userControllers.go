package controllers

import (
	"net/http"
	"pebruwantoro/middleware/config"
	"pebruwantoro/middleware/lib/database"
	"pebruwantoro/middleware/middlewares"
	"pebruwantoro/middleware/models"
	"strconv"

	"github.com/labstack/echo"
)

// Getting all users without authentication JWT
func GetUsersControllers(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

// Getting user by id without authentication JWT
func GetUserControllers(c echo.Context) error {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	var count int64
	config.DB.Model(models.User{}).Where("id=?", id).Count(&count)
	if count == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "not found",
		})
	}
	//Getting user by id
	getUser, err := database.GetOneUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get user by id",
		"users":  getUser,
	})
}

// Creating new user without authentication JWT
func CreateUserControllers(c echo.Context) error {
	var user models.User
	//Binding input data
	c.Bind(&user)
	newUser, err := database.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success create new user",
		"users":  newUser,
	})
}

// Deleting user without authentication JWT
func DeleteUserControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	delete_user, err := database.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success delete user",
		"users":  delete_user,
	})
}

// Updating user without authentication JWT
func UpdateUserControllers(c echo.Context) error {
	var user models.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	//Getting user by id
	get_user, _ := database.GetOneUser(id)
	user = get_user
	// Replacing user by id with new user data
	c.Bind(&user)
	update_user, err := database.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success update user",
		"users":  update_user,
	})
}

//Login without authentication JWT
func LoginUsersController(c echo.Context) error {
	userData := models.User{}
	c.Bind(&userData)
	users, err := database.LoginUsers(userData.Email, userData.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})
}

// Getting user by id with authentication JWT
func GetUserDetailControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Access user data use Token JWT
	loggedInUserId := middlewares.ExtractTokenUserId(c)
	if loggedInUserId != id {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized access, can not access the other databases")
	}
	// getting own user data
	users, err := database.GetDetailUsers((id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

// Deleting user by id with authentication JWT
func DeleteOneUserControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Access user data use Token JWT
	loggedInUserId := middlewares.ExtractTokenUserId(c)
	if loggedInUserId != id {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized access, can not access the other databases")
	}
	// Deleting process
	delete_user, err := database.DeleteOneUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  delete_user,
	})
}

//Updating user data
func UpdateOneUserControllers(c echo.Context) error {
	var user models.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	// Access user data use Token JWT
	loggedInUserId := middlewares.ExtractTokenUserId(c)
	if loggedInUserId != id {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized access, can not access the other databases")
	}
	// Getting user data
	get_user, _ := database.GetDetailUsers(id)
	user = get_user
	// Replacing user data with new data
	c.Bind(&user)
	update_user, err := database.UpdateOneUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success update user",
		"users":  update_user,
	})
}
