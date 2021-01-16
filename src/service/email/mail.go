package email

import (
	"github.com/mahendraHegde/email-service/src/config"
	mailjet "github.com/mailjet/mailjet-apiv3-go"
	logger "github.com/sirupsen/logrus"
)

type MailClient struct {
	Client *mailjet.Client
}

//NewMailClient exported
func NewMailClient(conf config.MailJetConfig) *MailClient {
	return &MailClient{Client: mailjet.NewMailjetClient(conf.ApiKey.Public, conf.ApiKey.Private)}
}

//SendMail exported
func (c *MailClient) SendMail(info []mailjet.InfoMessagesV31) (err error) {
	msgs := mailjet.MessagesV31{Info: info}
	res, err := c.Client.SendMailV31(&msgs)
	if err != nil {
		logger.Errorf("Failed to send email %v", err)
		return
	}
	logger.Info("Email Successfuly sent to %v,res %v", info[0].To, res)
	return
}
