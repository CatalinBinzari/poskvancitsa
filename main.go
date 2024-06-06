package main

import (
	"log"
	"poskvancitsa/config"
	"poskvancitsa/storage/mongo"
	"poskvancitsa/telegram"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {

	cfg := config.MustLoad()
	//storage := files.New(storagePath)

	storage := mongo.New(cfg.MongoConnectionString, 10*time.Second)

	log.Print("db started")

	pref := tele.Settings{
		Token:  cfg.TgBotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	processor := telegram.New(b, storage)

	log.Print("telebot started")

	err = processor.Exec()
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Print("service started")

	processor.Bot.Start()
}

// r := b.NewMarkup()

// // Reply buttons:
// r.Text("Hello!")
// r.Contact("Send phone number")
// r.Location("Send location")
// r.Poll(tele.PollQuiz)

// // Inline buttons:
// r.Data("Show help", "help") // data is optional
// r.Data("Delete item", "delete", item.ID)
// r.URL("Visit", "https://google.com")
// r.Query("Search", query)
// r.QueryChat("Share", query)
// r.Login("Login", &tele.Login{...})

// b.Handle("/cumparaturi", func(c tele.Context) error {
// 	return c.Send("Lista de cumparaturi. Alege o actiune 👇", cumparaturiSelector)
// })

// b.Handle(tele.OnText, func(c tele.Context) error {
// 	// All the text messages that weren't
// 	// captured by existing handlers.

// 	var (
// 		user = c.Sender()
// 		text = c.Text()
// 	)

// 	fmt.Printf("CATAL context %+v, %+v\n", user, text)

// 	// Use full-fledged bot's functions
// 	// only if you need a result:
// 	msg, err := b.Send(user, text)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Printf("CATAL message: %+v\n", msg)

// 	// Instead, prefer a context short-hand:
// 	return c.Send(text)
// })

// b.Handle(tele.OnChannelPost, func(c tele.Context) error {
// 	// Channel posts only.
// 	msg := c.Message()
// })

// b.Handle(tele.OnPhoto, func(c tele.Context) error {
// 	// Photos only.
// 	photo := c.Message().Photo
// 	fmt.Println(proto)
// })

// b.Handle(tele.OnCallback, func(cont tele.Context) error {
// 	fmt.Println("on callback called")
// 	call := cont.Callback()
// 	// fmt.Printf("callback %+v", call)
// 	// switch_inline_query := []tele.InlineButton{{Text: "text", InlineQuery: ""}, {Text: "text2", InlineQuery: ""}}
// 	// data, ok := calllist[call.Data]
// 	_ = b.Respond(call, &tele.CallbackResponse{
// 		Text:      "produs sters",
// 		ShowAlert: false,
// 	})
// 	// _, err = b.EditReplyMarkup(call, &tele.ReplyMarkup{
// 	// 	InlineKeyboard: [][]tele.InlineButton{switch_inline_query},
// 	// })
// 	// var err error
// 	// if !ok {
// 	// 	_ = b.Respond(call, &tb.CallbackResponse{
// 	// 		Text:      "报价失效咯~ 请重新发起查询",
// 	// 		ShowAlert: true,
// 	// 	})
// 	// 	_, err = b.EditReplyMarkup(call, &tb.ReplyMarkup{
// 	// 		InlineKeyboard: [][]tb.InlineButton{switch_inline_query},
// 	// 	})
// 	// 	return err
// 	// }
// 	return nil
// })
// b.Handle("/tags", func(c tele.Context) error {
// 	tags := c.Args() // list of arguments splitted by a space
// 	for i, tag := range tags {
// 		// iterate through passed arguments
// 		fmt.Printf("%+v, %+v", i, tag)
// 	}

// 	a := &tele.Audio{File: tele.FromDisk("simple-acoustic-folk-138360.mp3")}

// 	fmt.Println(a.OnDisk())  // true
// 	fmt.Println(a.InCloud()) // false
// 	fmt.Println(a.FileID)

// 	// Will upload the file from disk and send it to the recipient
// 	msg, _ := b.Send(c.Recipient(), a)
// 	// p := &tele.Photo{File: tele.FromDisk("chicken.jpg")}
// 	// v := &tele.Video{File: tele.FromURL("https://www.youtube.com/watch?v=0CKrh9g7YCE&ab_channel=Noukash")}
// 	// v := &tele.Photo{File: tele.FromDisk("chicken.jpg")}

// 	b.Send(c.Recipient(), "message", &tele.SendOptions{
// 		ReplyTo:    msg,
// 		Protected:  true,
// 		HasSpoiler: true,
// 	})
// 	// b.Send(c.Recipient(), v)
// 	return nil
// 	// return c.Send("got it", msgs, err)
// })

// b.Handle(tele.OnText, func(m *tele.Message) error {
// 	return m.Send(m.Sender, "hello world")
// }

// b.Handle(tele.OnCallback, func(c *tele.Callback) {
// 	log.Println("got inline " + c.Data)

// 	//switch classid

// 	//b.Edit(

// 	b.Respond(c, &tele.CallbackResponse{Text: "clicked " + c.Data})
// })
// b.Handle(&cumparaturiRemBtn, func(c *tele.Callback) {
// 	// on inline button pressed (callback!)

// 	// always respond!
// 	c.Respond(&tele.CallbackResponse{"test"})
// })

// func main() {
// 	pref := tele.Settings{
// 		// Token:  os.Getenv("TOKEN"),
// 		Token:  "7396824811:AAHwkaF5l0Wr7qY6iwdr3b6bgm-KlpxNdgA",
// 		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
// 	}

// 	b, err := tele.NewBot(pref)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	// b, _ := tele.NewBot(tele.Settings{})

// 	// This button will be displayed in user's
// 	// reply keyboard.
// 	replyBtn := tele.ReplyButton{Text: "🌕 Button #1"}
// 	replyKeys := [][]tele.ReplyButton{
// 		[]tele.ReplyButton{replyBtn},
// 		// ...
// 	}

// 	// And this one — just under the message itself.
// 	// Pressing it will cause the client to send
// 	// the bot a callback.
// 	//
// 	// Make sure Unique stays unique as it has to be
// 	// for callback routing to work.
// 	inlineBtn := tele.InlineButton{
// 		Unique: "sad_moon",
// 		Text:   "🌚 Button #2",
// 	}
// 	inlineKeys := [][]tele.InlineButton{
// 		[]tele.InlineButton{inlineBtn},
// 		// ...
// 	}

// 	b.Handle(&replyBtn, func(m *tele.Message) {
// 		// on reply button pressed
// 	})

// 	b.Handle(&inlineBtn, func(c *tele.Callback) {
// 		// on inline button pressed (callback!)

// 		// always respond!
// 		c.Respond(&tele.CallbackResponse{})
// 	})

// 	// Command: /start <PAYLOAD>
// 	b.Handle("/start", func(m *tele.Message) {
// 		if !m.Private() {
// 			return
// 		}

// 		b.Send(m.Sender, "Hello!", &tele.ReplyMarkup{
// 			ReplyKeyboard:  replyKeys,
// 			InlineKeyboard: inlineKeys,
// 		})
// 	})

// 	b.Start()
// }
