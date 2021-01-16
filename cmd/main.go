package main

import (
	"fmt"

	conf "github.com/mahendraHegde/email-service/src/config"
)

func main() {
	config, err := conf.LoadConfig("./")
	if err != nil {
		fmt.Printf("Unable to read config, %v", err)
	}
	fmt.Println(config)

	// port := ":" + viper.GetString("port")
	// if port == ":" {
	// 	port = ":" + strconv.Itoa(configuration.ServerPort)
	// }
	// port := ":" + strconv.Itoa(configuration.Server.Port)

	// lis, err := net.Listen("tcp", port)
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// s := grpc.NewServer()
	// pb.RegisterEmailServer(s, &server.Server{})
	// log.Printf("Listening on PORT %v", port)
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
}
