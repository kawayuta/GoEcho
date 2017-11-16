package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	. "./handler"
)


func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes

	e.POST("/users/", InsertUser)

	e.GET("/user/:id", SelectUser)

	e.GET("/users", SelectUsers)
	e.GET("/users/", SelectUsers)

	e.PUT("/users/", UpdateUser)

	e.DELETE("/users/:id/", DeleteUser)

	// Start server
	e.Start(":1323")
}
