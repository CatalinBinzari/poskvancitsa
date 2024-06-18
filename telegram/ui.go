package telegram

import tele "gopkg.in/telebot.v3"

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
