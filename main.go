package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	conf "github.com/mahendraHegde/email-service/src/config"
	"github.com/mahendraHegde/email-service/src/graph"
	"github.com/mahendraHegde/email-service/src/graph/generated"
	emailService "github.com/mahendraHegde/email-service/src/service/email"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {

	config, err := conf.LoadConfig("./")
	if err != nil {
		fmt.Printf("Unable to read config, %v", err)
	}

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://0.0.0.0:5000", "https://hegdeflutes.com/"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	//Create MailClient
	mailService := emailService.NewMailClient(config.MailJet)

	resolvers := &graph.Resolver{
		MailService: mailService,
		Config:      config,
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers}))

	router.Handle("/", playground.Handler("Notification services", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%v/ for GraphQL playground", config.Server.Port)
	err = http.ListenAndServe(":"+strconv.Itoa(config.Server.Port), router)
	if err != nil {
		panic(err)
	}
}
