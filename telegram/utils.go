package telegram

import (
	"fmt"
	"log/slog"
	"math/rand"
	"poskvancitsa/storage"
	"time"

	tele "gopkg.in/telebot.v3"
)

func startCommand(c tele.Context) error {
	err := c.Send("Bun venit", menu)
	if err != nil {
		return err
	}

	return c.Send("Poskvon4imsea? 🥎👻", activitiesSelector)
}

func unknownAction(c tele.Context) error {
	c.Send("Nu imi e clar ce doresti, te rog alege actiunea prin apasarea unuia din butoane! 😜")
	return startCommand(c)
}

func todoAction(c tele.Context) error {
	c.Send("Will be done soon! 😜")
	return startCommand(c)
}

func failedAction(c tele.Context) error {
	return c.Send("Operatiunea nu a reusit. Incercati mai tarziu. 🥲", &tele.SendOptions{
		ReplyTo: c.Message(),
	})
}

func showCumparaturi(c tele.Context, cumparaturi_type int) error {
	action := userAction{
		userCommnd: "FocusCumparaturi",
	}
	userActionsMap[c.Sender().ID] = action

	InlineShopButtonsList2 := make([][]tele.InlineButton, 10)
	for i := range InlineShopButtonsList2 {
		InlineShopButtonsList2[i] = make([]tele.InlineButton, 0, 10)
	}

	var list []storage.ShopItem
	var err error
	if cumparaturi_type == COMMON_CUMPARATURI {
		list, err = processor.storage.ShopItems()
	} else if cumparaturi_type == MY_CUMPARATURI {
		list, err = processor.storage.ShopItems(c.Sender().ID)
	} else {
		return failedAction(c)
	}
	if err != nil {
		slog.Error("showCumparaturi", "err", err)
		return failedAction(c)
	}
	if len(list) == 0 {
		return c.Send("no items, please add. 🥲", &tele.SendOptions{
			ReplyTo: c.Message(),
		})
	}

	numnerOfElementsPerRow := 2
	row := 0
	for idx, element := range list {
		if idx%numnerOfElementsPerRow == 0 {
			row += 1
		}
		payload := fmt.Sprintf("%s||%s", element.ID, element.ItemName)
		buttonText := fmt.Sprintf("%d🧮 %s", element.Count, element.ItemName)

		btn := tele.InlineButton{Text: buttonText, Data: payload}
		InlineShopButtonsList2[row] = append(InlineShopButtonsList2[row], btn)
	}

	_, err = processor.Bot.Send(c.Sender(), "~🛒🛍️", &tele.ReplyMarkup{
		InlineKeyboard: InlineShopButtonsList2,
	})
	if err != nil {
		slog.Error("showCumparaturi", "err", err)
		return failedAction(c)
	}
	return err
}

func notifyUsers(c tele.Context, msg string) {
	for _, id := range UserIdList {
		if id == c.Sender().ID { // do not notify yourself
			continue
		}
		var user tele.User
		user.ID = id

		processor.Bot.Send(&user, msg)
	}
}

func remindUsers(msg string) {
	for _, id := range UserIdList {
		var user tele.User
		user.ID = id

		_, err := processor.Bot.Send(&user, msg)
		if err != nil {
			slog.Error("remindUsers", "user", user.ID, "err", err)
		}
	}
}

func skvon4Users(c tele.Context, msg string) {

	skvon4er := ""

	for name, id := range UserIdList {
		if id == c.Sender().ID {
			skvon4er = name
			break
		}
	}
	for _, id := range UserIdList {
		if id == c.Sender().ID { // do not notify yourself
			continue
		}
		var user tele.User
		user.ID = id

		processor.Bot.Send(&user, skvon4er+msg)
	}
}

func pickRandomItem(list []storage.ShopItem) storage.ShopItem {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random index within the bounds of the list
	randomIndex := rand.Intn(len(list))

	// Return the item at the random index
	return list[randomIndex]
}
