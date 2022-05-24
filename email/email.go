package email

import (
	"net/smtp"
	"sendemail.com/config"
	"sendemail.com/types"
)

func SendEmail(email_info *types.EmailSpec, config config.Config, subject string, body string) error {

	to := []string{email_info.Email}

	address := config.SMTP_HOST + ":" + config.SMTP_PORT
	//build msg
	message := []byte(subject + body)
	//create auth

	auth := smtp.PlainAuth("", config.From, config.SMTP_PASSWORD, config.SMTP_HOST)

	// send mail
	err := smtp.SendMail(address, auth, config.From, to, message)
	if err != nil {
		return err
	}

	return err

}
