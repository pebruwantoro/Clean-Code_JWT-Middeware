package routes

import (
	"pebruwantoro/middleware/constants"
	"pebruwantoro/middleware/controllers"
	"pebruwantoro/middleware/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

	//CRUD without authentication JWT Token
	e.GET("/users", controllers.GetUsersControllers)
	e.GET("/users/:id", controllers.GetUserControllers)
	e.POST("/users", controllers.CreateUserControllers)
	e.DELETE("/users/:id", controllers.DeleteUserControllers)
	e.PUT("/users/:id", controllers.UpdateUserControllers)

	//login to get JWT Token
	e.POST("/login", controllers.LoginUsersController)

	// CRUD with authentication JWT Token
	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/users/:id", controllers.GetUserDetailControllers)
	eJwt.DELETE("/users/:id", controllers.DeleteOneUserControllers)
	eJwt.PUT("/users/:id", controllers.UpdateOneUserControllers)

	// Basic authentication
	eAuth := e.Group("")
	eAuth.Use(middleware.BasicAuth(middlewares.BasicAuthDb))
}
