package routes

import (
	"github.com/b4ysolutions/loki/controller"

	"github.com/labstack/echo"
)

// UserRoutes is used to set all user routes
func UserRoutes(echo *echo.Echo) {
	echo.GET("/users", controller.ListUsers)
	echo.GET("/users/:id", controller.GetUsers)
	echo.POST("/users", controller.CreateUsers)
	echo.PUT("/users", controller.UpdateUsers)
	echo.DELETE("/users/:id", controller.DeleteUsers)
	echo.PUT("/users/login", controller.UsersLogin)
	echo.GET("/users/logout/:id", controller.UsersLogout)
}
