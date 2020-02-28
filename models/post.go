package models

type Post struct {
	Title       string   `gorm:"-"`
	Description string   `gorm:"-"`
	URL         string   `gorm:"column:url"`
	Categories  []string `gorm:"-"`
}

func (Post) TableName() string {
	return "blog_url_news"
}
