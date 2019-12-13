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
	// 	config *config.Config
}

// NewResource constructor func
func (r *AuthResource) NewResource(path string) *AuthResource {
	return &AuthResource{
		path: path,
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
	db := db.NewDB("./auth/db")

	router.Get("/auth", r.auth(db))

	log.WithFields(log.Fields{"name": "auth"}).Info("Created Auth Endpoint")

	return router
}

func (r *AuthResource) auth(db *db.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// Check if req.Header is truthy
		key, err := db.GetData(req.Header["Eth-Id"][0])
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
