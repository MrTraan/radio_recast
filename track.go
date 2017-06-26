package main

import "github.com/jinzhu/gorm"

type Track struct {
	gorm.Model
	YoutubeURL string
	Title      string
	Upvotes    uint64
	Downvotes  uint64
	Filename   string
}
