package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	if err := a.db.Create(&t).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *App) listTrackHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
