package model

// User represents the model for an user
type User struct {
	ID       int    `json:"id" gorm:"column:user_id;primary_key;AUTO_INCREMENT;not null"`
	Name     string `json:"name" gorm:"column:user_name;type:varchar(100);not null"`
	Email    string `json:"email" gorm:"column:user_email;type:varchar(100);not null;unique"`
	Password string `json:"password" gorm:"column:user_password;type:varchar(100);not null"`
	Status   int    `json:"status" gorm:"column:user_status;type:int;not null"`
	IsLogged bool   `json:"islogged" gorm:"column:user_islogged;type:boolean;not null"`
}
