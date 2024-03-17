package groupie_tracker

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

const (
	artistsURL = "https://groupietrackers.herokuapp.com/api/artists"
)

func HandleArtists(w http.ResponseWriter, r *http.Request) {
	// Получаем данные об артистах из внешнего API
	resp, err := http.Get(artistsURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Декодируем ответ JSON в структуру ArtistResponse
	var artistData []struct {
		ID    int    `json:"id"`
		Image string `json:"image"`
		Name  string `json:"name"`
	}
	err = json.NewDecoder(resp.Body).Decode(&artistData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Кодируем данные в формат JSON и отправляем клиенту
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(artistData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleArtistInfo(w http.ResponseWriter, r *http.Request) {
	// Парсим ID артиста из URL запроса
	artistID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	// Получаем данные об артистах из внешнего API
	respArtists, err := http.Get(artistsURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer respArtists.Body.Close()

	// Получаем данные о связях из внешнего API
	artistIDStr := strconv.Itoa(artistID)
	url := "https://groupietrackers.herokuapp.com/api/relation/" + artistIDStr
	respRelation, err := http.Get(url)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer respRelation.Body.Close()

	// Декодируем данные в структуру Relation
	var relationData *Relation
	err = json.NewDecoder(respRelation.Body).Decode(&relationData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Декодируем ответы JSON
	var artistData []Artist
	err = json.NewDecoder(respArtists.Body).Decode(&artistData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Ищем артиста по ID
	var selectedArtist *Artist
	for _, artist := range artistData {
		if artist.ID == artistID {
			selectedArtist = &artist
			break
		}
	}
	// Проверяем, найден ли артист
	if selectedArtist == nil {
		http.Error(w, "Artist not found", http.StatusNotFound)
		return
	}

	// Отображаем HTML шаблон с данными об артисте и связях
	tmplPath := filepath.Join("artistInfo.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Создаем структуру, содержащую информацию об артисте и связях
	artistInfo := &ArtistInfo{
		Artist:   selectedArtist,
		Relation: relationData,
	}

	if err := tmpl.Execute(w, artistInfo); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var tmplPath string
		// Determine the template path based on the requested URL
		if r.URL.Path == "/" {
			tmplPath = filepath.Join("index.html")
		} else {
			fmt.Printf("Unknown URL path: %s", r.URL.Path)
			return
		}
		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			fmt.Printf("Error parsing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		if err := tmpl.Execute(w, nil); err != nil {
			fmt.Printf("Error executing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Status Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
