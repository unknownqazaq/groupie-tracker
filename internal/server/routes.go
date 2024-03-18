package server

import (
	"net/http"
)

// SetupRoutes настраивает маршруты для HTTP-сервера.
func SetupRoutes() http.Handler {
	mux := http.NewServeMux()
	// Определение обработчиков HTTP
	mux.HandleFunc("/artists", HandleArtists)
	mux.HandleFunc("/artistInfo", HandleArtistInfo)
	mux.HandleFunc("/", PageHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./internal/static"))))

	return mux
}
