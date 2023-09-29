package emailverification

import (
	"fmt"
	"net/smtp"

	"github.com/hktrib/SlackPlus/util"
)

func SendMail(config *util.Config, recipientEmailAddress *string) {
	auth := smtp.PlainAuth(
		"",
		config.SendingEmailAddress,
		config.SendingEmailAddressPwd,
		"smtp.gmail.com",
	)

	msg := "Subject: Hare Krsna Email\nEmail Body Woohooo"

	var recipientEa []string
	recipientEa = append(recipientEa, *recipientEmailAddress)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		config.SendingEmailAddress,
		recipientEa,
		[]byte(msg),
	)
	if err != nil {
		fmt.Println(err)
	}
}
