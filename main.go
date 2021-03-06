package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPasswd := os.Getenv("POSTGRES_PWD")
	dbStr := fmt.Sprintf("host=localhost user=%s dbname=radio_recast sslmode=disable password=%s", dbUser, dbPasswd)
	db, err := gorm.Open("postgres", dbStr)
	if err != nil {
		panic(err)
	}
	app := App{db}
	defer db.Close()
	db.AutoMigrate(&Track{})

	r := mux.NewRouter()
	r.HandleFunc("/tracks", app.listTrackHandler).Methods("GET")
	r.HandleFunc("/tracks", app.createTrackHandler).Methods("POST")

	// Static file server
	s := http.StripPrefix("/music/", http.FileServer(http.Dir("./music/")))
	r.PathPrefix("/music/").Handler(s)

	r.HandleFunc("/", app.index).Methods("GET")
	http.ListenAndServe(":4242", r)
}
