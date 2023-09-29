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

	// subject := "Subject: Hare Krsna Email"
	// link := "<a href=\"" + verificationLink + "\">Click here to verify</a>"
	// msg := subject + "\n\nPlease " + link

	msg := "Subject: Hare Krsna Email\n Please click on verification link"

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
