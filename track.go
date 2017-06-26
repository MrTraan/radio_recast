package main

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Track struct {
	gorm.Model
	YoutubeURL string `gorm:"column:youtube_url"`
	Title      string `gorm:"column:title"`
	Upvotes    uint64 `gorm:"column:upvotes"`
	Downvotes  uint64 `gorm:"column:downvotes"`
	Filename   string `gorm:"column:filename"`
}

func (t *Track) BeforeSave() error {
	if len(t.YoutubeURL) == 0 {
		return errors.New("Empty Youtube Url")
	}

	if len(t.Title) == 0 {
		return errors.New("Empty Title")
	}

	if len(t.Filename) == 0 {
		return errors.New("Empty Title")
	}

	return nil
}
