package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/boltdb/bolt"
)

// BUCKET is the main Bucket name for boltdb.
const BUCKET = "inspect"

// InitDb initializes the database.
func InitDb() {
	db, err := bolt.Open(filepath.Join(ConfigPath(), "go_inspect.db"), 0600, nil)
	if err != nil {
		fmt.Println("Error while opening database: ", err)
		os.Exit(1)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(BUCKET)); err != nil {
			return err
		}
		return nil
	})

	log.Println("Database created.")
}
