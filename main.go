package main

import (
	"log"
	"poskvancitsa/config"
	"poskvancitsa/storage/mongo"
	"poskvancitsa/telegram"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	cfg := config.MustLoad()

	storage := mongo.New(cfg.MongoConnectionString, 10*time.Second)

	log.Print("db cxn started")

	pref := tele.Settings{
		Token:  cfg.TgBotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	processor := telegram.New(b, storage)

	log.Print("telebot started")

	err = processor.Exec()
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Print("service started")

	processor.Bot.Start()
}
