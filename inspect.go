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

func main() {
	filepath.Walk(".", walkFun)
}

func walkFun(path string, info os.FileInfo, err error) error {
	if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
		fmt.Println(path)
		return nil
	}
	return nil
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
