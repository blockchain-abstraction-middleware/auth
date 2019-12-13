package db

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

// DB will contrain the read/write db logic
type DB struct {
	leveldb *leveldb.DB
}

// NewDB returns a new instance of the db
func NewDB(path string) *DB {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		fmt.Println(err)
	}

	_ = db.Put([]byte("0xb4FC6774a95A86134768d81A8eE19Cfbf5A171F6"), []byte("1key12345"), nil)
	_ = db.Put([]byte("0x7A7645f880248176998C6aB717b67ff60aAb106B"), []byte("2key12345"), nil)
	_ = db.Put([]byte("0x9290E64f25FEF981963566079e95fb2E6728Ab61"), []byte("3key12345"), nil)
	_ = db.Put([]byte("0xfe24C938Ca2EA10c56dc4c7AE81fA3AAe8854Bde"), []byte("4key12345"), nil)

	// defer db.Close()

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
