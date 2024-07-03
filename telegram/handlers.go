package telegram

import (
	"fmt"
	"log/slog"
	"poskvancitsa/storage"
	"strconv"
	"strings"

	tele "gopkg.in/telebot.v3"
)

func handleStart(c tele.Context) error {
	slog.Info("handleStart", "user", c.Sender().ID)
	err := startCommand(c)
	if err != nil {
		slog.Error("handleStart", "user", c.Sender().ID, "err", err)
	}
	return err
}

func handleCumparaturiSectionBtn(c tele.Context) error {
	slog.Info("handleCumparaturiSectionBtn", "user", c.Sender().ID)
	err := c.Edit("poshopyatsa", cumparaturiSelector)
	if err != nil {
		slog.Error("handleCumparaturiSectionBtn", "user", c.Sender().ID, "err", err)
	}
	return err
}

func handleOntext(c tele.Context) error {
	slog.Info("handleOntext", "user", c.Sender().ID, "message", c.Message().Text)
	switch c.Text() {
	case menuCumparaturiShowCommStr:
		return handleCumparaturiShowCommBtn(c)
	case menuBtnSkvnon4Str:
		return handleBtnSkvnon4Str(c)
	case menuBtnLovecoinsStr:
		return todoAction(c)
	}

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
			slog.Error("handleOntext", "user", c.Sender().ID, "err", err)
			return failedAction(c)
		}

		c.Send("Adaugat! 🥳", &tele.SendOptions{
			ReplyTo: c.Message(),
		})
		notifyUsers(c, shopitem.ItemName+ADD_CUMPARATURI)
		return nil
	} else if action.userCommnd == "modifyCumparaturi" {
		err := processor.storage.ModifyNameShopItem(action.userText, c.Message().Text)
		if err != nil {
			slog.Error("handleOntext", "user", c.Sender().ID, "err", err)
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
	slog.Info("handleOnCallback", "user", c.Sender().ID)
	userID := c.Sender().ID
	action, ok := userActionsMap[userID]
	if !ok {
		return unknownAction(c)
	}
	if action.userCommnd == "" {
		return unknownAction(c)
	}

	call := c.Callback()
	if call.Data == "" {
		slog.Error("handleOnCallback Unknown callback data", "user", c.Sender().ID)
	}

	parts := strings.Split(call.Data, "||")
	if len(parts) != 2 {
		return unknownAction(c)
	}
	action.userText = parts[0]
	action.userShopItemName = parts[1]
	userActionsMap[userID] = action

	switch action.userCommnd {
	case "FocusCumparaturi":
		return c.Send(userActionsMap[userID].userShopItemName,
			&tele.SendOptions{
				ReplyTo: c.Message(),
			},
			shopItemFocusSelector)

	case "RemCumparaturi":
		return handleDeleteShopItemBtn(c)
	}
	return nil
}

func handleCumparaturiShowCommBtn(c tele.Context) error {
	slog.Info("handleCumparaturiShowCommBtn", "user", c.Sender().ID)
	return showCumparaturi(c, COMMON_CUMPARATURI)
}

func handleCumparaturiShowMyBtn(c tele.Context) error {
	slog.Info("handleCumparaturiShowMyBtn", "user", c.Sender().ID)
	return showCumparaturi(c, MY_CUMPARATURI)
}

func handleCumparaturiAddBtn(c tele.Context) error {
	slog.Info("handleCumparaturiAddBtn", "user", c.Sender().ID)
	action := userAction{
		userCommnd: "addCumparaturi",
	}
	userActionsMap[c.Sender().ID] = action

	return c.Send("Ce vrei sa adaugi?", tele.ForceReply)
}

func handleModifyShopItemBtn(c tele.Context) error {
	slog.Info("handleModifyShopItemBtn", "user", c.Sender().ID)
	action, ok := userActionsMap[c.Sender().ID]
	if !ok {
		return failedAction(c)
	}
	action.userCommnd = "modifyCumparaturi"
	userActionsMap[c.Sender().ID] = action

	return c.Send("Da un nume nou:", tele.ForceReply)
}

func handleDeleteShopItemBtn(c tele.Context) error {
	slog.Info("handleDeleteShopItemBtn", "user", c.Sender().ID)
	action, ok := userActionsMap[c.Sender().ID]
	if !ok {
		return failedAction(c)
	}
	if action.userText == "" {
		return failedAction(c)
	}

	err := processor.storage.RemoveShopItem(action.userText)
	if err != nil {
		slog.Error("handleDeleteShopItemBtn", "user", c.Sender().ID, "err", err)
		return failedAction(c)
	}

	var strikethroughUserShopItemName string
	for _, r := range action.userShopItemName {
		strikethroughUserShopItemName += string(r) + strikethrough
	}

	err = c.Send(fmt.Sprintf("%s 😵", strikethroughUserShopItemName))
	notifyUsers(c, action.userShopItemName+DEL_CUMPARATURI)
	return err
}

func handleMinusShopItemBtn(c tele.Context) error {
	slog.Info("handleMinusShopItemBtn", "user", c.Sender().ID)
	action := userActionsMap[c.Sender().ID]
	if action.userText == "" {
		return failedAction(c)
	}

	err := processor.storage.ChangeShopItemCount(action.userText, -1)
	if err != nil {
		slog.Error("handleMinusShopItemBtn", "user", c.Sender().ID, "err", err)
		return failedAction(c)
	}

	return c.Send("-1 👍")
}

func handlePlusShopItemBtn(c tele.Context) error {
	slog.Info("handlePlusShopItemBtn", "user", c.Sender().ID)
	action := userActionsMap[c.Sender().ID]
	if action.userText == "" {
		return failedAction(c)
	}

	err := processor.storage.ChangeShopItemCount(action.userText, 1)
	if err != nil {
		slog.Error("handlePlusShopItemBtn", "user", c.Sender().ID, "err", err)
		return failedAction(c)
	}

	return c.Send("+1 👍")
}

func handleCumparaturiRemBtn(c tele.Context) error {
	slog.Info("handleCumparaturiRemBtn", "user", c.Sender().ID)
	showCumparaturi(c, COMMON_CUMPARATURI)
	action := userAction{
		userCommnd: "RemCumparaturi",
	}
	userActionsMap[c.Sender().ID] = action
	return c.Send("Alege ce doresti sa stergi.")
}

func handleBtnSkvnon4Str(c tele.Context) error {
	slog.Info("handleBtnSkvnon4Str", "user", c.Sender().ID)
	skvon4Users(c, " te skvon4este 😘🐈")
	// todo stats, cate skvon4 ai trimis si cate ai primit
	return c.Send("Skvon4 😘🐈 trimis")
}
