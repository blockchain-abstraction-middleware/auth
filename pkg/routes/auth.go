package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/blockchain-abstraction-middleware/rest-api/pkg/logger"
	"github.com/go-chi/chi"
	"github.com/syndtr/goleveldb/leveldb"
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

	db, err := leveldb.OpenFile("./auth/db", nil)

	if err != nil {
		fmt.Println(err)
	}

	_ = db.Put([]byte("0xb4FC6774a95A86134768d81A8eE19Cfbf5A171F6"), []byte("key12345!"), nil)
	_ = db.Put([]byte("0x7A7645f880248176998C6aB717b67ff60aAb106B"), []byte("key12345!"), nil)
	_ = db.Put([]byte("0x9290E64f25FEF981963566079e95fb2E6728Ab61"), []byte("key12345!"), nil)
	_ = db.Put([]byte("0xfe24C938Ca2EA10c56dc4c7AE81fA3AAe8854Bde"), []byte("key12345!"), nil)

	// defer db.Close()

	router.Get("/auth", r.auth(db))

	log.WithFields(log.Fields{"name": "auth"}).Info("Created Auth Endpoint")

	return router
}

func (r *AuthResource) auth(db *leveldb.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		// GET HEADERS

		// CHECK IN DB

		iter := db.NewIterator(nil, nil)
		for iter.Next() {
			// Remember that the contents of the returned slice should not be modified, and
			// only valid until the next call to Next.
			key := iter.Key()
			value := iter.Value()
			fmt.Println(string(key), string(value))
			log.Info(string(key), string(value))
		}
		iter.Release()
		// err = iter.Error()

		payload := struct {
			AddressList string `json:"addressList"`
			StatusCode  int    `json:"statusCode"`
		}{
			"gameJamAddressList[0].String()",
			200,
		}

		log.WithFields(log.Fields{"name": "auth"}).Info("Validated user: ?")

		json, _ := json.Marshal(payload)

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(json)
	}
}
