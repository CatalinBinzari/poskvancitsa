package telegram

import tele "gopkg.in/telebot.v3"

func generateUI() {
	menu.Reply(
		menu.Row(menuBtnLovecoins, menuBtnSkvnon4),
	)

	activitiesSelector.Inline(
		activitiesSelector.Row(cumparaturiSectionBtn),
	)

	cumparaturiSelector.Inline(
		cumparaturiSelector.Row(cumparaturiShowMyBtn),
		cumparaturiSelector.Row(cumparaturiShowCommBtn),
		cumparaturiSelector.Row(cumparaturiAddBtn, cumparaturiRemBtn),
	)

	shopItemFocusSelector.Inline(
		shopItemFocusSelector.Row(minusShopItemBtn, plusShopItemBtn),
		shopItemFocusSelector.Row(modifyShopItemBtn),
		shopItemFocusSelector.Row(deleteShopItemBtn),
	)
}

func startCommand(c tele.Context) error {
	err := c.Send("Poskvon4imsea?", menu)
	if err != nil {
		return err
	}

	err = c.Send("Bun venit! 🥎👻", activitiesSelector)
	return err
}
