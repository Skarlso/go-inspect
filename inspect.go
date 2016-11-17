package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/tucnak/telebot"
)

var ignoredFolders []string

func init() {
	ignoredFolders = LoadIgnoredFolders().Folders
	fmt.Println("Ignored folders:", ignoredFolders)
}

func main() {
	pathToWalk := os.Getenv("GO_INSPECT_PATH")
	if pathToWalk == "" {
		fmt.Println("Please set GO_INSPECT_PATH to the downloaded Go directory")
		os.Exit(1)
	}
	filepath.Walk("/Users/hannibal/golang/src/github.com/go", walkFun)
}

func walkFun(path string, info os.FileInfo, err error) error {
	if info.IsDir() || !strings.HasSuffix(path, ".go") {
		return nil
	}
	if isIgnoredFolder(path) {
		return nil
	}

	fmt.Println(path)
	return nil
}

func isIgnoredFolder(path string) bool {
	for _, f := range ignoredFolders {
		if strings.Contains(path, f) {
			return true
		}
	}
	return false
}

func sendMessage() {
	bot, err := telebot.NewBot(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}

	messages := make(chan telebot.Message, 100)
	go bot.Listen(messages, 1*time.Second)

	for {
		c := <-messages
		log.Println("Message was:", c.Text)
	}
	// for message := range messages {
	// 	if message.Text == "/hi" {
	// 		bot.SendMessage(message.Chat,
	// 			"Hello, "+message.Sender.FirstName+"!", nil)
	// 	}
	// }
}
