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
	if len(name) == 0 {
		name = `"NA"`
	}
	detail := data.Detail
	if len(detail) == 0 {
		detail = `"NA"`
	}
	email := data.Email
	if len(email) == 0 {
		email = `"NA"`
	}
	variables := make(map[string]interface{})
	variables["subject"] = subject
	variables["name"] = name
	variables["detail"] = detail
	variables["email"] = email

	return []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			Variables: variables,
			From: &mailjet.RecipientV31{
				Email: data.To,
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
