package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

var configPath string

// IgnoredFolders is a struct to collect ignored folders.
type IgnoredFolders struct {
	Folders []string `json:"folders"`
}

// ConfigPath retrieves the main configuration path.
func ConfigPath() string {
	// Get configuration path
	usr, err := user.Current()
	if err != nil {
		log.Println("Error occurred while getting user path:", err)
		os.Exit(1)
	}
	return filepath.Join(usr.HomeDir, ".config", "go-inspect")
}

func init() {
	configPath = ConfigPath()
}

// LoadIgnoredFolders Loads the configuration file into the representive struct.
func LoadIgnoredFolders() (ignoredFolders *IgnoredFolders) {
	dat, err := ioutil.ReadFile(filepath.Join(configPath, "ignored_folders.json"))
	if err != nil {
		log.Println("Error occurred while loading configuration file:", err)
		os.Exit(1)
	}
	ignoredFolders = &IgnoredFolders{}
	err = json.Unmarshal(dat, &ignoredFolders)
	if err != nil {
		log.Println("Error occurred while loading config file:", err)
		os.Exit(1)
	}
	return ignoredFolders
}
