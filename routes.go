package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os/exec"

	"github.com/jinzhu/gorm"
)

// App holds the database connection
type App struct {
	db *gorm.DB
}

func loadJSONFromRequest(r *http.Request, target interface{}) error {
	if r.Body == nil {
		return fmt.Errorf("Body is empty")
	}
	if err := json.NewDecoder(r.Body).Decode(target); err != nil {
		return err
	}
	return nil
}

func (a *App) downloadTrack(rawURL string) (string, error) {
	cleanURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	viewIDs := cleanURL.Query().Get("v")
	if len(viewIDs) == 0 {
		return "", fmt.Errorf("Invalid url: %s", rawURL)
	}
	viewID := viewIDs
	filename := MusicFolder + "/" + string(viewID) + ".mp3"
	cmd := exec.Command("youtube-dl", "-x", "--audio-format", "mp3", "-o", filename, rawURL)
	err = cmd.Start()
	if err != nil {
		return "", err
	}
	log.Printf("Downloading...")
	err = cmd.Wait()
	return filename, err
}

func (a *App) createTrackHandler(w http.ResponseWriter, r *http.Request) {
	var t Track
	err := loadJSONFromRequest(r, &t)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	t.Upvotes = 0
	t.Downvotes = 0

	t.Filename, err = a.downloadTrack(t.YoutubeURL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if err := a.db.Create(&t).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *App) listTrackHandler(w http.ResponseWriter, r *http.Request) {
	db := a.db
	var tracks []Track
	if err := db.Find(&tracks).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&tracks)
}
