package utils

import (
	"github.com/getsentry/sentry-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"rss-to-telegram/models"
)

var DB *gorm.DB
var err error

func CreateConnectionDB(database string) {
	DB, err = gorm.Open("mysql", database)
	if err != nil {
		sentry.CaptureException(err)
	}
}

func IsNewPost(post *models.Post) bool {
	var p models.Post
	DB.Where(&models.Post{URL: post.URL}).First(&p)
	return p.URL == ""
}

func InsertPost(post *models.Post) {
	DB.Create(&post)
}
