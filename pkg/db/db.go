package db

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

// DB will contain the read/write db logic
type DB struct {
	leveldb *leveldb.DB
}

// NewDB returns a new instance of the db
func NewDB(path string) *DB {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		fmt.Println(err)
	}

	_ = db.Put([]byte("0x8c6253A7dCCE198b4385f17f390bC6fcE34A19Ea"), []byte("4key12345"), nil)

	return &DB{
		leveldb: db,
	}
}

// GetData returns data is a given key exists
func (db *DB) GetData(key string) (string, error) {
	value := []byte("")
	value, err := db.leveldb.Get([]byte(key), nil)

	return string(value), err
}

// PutData puts key, value data into the database
func (db *DB) PutData(key string, value string) error {
	err := db.leveldb.Put([]byte(key), []byte(value), nil)

	return err
}
