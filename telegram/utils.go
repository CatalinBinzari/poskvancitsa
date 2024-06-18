package telegram

import (
	"fmt"
	"poskvancitsa/storage"

	tele "gopkg.in/telebot.v3"
)

func startCommand(c tele.Context) error {
	err := c.Send("Poskvon4imsea?", menu)
	if err != nil {
		return err
	}

	err = c.Send("Bun venit! 🥎👻", activitiesSelector)
	return err
}

func unknownAction(c tele.Context) error {
	c.Send("Nu imi e clar ce doresti, te rog alege actiunea prin apasarea unuia din butoane! 😜")
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
		fmt.Println(err)
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
		fmt.Println(err)
		return failedAction(c)
	}
	return err
}
