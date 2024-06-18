package telegram

import (
	"poskvancitsa/storage"

	tele "gopkg.in/telebot.v3"
)

const strikethrough = "\u0336" // Combining Long Stroke Overlay

var (
	// Universal markup builders.
	menu                  = &tele.ReplyMarkup{ResizeKeyboard: true}
	activitiesSelector    = &tele.ReplyMarkup{}
	cumparaturiSelector   = &tele.ReplyMarkup{}
	shopItemFocusSelector = &tele.ReplyMarkup{}

	menuBtnLovecoins = menu.Text("Lovecoinsss 💰")
	menuBtnSkvnon4   = menu.Text("Skvon4 😘🐈")

	cumparaturiSectionBtn = activitiesSelector.Data("Cumparaturi 🛒🛍️", "cumparaturiSection", "test")

	cumparaturiShowMyBtn   = cumparaturiSelector.Data("🙋🏻‍♂️ Arata lista mea", "cumparaturiShowMyBtn", "test")
	cumparaturiShowCommBtn = cumparaturiSelector.Data("👩🏻‍❤️‍👨🏻 Arata lista comuna", "cumparaturiShowCommBtn", "test")
	cumparaturiAddBtn      = cumparaturiSelector.Data("✍️ Adauga", "cumparaturiAdd", "test")
	cumparaturiRemBtn      = cumparaturiSelector.Data("❌ Sterge", "cumparaturiRemove", "test keyword")

	minusShopItemBtn  = shopItemFocusSelector.Data("➖", "minusShopItemBtn", "test")
	plusShopItemBtn   = shopItemFocusSelector.Data("➕", "plusShopItemBtn", "test")
	modifyShopItemBtn = shopItemFocusSelector.Data("⚙️ Modify", "modifyShopItemBtn", "test")
	deleteShopItemBtn = shopItemFocusSelector.Data("🚫 Delete", "deleteShopItemBtn", "test keyword")
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
	p.Bot.Handle(&cumparaturiSectionBtn, func(c tele.Context) error {
		return c.Edit("poshopyatsa", cumparaturiSelector)
	})

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
