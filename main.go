package main

import (
	"os"
	"time"

	goClacks "github.com/JonathanMH/goClacks/echo"
	"github.com/b4ysolutions/loki/config"
	"github.com/b4ysolutions/loki/helpers/mysql"
	"github.com/b4ysolutions/loki/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Init Configs
	config.Init()

	// Mysql Init
	mysql.Init()

	// Echo instance
	api := echo.New()
	api.Use(goClacks.Terrify)

	// Middleware
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())

	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Set User Routes
	routes.UserRoutes(api)

	// Set log output file
	file, err := os.OpenFile("./log/log_"+time.Now().Format("20060102_150405")+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	api.Logger.SetOutput(file)
	api.Logger.SetLevel(log.DEBUG)
	log.SetOutput(file)
	log.SetLevel(log.DEBUG)

	// Start Service
	api.Logger.Debug(api.Start(":8787"))
}
