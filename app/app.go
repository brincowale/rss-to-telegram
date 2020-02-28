package app

import (
	"errors"
	"github.com/getsentry/sentry-go"
	"github.com/mmcdole/gofeed"
	"rss-to-telegram/models"
	"rss-to-telegram/utils"
)

func GetPosts(config utils.Config) []*models.Post {
	var postList []*models.Post
	fp := gofeed.NewParser()
	for _, url := range config.ScrapeUrlList {
		feed, err := fp.ParseURL(url)
		if err != nil {
			sentry.CaptureException(err)
			continue
		}
		for _, item := range feed.Items {
			post := &models.Post{
				Title:       item.Title,
				Description: item.Description,
				URL:         item.Link,
				Categories:  item.Categories,
			}
			postList = append(postList, post)
		}
		if len(feed.Items) == 0 {
			sentry.CaptureException(errors.New("Zero posts scrapped from: " + url))
		}
	}
	return postList
}
