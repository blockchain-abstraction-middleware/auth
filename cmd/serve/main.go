package main

import (
	"github.com/blockchain-abstraction-middleware/auth/pkg/authorization"
	"github.com/blockchain-abstraction-middleware/auth/pkg/contracts/Authorization"
	db "github.com/blockchain-abstraction-middleware/auth/pkg/db"
	keygen "github.com/blockchain-abstraction-middleware/auth/pkg/keygen"
	auth "github.com/blockchain-abstraction-middleware/auth/pkg/routes"
	"github.com/blockchain-abstraction-middleware/game-jam-abstraction/pkg/ethereum"
	log "github.com/blockchain-abstraction-middleware/rest-api/pkg/logger"
	"github.com/blockchain-abstraction-middleware/rest-api/pkg/routes"
	"github.com/blockchain-abstraction-middleware/rest-api/pkg/server"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
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

	// privKey, err := crypto.HexToECDSA(r.config.Keys.Admin)
	// if err != nil {
	// 	log.WithError(err).Error("Failed to load key")
	// }

	ec, err := ethereum.CreateEthClient("wss://rinkeby.infura.io/ws/v3/" + "833f712dc639422ba10ead3cf74ee0bf")
	if err != nil {
		log.WithError(err).Error("Failed to initialize ethereum client")
	}

	authorization, err := authorization.NewAuthorization(ec, "0xce1077cb99709dba18333ceed9752e9bbc967dff", nil)
	if err != nil {
		log.WithError(err).Error("Failed to create game jam manager")
	}

	as := make(chan *Authorization.AuthorizationSubscribed)

	sub, err := authorization.Manager.WatchSubscribed(&bind.WatchOpts{}, as, nil)
	if err != nil {
		log.WithError(err).Error("Failed to watch events")
	}

	log.WithFields(log.Fields{"Contract": "Auth", "Address": authorization}).Info("Created auth abstraction")

	db := db.NewDB("./auth/db")

	log.Info("Started database")

	go listen(sub, db, as)

	log.Info("Started event listener")

	healthResource := routes.HealthResource{}
	swaggerResource := routes.SwaggerResource{}
	authResource := auth.AuthResource{}

	srv.RegisterResource(healthResource.NewResource("/health"))
	srv.RegisterResource(swaggerResource.NewResource("/swagger", "/api/v1/swagger"))

	srv.RegisterResource(authResource.NewResource("/auth", db))

	if err := srv.Run(); err != nil {
		log.WithError(err).Fatal("Serving failed")
	}
}

func listen(sub event.Subscription, db *db.DB, as chan *Authorization.AuthorizationSubscribed) {
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case event := <-as:
			log.Info("User subscribed: " + event.Account.String())
			err := db.PutData(event.Account.String(), keygen.GenerateToken(32))
			if err != nil {
				log.WithError(err).Error("Failed to put data in db")
			}
		}
	}
}
