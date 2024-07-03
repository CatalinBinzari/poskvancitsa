package telegram

import (
	"poskvancitsa/storage"

	tele "gopkg.in/telebot.v3"
)

const strikethrough = "\u0336" // Combining Long Stroke Overlay

const (
	COMMON_CUMPARATURI = 0
	MY_CUMPARATURI     = 1
)

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

var UserIdList = map[string]int64{
	"Iunona":  570730943,
	"Catalin": 6583361128,
}

const (
	ADD_CUMPARATURI = " a fost adaugat pe lista de cumparaturi."
	DEL_CUMPARATURI = " a fost sters de pe lista de cumparaturi."
)

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
	p.Bot.Handle(&cumparaturiShowMyBtn, handleCumparaturiShowMyBtn)
	p.Bot.Handle(&cumparaturiAddBtn, handleCumparaturiAddBtn)
	p.Bot.Handle(&cumparaturiRemBtn, handleCumparaturiRemBtn)
	p.Bot.Handle(&minusShopItemBtn, handleMinusShopItemBtn)
	p.Bot.Handle(&plusShopItemBtn, handlePlusShopItemBtn)
	p.Bot.Handle(&modifyShopItemBtn, handleModifyShopItemBtn)
	p.Bot.Handle(&deleteShopItemBtn, handleDeleteShopItemBtn)

	return nil
}
