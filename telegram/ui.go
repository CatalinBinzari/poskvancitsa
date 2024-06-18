package telegram

import tele "gopkg.in/telebot.v3"

var (
	menuBtnSkvnon4Str          = "Skvon4 😘🐈"
	menuBtnLovecoinsStr        = "Lovecoinsss 💰"
	menuCumparaturiShowCommStr = "👩🏻‍❤️‍👨🏻 Arata lista comuna"

	// Universal markup builders.
	menu                  = &tele.ReplyMarkup{ResizeKeyboard: true}
	activitiesSelector    = &tele.ReplyMarkup{}
	cumparaturiSelector   = &tele.ReplyMarkup{}
	shopItemFocusSelector = &tele.ReplyMarkup{}

	menuBtnLovecoins        = menu.Text(menuBtnLovecoinsStr)
	menuBtnSkvnon4          = menu.Text(menuBtnSkvnon4Str)
	menuCumparaturiShowComm = menu.Text(menuCumparaturiShowCommStr)

	cumparaturiSectionBtn = activitiesSelector.Data("Cumparaturi 🛒🛍️", "cumparaturiSection", "test")

	cumparaturiShowMyBtn   = cumparaturiSelector.Data("🙋🏻‍♂️ Arata lista mea", "cumparaturiShowMyBtn", "test")
	cumparaturiShowCommBtn = cumparaturiSelector.Data(menuCumparaturiShowCommStr, "cumparaturiShowCommBtn", "test")
	cumparaturiAddBtn      = cumparaturiSelector.Data("✍️ Adauga", "cumparaturiAdd", "test")
	cumparaturiRemBtn      = cumparaturiSelector.Data("❌ Sterge", "cumparaturiRemove", "test keyword")

	minusShopItemBtn  = shopItemFocusSelector.Data("➖", "minusShopItemBtn", "test")
	plusShopItemBtn   = shopItemFocusSelector.Data("➕", "plusShopItemBtn", "test")
	modifyShopItemBtn = shopItemFocusSelector.Data("⚙️ Modify", "modifyShopItemBtn", "test")
	deleteShopItemBtn = shopItemFocusSelector.Data("🚫 Delete", "deleteShopItemBtn", "test keyword")
)

func generateUI() {
	menu.Reply(
		menu.Row(menuCumparaturiShowComm),
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
