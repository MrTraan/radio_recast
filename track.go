package main

import (
	"errors"
	"net/url"

	"github.com/jinzhu/gorm"
)

// MusicFolder linter break my balls
const MusicFolder = "./music"

// Track lol
type Track struct {
	gorm.Model
	YoutubeURL string `gorm:"column:youtube_url" json:"youtube_url"`
	Title      string `gorm:"column:title" json:"title"`
	Upvotes    uint64 `gorm:"column:upvotes" json:"upvotes"`
	Downvotes  uint64 `gorm:"column:downvotes" json:"downvotes"`
	Filename   string `gorm:"column:filename" json:"filename"`
}

// BeforeSave stuff
func (t *Track) BeforeSave() error {
	if len(t.YoutubeURL) == 0 {
		return errors.New("Empty Youtube Url")
	}

	if len(t.Title) == 0 {
		return errors.New("Empty Title")
	}

	if len(t.Filename) == 0 {
		url, err := url.Parse(t.YoutubeURL)
		if err != nil {
			return err
		}
		viewIDs := url.Query().Get("v")
		if len(viewIDs) == 0 {
			return errors.New("Invalid youtube url")
		}
		viewID := viewIDs[0]
		filename := MusicFolder + "/" + string(viewID) + ".mp3"

		t.Filename = filename
	}

	return nil
}
