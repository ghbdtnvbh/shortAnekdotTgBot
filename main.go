package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/anaskhan96/soup"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func randomAnekdot() string {
	url := "https://baneks.site/random"
	resp, _ := soup.Get(url)
	site := soup.HTMLParse(resp)
	block := site.Find("section", "itemprop", "description").FindAll("p")

	var result string

	for _, v := range block {
		p := regexp.MustCompile(`<*.p>`)
		br := regexp.MustCompile(`<br/>`)
		qute := regexp.MustCompile(`&#\d+?;`)
		replaceTagP := p.ReplaceAllString(v.HTML(), "")
		replaceTagBr := br.ReplaceAllString(replaceTagP, "\n")
		result = qute.ReplaceAllString(replaceTagBr, `"`)

	}
	return result
}

func shortAnekdot() string {

	for {

		text := randomAnekdot()
		if len([]rune(text)) < 140 {
			return text

		}
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI("2002027839:AAFqwfk3z1-0RmCTtkvAIdE_1MEHHvkTgcs")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch strings.ToLower(update.Message.Text) {
		case "анекдот", "anekdot":
			msg.Text = shortAnekdot()
		default:
			msg.Text = "I don't know that command\nSend анекдот or anekdot"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
