package main

import (
	"github.com/getsentry/sentry-go"
	"rss-to-telegram/app"
	"rss-to-telegram/utils"
	"time"
)

func main() {
	configs := utils.ReadConfig()
	_ = sentry.Init(sentry.ClientOptions{
		Dsn: configs.Sentry,
	})
	utils.CreateConnectionDB(configs.DBConnection)
	for _, post := range app.GetPosts(configs) {
		if utils.IsNewPost(post) {
			message := utils.CreateMessage(post)
			if utils.SendTelegramMessage(message, configs, post) {
				time.Sleep(2 * time.Second)
				utils.InsertPost(post)
			}
		}
	}
}
