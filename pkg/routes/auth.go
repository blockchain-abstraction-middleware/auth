package routes

import (
	"net/http"

	db "github.com/blockchain-abstraction-middleware/auth/pkg/db"
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

	return router
}

func (r *AuthResource) auth() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		log.WithFields(log.Fields{"name": "auth"}).Info("Starting Auth Request")
		// Check if req.Header is truthy3
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
