package graph

import (
	"github.com/mahendraHegde/email-service/src/config"
	"github.com/mahendraHegde/email-service/src/service/email"
)

//Resolver Exported
type Resolver struct {
	MailService *email.MailClient
	Config      config.Configurations
}
