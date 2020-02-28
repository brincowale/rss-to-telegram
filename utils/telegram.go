package utils

import (
	"github.com/getsentry/sentry-go"
	"github.com/go-errors/errors"
	"github.com/grokify/html-strip-tags-go"
	"github.com/parnurzeal/gorequest"
	"github.com/tidwall/gjson"
	"net/http"
	"rss-to-telegram/models"
	"strings"
	"time"
)

var request *gorequest.SuperAgent

func init() {
	request = gorequest.New().
		Set("Content-Type", "application/json").
		Timeout(30*time.Second).
		Retry(3, 5*time.Second, http.StatusInternalServerError)
}

type Message struct {
	ChatId string `json:"chat_id"`
	Text   string `json:"text"`
}

func SendTelegramMessage(message string, config Config, post *models.Post) bool {
	var URL = "https://api.telegram.org/bot" + config.TelegramApiKey + "/sendMessage"
	data := Message{
		ChatId: config.TelegramChannel,
		Text:   message,
	}
	_, body, errList := request.Post(URL).Send(data).End()
	if errList != nil {
		for _, err := range errList {
			sentry.CaptureException(err)
		}
	}
	if gjson.Get(body, "ok").Bool() {
		return true
	}
	sentry.CaptureException(errors.New("Cannot send to Telegram: " + post.URL))
	return false
}

func CreateMessage(post *models.Post) string {
	categories := strings.Join(post.Categories, ", ")
	returnLine := "\n"
	str :=
		strip.StripTags(post.Title) + returnLine + returnLine +
			strip.StripTags(post.Description) + returnLine + returnLine +
			"Categories: " + strip.StripTags(categories) + returnLine + returnLine +
			post.URL + returnLine
	return str
}
