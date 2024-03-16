package groupie_tracker

import (
	"encoding/json"
	"log"
	"net/http"
)

// API представляет собой структуру с URL-адресами API.
type API struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

// New обрабатывает запросы к /api и возвращает URL-адреса API.
func New(w http.ResponseWriter, r *http.Request) {
	// Создание объекта API с URL-адресами.
	endpoints := API{
		Artists:   "https://groupietrackers.herokuapp.com/api/artists",
		Locations: "https://groupietrackers.herokuapp.com/api/locations",
		Dates:     "https://groupietrackers.herokuapp.com/api/dates",
		Relation:  "https://groupietrackers.herokuapp.com/api/relation",
	}

	// Кодирование структуры API в JSON.
	data, err := json.Marshal(endpoints)
	if err != nil {
		log.Println("Ошибка кодирования JSON:", err)
		http.Error(w, "Ошибка кодирования JSON", http.StatusInternalServerError)
		return
	}

	// Установка заголовков ответа и запись JSON-ответа клиенту.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		log.Println("Ошибка записи ответа:", err)
		http.Error(w, "Ошибка записи ответа", http.StatusInternalServerError)
		return
	}
}
