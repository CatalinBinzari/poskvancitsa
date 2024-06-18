package telegram

import (
	"fmt"
	"poskvancitsa/storage"
	"strconv"
	"strings"

	tele "gopkg.in/telebot.v3"
)

func handleStart(c tele.Context) error {
	return startCommand(c)
}

func handleCumparaturiSectionBtn(c tele.Context) error {
	return c.Edit("poshopyatsa", cumparaturiSelector)
}

func handleOntext(c tele.Context) error {
	userID := c.Sender().ID
	action, ok := userActionsMap[userID]
	if !ok {
		return unknownAction(c)
	}
	if action.userCommnd == "addCumparaturi" {
		action.userText = c.Message().Text

		shopitem := storage.AddShopItem{
			AddedBy:  strconv.FormatInt(userID, 10),
			Count:    1,
			ItemName: action.userText,
		}
		err := processor.storage.Save(&shopitem)
		if err != nil {
			return failedAction(c)
		}

		return c.Send("Adaugat! 🥳", &tele.SendOptions{
			ReplyTo: c.Message(),
		})
	} else if action.userCommnd == "modifyCumparaturi" {
		err := processor.storage.ModifyNameShopItem(action.userText, c.Message().Text)
		if err != nil {
			return failedAction(c)
		}
		return c.Send("Modificat! 🥳", &tele.SendOptions{
			ReplyTo: c.Message(),
		})
	} else {
		return unknownAction(c)
	}
}

func handleOnCallback(c tele.Context) error {
	userID := c.Sender().ID
	action, ok := userActionsMap[userID]
	if !ok {
		return unknownAction(c)
	}

	if action.userCommnd == "FocusCumparaturi" {
		call := c.Callback()
		fmt.Printf("%+v", call)
		if call.Data == "" {
			fmt.Println("Unknown callback data")
		}

		parts := strings.Split(call.Data, "||")
		action.userText = parts[0]
		action.userShopItemName = parts[1]
		userActionsMap[userID] = action
		return c.Send(userActionsMap[userID].userShopItemName,
			&tele.SendOptions{
				ReplyTo: c.Message(),
			},
			shopItemFocusSelector)
	}
	return nil
}

func handleCumparaturiShowCommBtn(c tele.Context) error {
	call := c.Callback()
	action := userAction{
		userCommnd: "FocusCumparaturi",
	}
	userActionsMap[c.Sender().ID] = action

	InlineShopButtonsList2 := make([][]tele.InlineButton, 10)
	for i := range InlineShopButtonsList2 {
		InlineShopButtonsList2[i] = make([]tele.InlineButton, 0, 10)
	}

	list, err := processor.storage.ShopItems()
	if err != nil {
		fmt.Println(err)
		_ = processor.Bot.Respond(call, &tele.CallbackResponse{
			Text:      "storage ERROR",
			ShowAlert: true,
		})
	}
	if len(list) == 0 {
		return processor.Bot.Respond(call, &tele.CallbackResponse{
			Text: "no items, please add",
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

	processor.Bot.Send(c.Sender(), "~🛒🛍️", &tele.ReplyMarkup{
		InlineKeyboard: InlineShopButtonsList2,
	})

	if err != nil {
		fmt.Println(err)
		_ = processor.Bot.Respond(call, &tele.CallbackResponse{
			Text:      "inline btn ERROR",
			ShowAlert: true,
		})
	}
	return processor.Bot.Respond(call, &tele.CallbackResponse{
		Text:      "Lista afisata ✨",
		ShowAlert: false,
	})
}

func handleCumparaturiAddBtn(c tele.Context) error {
	action := userAction{
		userCommnd: "addCumparaturi",
	}
	userActionsMap[c.Sender().ID] = action

	fmt.Println(userActionsMap)

	return c.Send("Ce vrei sa adaugi?", tele.ForceReply)
}

func handleMinusShopItemBtn(c tele.Context) error {
	action := userActionsMap[c.Sender().ID]
	if action.userText == "" {
		return failedAction(c)
	}

	err := processor.storage.ChangeShopItemCount(action.userText, -1)
	if err != nil {
		return failedAction(c)
	}

	return c.Send("-1 👍")
}

func handlePlusShopItemBtn(c tele.Context) error {
	action := userActionsMap[c.Sender().ID]
	if action.userText == "" {
		return failedAction(c)
	}

	err := processor.storage.ChangeShopItemCount(action.userText, 1)
	if err != nil {
		fmt.Println(err)
		return failedAction(c)
	}

	return c.Send("+1 👍")
}

func handleModifyShopItemBtn(c tele.Context) error {
	action := userActionsMap[c.Sender().ID]
	action.userCommnd = "modifyCumparaturi"
	userActionsMap[c.Sender().ID] = action

	return c.Send("Da un nume nou.", tele.ForceReply)
}

func handleDeleteShopItemBtn(c tele.Context) error {
	action := userActionsMap[c.Sender().ID]
	if action.userText == "" {
		return failedAction(c)
	}

	err := processor.storage.RemoveShopItem(action.userText)
	if err != nil {
		fmt.Println(err)
		return failedAction(c)
	}

	var strikethroughUserShopItemName string
	for _, r := range action.userShopItemName {
		strikethroughUserShopItemName += string(r) + strikethrough
	}

	return c.Send(fmt.Sprintf("%s 😵", strikethroughUserShopItemName))
}
