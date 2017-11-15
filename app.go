package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gocraft/dbr"
)


func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes

	e.POST("/users/", insertUser)

	e.GET("/user/:id", selectUser)

	e.GET("/users", selectUsers)
	e.GET("/users/", selectUsers)

	e.PUT("/users/", updateUser)

	e.DELETE("/users/:id/", deleteUser)

	// Start server
	e.Start(":1323")
}
