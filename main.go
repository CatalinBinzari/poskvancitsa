package main

import (
	"log"
	"log/slog"
	"poskvancitsa/config"
	"poskvancitsa/cronjob"
	"poskvancitsa/storage/mongo"
	"poskvancitsa/telegram"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	cfg := config.MustLoad()

	storage := mongo.New(cfg.MongoConnectionString, 10*time.Second)

	slog.Info("db connextion started", "db", "mongodb")

	pref := tele.Settings{
		Token:  cfg.TgBotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	processor := telegram.New(b, storage)

	slog.Info("telebot started")

	err = processor.Exec()
	if err != nil {
		log.Fatal(err)
		return
	}

	cronjob.StartCronjob(telegram.BuyReminders, "17:57")

	slog.Info("service started")

	processor.Bot.Start()
}
