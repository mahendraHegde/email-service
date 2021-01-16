package utils

import (
	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

type ContactMeMailData struct {
	Name       string
	Email      string
	Detail     string
	Subject    string
	To         string
	TemplateID int
}

func NewContactMeMaildata(name, email, detail, subject, to string, templateID int) ContactMeMailData {
	return ContactMeMailData{
		Name:       name,
		Email:      email,
		Detail:     detail,
		Subject:    subject,
		To:         to,
		TemplateID: templateID,
	}
}

//GetContactMeEmailvailables exported
func GenerateEmailMessageInfo(data ContactMeMailData) []mailjet.InfoMessagesV31 {
	subject := data.Subject
	if subject == "" {
		subject = `"NA"`
	}
	name := data.Name
	if name == "" {
		name = `"NA"`
	}
	detail := data.Detail
	if detail == "" {
		detail = `"NA"`
	}
	email := data.Email
	if email == "" {
		email = `"NA"`
	}
	variables := map[string]interface{}{
		subject: subject,
		name:    name,
		detail:  detail,
		email:   email,
	}
	return []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			Variables: variables,
			From: &mailjet.RecipientV31{
				Email: email,
				Name:  name,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: data.To,
				},
			},
			TemplateID:       data.TemplateID,
			TemplateLanguage: true,
			Subject:          subject,
		},
	}
}
