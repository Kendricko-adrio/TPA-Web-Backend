package mailjet

import (
	"fmt"
	mailjet "github.com/mailjet/mailjet-apiv3-go"
	"log"
)
func SendEmail (toEmail, otp string) {
	mailjetClient := mailjet.NewMailjetClient("ede341987c435df63f74d5ed23686b6d", "f53c24ff587dc8b1e44352de917250e5")
	messagesInfo := []mailjet.InfoMessagesV31 {
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: "kendrickoadrio314@gmail.com",
				Name: "Kendricko",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31 {
					Email: toEmail,
				},
			},
			Subject: "Staem OTP",
			TextPart: "My first Mailjet email",
			HTMLPart: "<h3>" + otp + "</h3>",
			CustomID: "AppGettingStartedTest",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo }
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}

func OnSale (toEmail, game string) {
	mailjetClient := mailjet.NewMailjetClient("ede341987c435df63f74d5ed23686b6d", "f53c24ff587dc8b1e44352de917250e5")
	messagesInfo := []mailjet.InfoMessagesV31 {
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: "kendrickoadrio314@gmail.com",
				Name: "Kendricko",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31 {
					Email: toEmail,
				},
			},
			Subject: "Staem OTP",
			TextPart: "My first Mailjet email",
			HTMLPart: "<h3>" + game + "is Now on Sale! grab it fast" + "</h3>",
			CustomID: "AppGettingStartedTest",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo }
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}
