package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/boltdb/bolt"
)

// BUCKET is the main Bucket name for boltdb.
const BUCKET = "inspect"

func createDbConnection() *bolt.DB {
	db, err := bolt.Open(filepath.Join(ConfigPath(), "go_inspect.db"), 0600, nil)
	if err != nil {
		log.Println("Error while opening database: ", err)
		os.Exit(1)
	}
	return db
}

// InitDb initializes the database.
func InitDb() {
	db := createDbConnection()
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println("Error while closing db connection: ", err)
			os.Exit(1)
		}
	}()
	err := db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(BUCKET)); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Println("Error while creating the bucket: ", err)
		os.Exit(1)
	}
	log.Println("Database created.")
}

// SaveFiles saves a file in db.
func SaveFiles(files []string) error {
	db := createDbConnection()
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET))
		for _, f := range files {
			err := b.Put([]byte(f), []byte("false"))
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

// ChooseRandomFile chooses a random file from the db which is not marked as read.
func ChooseRandomFile() string {
	return ""
}

// CheckIfExists checks if a file is in the db.
func CheckIfExists(file string) (bool, error) {
	return false, nil
}

// MarkFileAsRead Marks a file as visited, which will never be retrieved again but is still retained.
func MarkFileAsRead(file string) error {
	return nil
}
