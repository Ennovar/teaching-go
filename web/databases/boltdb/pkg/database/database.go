// Package database handles all communication between software and the database
package database

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/boltdb/bolt"
)

// Database constant
const (
	DB_MAIN = "datastore.db"
)

// Everytime you add a new bucket constant you should append to
// the Open function in the Update part that makes sure all top
// level buckets exist.
const (
	BUCKET_TEST = "test"
)

// Error codes
var (
	ErrKeyNotExist = errors.New("key does not exist")
)

type Datastore struct {
	handle *bolt.DB
}

// Open function will open the database and return a Datastore struct
// that has a handle within it for various datastore functions. When you
// call this function make sure to call it with the directory + database.DB_MAIN
// like so: database.Open("my/dir/path/" + database.DB_MAIN)
func Open(filepath string) (*Datastore, error) {
	db, err := bolt.Open(filepath, 0666, &bolt.Options{Timeout: 15 * time.Second})

	if err != nil {
		return nil, err
	}

	ds := &Datastore{
		handle: db,
	}

	// Ensure that all top-level buckets exist (bucket constants)
	err = ds.handle.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(BUCKET_TEST))

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return ds, nil
}

// Close is a function attached to the Datastore struct that will close the current bolt instance
func (ds *Datastore) Close() error {
	return ds.handle.Close()
}

// Get is a function attached to the Datastore struct that will get data from the current bolt instance
func (ds *Datastore) Get(bucket string, key []byte, result interface{}) error {
	return ds.handle.View(func(tx *bolt.Tx) error {
		dsValue := tx.Bucket([]byte(bucket)).Get(key)

		if dsValue == nil {
			return ErrKeyNotExist
		}

		return json.Unmarshal(dsValue, result)
	})
}

// Put is a function attached to the Datastore struct that will put data into the current bolt instance
func (ds *Datastore) Put(bucket string, key []byte, value interface{}) error {
	var err error
	dsValue, ok := value.([]byte)

	if !ok {
		dsValue, err = json.Marshal(value)

		if err != nil {
			return err
		}

	}

	return ds.handle.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(bucket)).Put(key, dsValue)
	})
}

// Delete is a function attached to the Datastore struct that will delete data from the current bolt instance
func (ds *Datastore) Delete(bucket string, key []byte) error {
	return ds.handle.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(bucket)).Delete(key)
	})
}