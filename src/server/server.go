package server

import (
	"context"
	"fmt"

	pb "github.com/mahendraHegde/email-service"
)

type Server struct {
	pb.UnimplementedEmailServer
}

func (s *Server) SendContactMeEmail(ctx context.Context, req *pb.ContactMeEmailRequest) (*pb.ContactMeEmailReply, error) {
	return &pb.ContactMeEmailReply{Message: fmt.Sprintf("%v", req)}, nil
}
