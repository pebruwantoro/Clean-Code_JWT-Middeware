package main

import (
	"pebruwantoro/middleware/config"
	m "pebruwantoro/middleware/middlewares"
	"pebruwantoro/middleware/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.InitDB()
	// implement middleware logger
	m.LogMiddlewares(e)
	routes.New(e)
	//start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
