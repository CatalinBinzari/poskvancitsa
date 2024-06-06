package telegram

import (
	"fmt"
	"poskvancitsa/storage"

	tele "gopkg.in/telebot.v3"
)

var (
	// Universal markup builders.
	menu                = &tele.ReplyMarkup{ResizeKeyboard: true}
	activitiesSelector  = &tele.ReplyMarkup{}
	cumparaturiSelector = &tele.ReplyMarkup{}

	// Reply buttons.
	menuBtnLovecoins = menu.Text("Lovecoinsss 💰")
	menuBtnSkvnon4   = menu.Text("Skvon4 😘🐈")

	// Inline buttons.
	//
	// Pressing it will cause the client to
	// send the bot a callback.
	//
	// Make sure Unique stays unique as per button kind
	// since it's required for callback routing to work.
	//

	cumparaturiSectionBtn = activitiesSelector.Data("Cumparaturi 🛒🛍️", "cumparaturiSection", "test")

	cumparaturiShowBtn = cumparaturiSelector.Data("📋 Arata lista", "cumparaturiShow", "test")
	cumparaturiAddBtn  = cumparaturiSelector.Data("✍️ Adauga", "cumparaturiAdd", "test")
	cumparaturiRemBtn  = cumparaturiSelector.Data("❌ Sterge", "cumparaturiRemove", "test keyword")
)

type userAction struct {
	userCommnd string
	userText   string
}

type Processor struct {
	Bot     *tele.Bot
	storage storage.Storage
}

func New(b *tele.Bot, storage storage.Storage) *Processor {
	return &Processor{
		Bot:     b,
		storage: storage,
	}
}

func (p *Processor) Exec() error {
	generateUI()
	userActionsMap := make(map[int64]userAction, 10)

	return p.handlers(userActionsMap)
}

func (p *Processor) handlers(userActionsMap map[int64]userAction) error {

	p.Bot.Handle("/start", func(c tele.Context) error {
		return startCommand(c)
	})

	p.Bot.Handle(tele.OnText, func(c tele.Context) error {
		// Incoming inline message.
		userID := c.Sender().ID
		action, ok := userActionsMap[userID]
		if !ok {
			c.Send("Nu imi e clar ce doresti, te rog alege actiunea prin apasarea unuia din butoane! 😜")
			return startCommand(c)
		}

		action.userText = c.Message().Text

		// fmt.Println("OnQ triggered, ", msg)
		resp := fmt.Sprintf("user %d, wants to %s %s\n", userID, action.userCommnd, action.userText)
		// add to database this entry
		return c.Send(resp)
	})

	p.Bot.Handle(&cumparaturiSectionBtn, func(c tele.Context) error {
		return c.Edit("poshopyatsa", cumparaturiSelector)
	})

	p.Bot.Handle(&cumparaturiShowBtn, func(c tele.Context) error {
		call := c.Callback()
		fmt.Printf("callback %+v", call)
		fmt.Printf("\n\nmsg payload %+v", call.Message)
		// switch_inline_query := []tele.InlineButton{{Text: "text", InlineQuery: ""}, {Text: "text2", InlineQuery: ""}}
		// data, ok := calllist[call.Data]
		_ = p.Bot.Respond(call, &tele.CallbackResponse{
			Text:      "callback of cumparaturiShowBtn",
			ShowAlert: false,
		})
		return nil
	})

	p.Bot.Handle(&cumparaturiAddBtn, func(c tele.Context) error {
		action := userAction{
			userCommnd: "addCumparaturi",
		}
		userActionsMap[c.Sender().ID] = action

		fmt.Println(userActionsMap)

		return c.Send("Ce vrei sa adaugi?", tele.ForceReply)
	})

	return nil
}

func startCommand(c tele.Context) error {
	err := c.Send("Poskvon4imsea?", menu)
	if err != nil {
		return err
	}

	err = c.Send("Bun venit! 🥎👻", activitiesSelector)
	return err
}

func generateUI() {
	menu.Reply(
		menu.Row(menuBtnLovecoins, menuBtnSkvnon4),
	)

	activitiesSelector.Inline(
		activitiesSelector.Row(cumparaturiSectionBtn),
	)

	cumparaturiSelector.Inline(
		cumparaturiSelector.Row(cumparaturiShowBtn),
		cumparaturiSelector.Row(cumparaturiAddBtn, cumparaturiRemBtn),
	)
}
