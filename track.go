package main

import "github.com/jinzhu/gorm"

type Track struct {
	gorm.Model
	YoutubeURL string `gorm:"column:youtube_url"`
	Title      string `gorm:"column:title"`
	Upvotes    uint64 `gorm:"column:upvotes"`
	Downvotes  uint64 `gorm:"column:downvotes"`
	Filename   string `gorm:"column:filname"`
}
