package main

import (
	"log"
	"os"
	"time"

	"github.com/tucnak/telebot"
)

func main() {
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
