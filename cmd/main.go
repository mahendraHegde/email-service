package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	conf "github.com/mahendraHegde/email-service/src/config"
	"github.com/mahendraHegde/email-service/src/graph"
	"github.com/mahendraHegde/email-service/src/graph/generated"
	emailService "github.com/mahendraHegde/email-service/src/service/email"
)

const defaultPort = "8080"

func main() {

	config, err := conf.LoadConfig("./")
	if err != nil {
		fmt.Printf("Unable to read config, %v", err)
	}

	//Create MailClient
	mailService := emailService.NewMailClient(config.MailJet)

	resolvers := &graph.Resolver{
		MailService: mailService,
		Config:      config,
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%v/ for GraphQL playground", config.Server.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Server.Port), nil))
}
