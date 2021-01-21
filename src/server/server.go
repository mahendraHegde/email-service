package server

import (
	"context"

	pb "github.com/mahendraHegde/email-service"
	"github.com/mahendraHegde/email-service/src/config"
	emailService "github.com/mahendraHegde/email-service/src/service/email"
	emailUtils "github.com/mahendraHegde/email-service/src/service/email/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//Server Exported
type Server struct {
	pb.UnimplementedEmailServer
	MailService *emailService.MailClient
	Config      config.Configurations
}

//SendContactMeEmail Exported
func (s *Server) SendContactMeEmail(ctx context.Context, req *pb.ContactMeEmailRequest) (*pb.ContactMeEmailReply, error) {
	data := emailUtils.NewContactMeMaildata(req.GetName(), req.GetEmail(), req.GetDetail(), req.GetSubject(), s.Config.Me.Email, s.Config.MailJet.Templates.ContactMe)
	info := emailUtils.GenerateEmailMessageInfo(data)
	if err := s.MailService.SendMail(info); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to send email")
	}
	return &pb.ContactMeEmailReply{Message: "Success"}, nil
}
