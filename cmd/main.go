package main

import (
	"fmt"
	"log"
	"net"
	"strconv"

	pb "github.com/mahendraHegde/email-service"
	conf "github.com/mahendraHegde/email-service/src/config"
	"github.com/mahendraHegde/email-service/src/server"
	emailService "github.com/mahendraHegde/email-service/src/service/email"
	"google.golang.org/grpc"
)

func main() {
	config, err := conf.LoadConfig("./")
	if err != nil {
		fmt.Printf("Unable to read config, %v", err)
	}

	//Create MailClient
	mailService := emailService.NewMailClient(config.MailJet)

	//create Server
	server := server.Server{MailService: mailService}

	port := ":" + strconv.Itoa(config.Server.Port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEmailServer(s, &server)
	log.Printf("Listening on PORT %v", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
