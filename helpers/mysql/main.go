package mysql

import (
	"github.com/b4ysolutions/loki/helpers"
	"github.com/b4ysolutions/loki/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

// Connect - Return a valid mysql database connection
func Connect() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", viper.GetString("mysql.username")+":"+viper.GetString("mysql.password")+"@("+viper.GetString("mysql.hostname")+":"+viper.GetString("mysql.port")+")/"+viper.GetString("mysql.database")+"?charset=utf8&parseTime=True&loc=Local")
	return db, err
}

// Init - Delete tables and create tables again
func Init() {

	db, err := Connect()
	if err != nil {
		log.Error(helpers.OutputMessage("ErrorDBConnection"))
	}
	defer db.Close()

	db.DropTableIfExists(&model.User{})
	log.Info("User Table Dropped")

	db.CreateTable(&model.User{})
	log.Info("User Table Created")
}
