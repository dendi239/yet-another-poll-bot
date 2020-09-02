package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dendi239/yet-another-poll-bot/pkg/grammar"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	webhookURL := os.Getenv("WEBHOOK_URL")
	webhook := tgbotapi.NewWebhook(webhookURL + bot.Token)

	_, err = bot.Request(webhook)
	if err != nil {
		panic(err)
	}

	info, err := bot.GetWebhookInfo()

	if err != nil {
		panic(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("failed to set webhook: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe("/"+bot.Token, nil)

	for update := range updates {
		log.Printf("Found update: %v", update)
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Default text")
		words := strings.Fields(update.Message.Text)

		func() {
			defer func() {
				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}
			}()

			switch {
			case len(words) == 0:
				msg.Text = "No text found"
				return

			case words[0][:1] == "/":
				message := strings.Join(words[1:], " ")

				switch words[0] {
				case "/tokenize":
					t, err := grammar.Tokenize(message)
					if err != nil {
						msg.Text = fmt.Sprintf("tokenize error: %v", err)
						return
					}
					msg.Text = fmt.Sprintf("tokenize(%v) = %v", message, t)

				case "/tokenize_implication":
					t1, t2, err := grammar.TokenizeImplication(message)
					if err != nil {
						msg.Text = fmt.Sprintf("tokenize error: %v", err)
						return
					}
					msg.Text = fmt.Sprintf("%v => %v", t1, t2)

				case "/normalize":
					ts, err := grammar.Tokenize(message)
					if err != nil {
						msg.Text = fmt.Sprintf("tokenize error: %v", err)
						return
					}

					term, _, err := grammar.Parse(ts)
					if err != nil {
						msg.Text = fmt.Sprintf("parse error: %v", err)
						return
					}

					msg.Text = fmt.Sprintf("parsed: %v", term)

				default:
					return
				}
			}

			msg.ReplyToMessageID = update.Message.MessageID
		}()
	}
}
