package telegram

import tele "gopkg.in/telebot.v3"

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
