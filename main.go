package main

import (
	//"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	//"os"
	"sendemail.com/config"
	"sendemail.com/controller"
)

func main() {
	e := echo.New()

	config := config.InitConfig()
	port := config.Port
	e.GET("/sendmail", controller.SendMail)

	e.Start(":" + port)
}
