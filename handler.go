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
	artistsURL  = "https://groupietrackers.herokuapp.com/api/artists/"
	relationURL = "https://groupietrackers.herokuapp.com/api/relation/"
)

func HandleArtists(w http.ResponseWriter, r *http.Request) {
	// Получаем данные об артистах из внешнего API
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
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

func fetchArtistData(artistID int) (*Artist, error) {
	artistIDStr := strconv.Itoa(artistID)
	resp, err := http.Get(artistsURL + artistIDStr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artistData *Artist
	if err := json.NewDecoder(resp.Body).Decode(&artistData); err != nil {
		return nil, err
	}

	return artistData, nil
}

func fetchRelationData(artistID int) (*Relation, error) {
	artistIDStr := strconv.Itoa(artistID)
	resp, err := http.Get(relationURL + artistIDStr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var relationData *Relation
	if err := json.NewDecoder(resp.Body).Decode(&relationData); err != nil {
		return nil, err
	}

	return relationData, nil
}

func HandleArtistInfo(w http.ResponseWriter, r *http.Request) {
	// Парсим ID артиста из URL запроса
	artistID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	// Получаем данные об артисте
	artistData, err := fetchArtistData(artistID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Получаем данные о связях
	relationData, err := fetchRelationData(artistID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		Artist:   artistData,
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
