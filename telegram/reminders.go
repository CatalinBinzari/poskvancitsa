package telegram

import (
	"log/slog"
)

func BuyReminders() {
	slog.Info("BuyReminders")

	list, err := processor.storage.ShopItems()
	if err != nil {
		slog.Error("BuyReminders", "err", err)
		return
	}
	if len(list) == 0 {
		remindUsers("Poate e ziua perfecta pentru a adauga ceva pe lista de cumparaturi?")
	}

	remindUsers("Poate e ziua perfecta pentru a cumpara " + pickRandomItem(list).ItemName + " ?")
}
