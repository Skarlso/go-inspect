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

type IgnoredFolders struct {
	Folders []string `json:"folders"`
}

// Path retrieves the main configuration path.
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

// LoadEC2Config Loads the EC2 configuration file into the representive struct.
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
