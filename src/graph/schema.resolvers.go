package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mahendraHegde/email-service/src/graph/generated"
	"github.com/mahendraHegde/email-service/src/graph/model"
)

func (r *mutationResolver) SendContactUsEmail(ctx context.Context, input model.ContactUsEmail) (*model.Response, error) {
	toEmail := r.Config.HegdeFlutes.Email
	templateId := r.Config.MailJet.Templates.HegdeFlutesContactUs
	if input.To == model.ToMahendra {
		toEmail = r.Config.Me.Email
		templateId = r.Config.MailJet.Templates.ContactMe
	}
	data := model.NewContactMeMaildata(input.Name, input.Email, input.Detail, input.Subject, toEmail, templateId)
	if err := r.MailService.SendMail(model.GenerateEmailMessageInfo(data)); err != nil {
		return nil, err
	}
	return &model.Response{Status: "Success"}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
