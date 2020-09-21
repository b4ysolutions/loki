package controller

import (
	"net/http"

	"github.com/b4ysolutions/loki/helpers"
	"github.com/b4ysolutions/loki/helpers/mysql"
	"github.com/b4ysolutions/loki/model"
	"github.com/b4ysolutions/loki/service"
	"github.com/labstack/echo"
)

// ListUsers - Loki.GET("/users", ListUsers)
func ListUsers(context echo.Context) error {

	db, err := mysql.Connect()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, helpers.Output("ErrorDBConnection", err))
	}
	defer db.Close()

	return context.JSON(http.StatusOK, service.ListUsers(db))
}

// GetUsers - Loki.GET("/users/:id", GetUsers)
func GetUsers(context echo.Context) error {

	db, err := mysql.Connect()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, helpers.Output("ErrorDBConnection", err))
	}
	defer db.Close()

	return context.JSON(http.StatusOK, service.GetUsers(db, context.Param("id")))
}

// CreateUsers - Loki.POST("/users", CreateUsers)
func CreateUsers(context echo.Context) error {

	var user = new(model.User)

	db, err := mysql.Connect()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, helpers.Output("ErrorDBConnection", err))
	}
	defer db.Close()

	if err := context.Bind(user); err != nil {
		return context.JSON(http.StatusInternalServerError, helpers.Output("ErrorBindJsonInputData", err))
	}

	return context.JSON(http.StatusCreated, service.CreateUsers(db, *user))
}

// UpdateUsers - Loki.PUT("/users", UpdateUsers)
func UpdateUsers(context echo.Context) error {

	var user = new(model.User)

	db, err := mysql.Connect()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, helpers.Output("ErrorDBConnection", err))
	}
	defer db.Close()

	if err := context.Bind(user); err != nil {
		return context.JSON(http.StatusInternalServerError, helpers.Output("ErrorBindJsonInputData", err))
	}

	return context.JSON(http.StatusCreated, service.UpdateUsers(db, user))
}

// DeleteUsers - Loki.DELETE("/users/:id", DeleteUsers)
func DeleteUsers(context echo.Context) error {

	db, err := mysql.Connect()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, helpers.Output("ErrorDBConnection", err))
	}
	defer db.Close()

	return context.JSON(http.StatusOK, service.DeleteUsers(db, context.Param("id")))

}

// UsersLogin - Loki.PUT("/users/login", UsersLogin)
func UsersLogin(context echo.Context) error {

	var login = new(model.Login)

	db, err := mysql.Connect()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, helpers.Output("ErrorDBConnection", err))
	}
	defer db.Close()
	if err := context.Bind(login); err != nil {
		return context.JSON(http.StatusInternalServerError, helpers.Output("ErrorBindJsonInputData", err))
	}

	return context.JSON(http.StatusCreated, service.UsersLogin(db, *login))
}

// UsersLogout - Loki.GET("/users/login/:id", UsersLogout)
func UsersLogout(context echo.Context) error {

	db, err := mysql.Connect()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, helpers.Output("ErrorDBConnection", err))
	}
	defer db.Close()

	return context.JSON(http.StatusCreated, service.UsersLogout(db, context.Param("id")))
}
