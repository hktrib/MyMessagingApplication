package emailverification

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/hktrib/SlackPlus/util"
)

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

func SendMail(config *util.Config, recipientEmailAddress *string,
	username *string, secretCode *string) {
	auth := smtp.PlainAuth(
		"",
		config.SendingEmailAddress,
		config.SendingEmailAddressPwd,
		"smtp.gmail.com",
	)

	verificationLink := fmt.Sprintf("http://localhost:5173/verifyUser?username=%v&secret_code=%v", *username, *secretCode)

	subject := fmt.Sprintf("Subject: My-Messaging-App: ||Account Verification Email for User: %v ||", *username)
	link := "<a href=\"" + verificationLink + "\">Click here to verify</a>"
	// msg := subject + "\n\nPlease " + link

	var recipient []string
	recipient = append(recipient, *recipientEmailAddress)

	request := Mail{
		Sender:  config.SendingEmailAddress,
		To:      recipient,
		Subject: subject,
		Body:    link,
	}

	msg := BuildMessage(request)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		config.SendingEmailAddress,
		recipient,
		[]byte(msg),
	)
	if err != nil {
		fmt.Println(err)
	}
}
func BuildMessage(mail Mail) string {
	var msg strings.Builder

	fmt.Fprintf(&msg, "From: %s\r\n", mail.Sender)
	fmt.Fprintf(&msg, "To: %s\r\n", strings.Join(mail.To, ";"))
	fmt.Fprintf(&msg, "Subject: %s\r\n", mail.Subject)
	msg.WriteString("\r\n")
	msg.WriteString(mail.Body)

	return msg.String()
}
