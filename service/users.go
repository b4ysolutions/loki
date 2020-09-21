package service

import (
	"strconv"

	"github.com/b4ysolutions/loki/helpers"
	"github.com/b4ysolutions/loki/model"
	"github.com/jinzhu/gorm"
)

var response model.Output

// ListUsers - Used to retrieve all registered users from database
func ListUsers(db *gorm.DB) model.Output {

	var users []model.User
	var response model.Output

	db.Find(&users)

	if len(users) > 0 {
		response = helpers.Output("SuccessListWithAllUsers", nil)
	} else {
		response = helpers.Output("SuccessNoUserFound", nil)
	}
	response.Users = users

	return response
}

// GetUsers - Used to retrieve an user from database
func GetUsers(db *gorm.DB, requestedID string) model.Output {

	user := new(model.User)
	user.ID, _ = strconv.Atoi(requestedID)

	db.Take(&user)

	if user.Email != "" {
		response = helpers.Output("SuccessUserFound", nil)
		response.Users = append(response.Users, *user)
	} else {
		response = helpers.Output("SuccessUserNotFound", nil)
	}

	return response
}

// CreateUsers - Used to create a new user in database
func CreateUsers(db *gorm.DB, user model.User) model.Output {

	db.Create(&user)

	if user.ID > 0 {
		response = helpers.Output("SuccessUserCreated", nil)
		response.Users = append(response.Users, user)
	} else {
		response = helpers.Output("ErrorCreateUser", nil)
	}

	return response
}

// UpdateUsers - Used to update an user in database
func UpdateUsers(db *gorm.DB, user *model.User) model.Output {

	newUser := new(model.User)
	newUser.ID = user.ID

	db.Take(&newUser)

	if newUser.Email != "" {
		newUser = user
		db.Save(&newUser)
	}

	if newUser == user {
		response = helpers.Output("SuccessUserUpdated", nil)
		response.Users = append(response.Users, *user)
	} else {
		response = helpers.Output("SuccessUserNotUpdated", nil)
	}

	return response
}

// DeleteUsers - Used to delete an user in database
func DeleteUsers(db *gorm.DB, requestedID string) model.Output {

	user := new(model.User)
	user.ID, _ = strconv.Atoi(requestedID)

	db.First(&user)

	if db.Unscoped().Delete(&user).RowsAffected == 1 {
		response = helpers.Output("SuccessUserRemoved", nil)
	} else {
		response = helpers.Output("SuccessUserNotRemoved", nil)
	}

	return response
}

// UsersLogin - Used to login an user in database
func UsersLogin(db *gorm.DB, login model.Login) model.Output {

	user := new(model.User)

	db.Where("user_email = ? AND user_password = ?", login.Email, login.Password).Find(&user)

	user.IsLogged = true

	if db.Model(model.User{}).Updates(user).RowsAffected == 1 {
		response = helpers.Output("SuccessUserLogin", nil)
	} else {
		response = helpers.Output("SuccessUserAlreadyLogged", nil)
	}
	response.Users = append(response.Users, *user)

	return response
}

// UsersLogout - Used to logout an user in database
func UsersLogout(db *gorm.DB, requestedID string) model.Output {

	user := new(model.User)
	user.ID, _ = strconv.Atoi(requestedID)

	db.First(&user)

	if user.IsLogged == true {

		user.IsLogged = false

		db.Save(&user)

		response = helpers.Output("SuccessUserLogout", nil)
		response.Users = append(response.Users, *user)
	} else {
		response = helpers.Output("SuccessUserNotLogged", nil)
	}

	return response
}
