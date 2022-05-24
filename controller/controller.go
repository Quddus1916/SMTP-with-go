package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sendemail.com/config"
	"sendemail.com/email"
	"sendemail.com/msg"
	"sendemail.com/types"
)

func SendMail(c echo.Context) error {

	var emailinfo = new(types.EmailSpec)
	config := config.InitConfig()
	if err := c.Bind(emailinfo); err != nil {
		return c.JSON(http.StatusBadRequest, "error in binding")
	}
	//make a header
	header := "welcome" + emailinfo.Username

	//make a body
	body := msg.CreateMsg(emailinfo.Username)
	//send mail
	err := email.SendEmail(emailinfo, config, header, body)
	if err != nil {
		return c.String(http.StatusOK, "email not send")
	}

	return c.String(http.StatusOK, "mail sent")
}
