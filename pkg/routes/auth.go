package routes

import (
	"encoding/json"
	"net/http"

	db "github.com/blockchain-abstraction-middleware/auth/pkg/db"
	verify "github.com/blockchain-abstraction-middleware/auth/pkg/verify"
	log "github.com/blockchain-abstraction-middleware/rest-api/pkg/logger"
	"github.com/go-chi/chi"
)

// Auth Resource implements server.Resource interface
type AuthResource struct {
	path string
	db   *db.DB
	// 	config *config.Config
}

// NewResource constructor func
func (r *AuthResource) NewResource(path string, db *db.DB) *AuthResource {
	return &AuthResource{
		path: path,
		db:   db,
		// config: config,
	}
}

// Path returns the Resource base path
func (r *AuthResource) Path() string {
	return r.path
}

// Routes bootstraps health routes
func (r *AuthResource) Routes() http.Handler {
	router := chi.NewRouter()

	router.Get("/auth", r.auth())
	log.WithFields(log.Fields{"name": "auth"}).Info("Created Auth Endpoint")

	router.Post("/serve-api-key", r.apiKey())
	log.WithFields(log.Fields{"name": "api-key"}).Info("Created API-KEY Endpoint")

	return router
}

func (r *AuthResource) auth() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		log.WithFields(log.Fields{"name": "auth"}).Info("Starting Auth Request")
		// TODO: Check if req.Header is truthy
		key, err := r.db.GetData(req.Header["Eth-Id"][0])
		if err != nil {
			res.WriteHeader(http.StatusForbidden)

			return
		}

		if key != req.Header["Bam-Key"][0] {
			res.WriteHeader(http.StatusForbidden)

			return
		}

		log.WithFields(log.Fields{"name": req.Header["Eth-Id"][0]}).Info("Validated user")
		res.WriteHeader(http.StatusOK)
	}
}

// GetAPIKeyPayload payload definition for getting an apikey
type GetAPIKeyPayload struct {
	HexEthAddress string `json:"hexEthAddress"`
	HexHash       string `json:"hexHash"`
	HexSignature  string `json:"hexSignature"`
}

func (r *AuthResource) apiKey() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		log.WithFields(log.Fields{"name": "auth"}).Info("Starting apikey Request")
		getAPIKeyPayload := new(GetAPIKeyPayload)
		json.NewDecoder(req.Body).Decode(&getAPIKeyPayload)

		pubAddress := verify.ParseAddressFromSignedMessage(getAPIKeyPayload.HexHash, getAPIKeyPayload.HexSignature)

		key, err := r.db.GetData(pubAddress.String())
		if err != nil {
			res.WriteHeader(http.StatusForbidden)

			return
		}

		payload := struct {
			APIKey     string `json:"apiKey"`
			StatusCode int    `json:"statusCode"`
		}{
			key,
			http.StatusOK,
		}

		json, _ := json.Marshal(payload)

		res.Header().Set("Content-Type", "application/json")
		res.Write(json)

		log.WithFields(log.Fields{"api user: ": getAPIKeyPayload.HexEthAddress}).Info("User collected api-key")
	}
}
