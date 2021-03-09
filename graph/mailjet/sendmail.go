package mailjet

import (
	"fmt"
	"github.com/kendricko-adrio/gqlgen-todos/graph/model"
	mailjet "github.com/mailjet/mailjet-apiv3-go"
	"log"
	"strconv"
)

func SendEmail(toEmail, otp string) {
	mailjetClient := mailjet.NewMailjetClient("ede341987c435df63f74d5ed23686b6d", "f53c24ff587dc8b1e44352de917250e5")
	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: "kendrickoadrio314@gmail.com",
				Name:  "Kendricko",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: toEmail,
				},
			},
			Subject:  "Staem OTP",
			TextPart: "My first Mailjet email",
			HTMLPart: "<h3>" + otp + "</h3>",
			CustomID: "AppGettingStartedTest",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}

func OnSale(toEmail, game string) {
	mailjetClient := mailjet.NewMailjetClient("ede341987c435df63f74d5ed23686b6d", "f53c24ff587dc8b1e44352de917250e5")
	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: "kendrickoadrio314@gmail.com",
				Name:  "Kendricko",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: toEmail,
				},
			},
			Subject:  "Staem OTP",
			TextPart: "My first Mailjet email",
			HTMLPart: "<h3>" + game + "is Now on Sale! grab it fast" + "</h3>",
			CustomID: "AppGettingStartedTest",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}

func Transaction(toEmail string, detail []*model.Game) {
	mailjetClient := mailjet.NewMailjetClient("ede341987c435df63f74d5ed23686b6d", "f53c24ff587dc8b1e44352de917250e5")

	var tr string

	for i := 0; i < len(detail); i++ {
		tr = tr + "<tr>" +
			"<td>" + detail[i].Name + "</td>" +
			"<td>" + strconv.Itoa(detail[i].Price) + "</td>" +
			"</tr>"
	}

	msg := "<!DOCTYPE html>\n" +
		"<html lang=\"en\">\n" +
		"<head>\n    " +
		"<meta charset=\"UTF-8\">\n    " +
		"<meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\n    " +
		"<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n    " +
		"<title>Document</title>\n    " +
		"<style>\n        " +
		".card{\n            " +
		"padding: 20px;\n            " +
		"width: 300px;\n            " +
		"height: 100%;\n            " +
		"background-color: gray;\n        " +
		"}\n    " +
		"</style>\n" +
		"</head>\n" +
		"<body>\n    " +
		"<div class=\"card\">\n       " +
		" <div>Hey Kamu! Thank you for purchasing our product</div>\n        " +
		"<table>\n            " +
		"<tr>\n                " +
		"<td>Game</td>\n               " +
		" <td>Price</td>\n            " +
		"</tr>\n" +
		tr +
		"</table>\n    " +
		"</div>\n" +
		"</body>\n" +
		"</html>"

	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: "kendrickoadrio314@gmail.com",
				Name:  "Kendricko",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: toEmail,
				},
			},
			Subject:  "Staem OTP",
			TextPart: "My first Mailjet email",
			HTMLPart: msg,
			CustomID: "AppGettingStartedTest",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}
