package main

import (
	groupie_tracker "groupie-tracker"
	"log"
	"net/http"
)

func main() {

	// Определение обработчиков HTTP
	http.HandleFunc("/artists", groupie_tracker.HandleArtists)
	http.HandleFunc("/artistInfo", groupie_tracker.HandleArtistInfo)
	http.HandleFunc("/", groupie_tracker.PageHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	srv := new(groupie_tracker.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
