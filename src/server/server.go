package server

import (
	"context"

	pb "github.com/mahendraHegde/email-service"
	"github.com/mahendraHegde/email-service/src/config"
	emailService "github.com/mahendraHegde/email-service/src/service/email"
	emailUtils "github.com/mahendraHegde/email-service/src/service/email/utils"
)

//Server Exported
type Server struct {
	pb.UnimplementedEmailServer
	MailService *emailService.MailClient
	config      config.Configurations
}

//SendContactMeEmail Exported
func (s *Server) SendContactMeEmail(ctx context.Context, req *pb.ContactMeEmailRequest) (*pb.ContactMeEmailReply, error) {
	data := emailUtils.NewContactMeMaildata(req.GetName(), req.GetEmail(), req.GetDetail(), req.GetSubject(), s.config.Me.Email, s.config.MailJet.Templates.ContactMe)
	info := emailUtils.GenerateEmailMessageInfo(data)
	s.MailService.SendMail(info)
	return &pb.ContactMeEmailReply{Message: "Success"}, nil
}
