package telegram

import (
	"poskvancitsa/storage"

	tele "gopkg.in/telebot.v3"
)

const strikethrough = "\u0336" // Combining Long Stroke Overlay

type userAction struct {
	userCommnd       string
	userText         string
	userShopItemName string
}

type Processor struct {
	Bot     *tele.Bot
	storage storage.Storage
}

var processor *Processor
var userActionsMap map[int64]userAction

func New(b *tele.Bot, storage storage.Storage) *Processor {
	processor = &Processor{
		Bot:     b,
		storage: storage,
	}

	return processor
}

func (p *Processor) Exec() error {
	generateUI()
	userActionsMap = make(map[int64]userAction, 10)

	return p.handlers()
}

func (p *Processor) handlers() error {

	p.Bot.Handle("/start", handleStart)
	p.Bot.Handle(&cumparaturiSectionBtn, handleCumparaturiSectionBtn)

	p.Bot.Handle(tele.OnText, handleOntext)
	p.Bot.Handle(tele.OnCallback, handleOnCallback)
	p.Bot.Handle(&cumparaturiShowCommBtn, handleCumparaturiShowCommBtn)
	p.Bot.Handle(&cumparaturiAddBtn, handleCumparaturiAddBtn)
	p.Bot.Handle(&minusShopItemBtn, handleMinusShopItemBtn)
	p.Bot.Handle(&plusShopItemBtn, handlePlusShopItemBtn)
	p.Bot.Handle(&modifyShopItemBtn, handleModifyShopItemBtn)
	p.Bot.Handle(&deleteShopItemBtn, handleDeleteShopItemBtn)

	return nil
}
