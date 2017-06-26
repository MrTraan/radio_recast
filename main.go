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
	defer db.Close()
	db.AutoMigrate(&Track{})

	r := mux.NewRouter()
	r.HandleFunc("/tracks", listTrackHandler).Methods("GET")
	r.HandleFunc("/tracks", createTrackHandler).Methods("POST")
	// r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)

}
