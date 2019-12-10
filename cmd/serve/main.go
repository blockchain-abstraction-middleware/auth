package main

import (
	auth "github.com/blockchain-abstraction-middleware/auth/pkg/routes"
	log "github.com/blockchain-abstraction-middleware/rest-api/pkg/logger"
	"github.com/blockchain-abstraction-middleware/rest-api/pkg/routes"
	"github.com/blockchain-abstraction-middleware/rest-api/pkg/server"
)

func main() {
	// abstractionConfig := config.NewConfig()

	serverConfig := server.Config{
		BasePath:       "/api/v1",
		Name:           "auth",
		Port:           8080,
		MetricsEnabled: false,
	}
	srv := server.New(&serverConfig)

	healthResource := routes.HealthResource{}
	swaggerResource := routes.SwaggerResource{}
	authResource := auth.AuthResource{}

	srv.RegisterResource(healthResource.NewResource("/health"))
	srv.RegisterResource(swaggerResource.NewResource("/swagger", "/api/v1/swagger"))

	srv.RegisterResource(authResource.NewResource("/auth"))

	if err := srv.Run(); err != nil {
		log.WithError(err).Fatal("Serving failed")
	}
}
