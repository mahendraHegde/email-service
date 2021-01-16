package mail

import (
	"email-service/config"
	logger "github.com/sirupsen/logrus"
	mailjet "github.com/mailjet/mailjet-apiv3-go"
	proto "email-service/email_proto"
)

type mailClient struct {
	Client mailjet.Client
}
type ContactMeMailData struct {
	Name       *string
	Email      *string
	Detail     *string
}

func GetContactMeEmailvailables(vars proto.ContactMeEmailRequest)map[string]string{
	return map[string]string{
		
	}
}

func NewMailClient(conf config.MailJetConfig) *mailClient {
	return &mailClient{Client: NewMailjetClient(conf.ApiKey.Public, conf.ApiKey.Private)}
}
func (c *mailClient) SendMail(info []mailjet.InfoMessagesV31) {
	msgs:=mailjet.MessagesV31{Info: info}
	res, err := c.Client.SendMailV31(&msgs)
	if err != nil {
		logger.Errorf("Failed to send email %v",err)
		return
	}
	logger.Info("Email Successfuly sent to %v",info[0].To)
}


func main () {
	mailjetClient := NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))
	messagesInfo := []mailjet.InfoMessagesV31 {
      mailjet.InfoMessagesV31{
        From: &mailjet.RecipientV31{
          Email: "kebmizsrxfpbejvkpq@ttirv.com",
          Name: "hegde flutes",
        },
        To: &mailjet.RecipientsV31{
          mailjet.RecipientV31 {
            Email: "passenger1@example.com",
            Name: "passenger 1",
          },
        },
        TemplateID: 2224822,
        TemplateLanguage: true,
        Subject: "hegdflutes - contact us",
        Variables: map[string]interface{}{
      "subject": "\"NA\"",
      "name": "\"\"",
      "detail": "\"NA:",
      "email": "\"\""
    },
  },
	messages := mailjet.MessagesV31{Info: messagesInfo }
	res, err := m.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v
", res)
}
}