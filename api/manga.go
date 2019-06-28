package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Manga type with ID, Name, and Author
type Manga struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var ListOfManga = map[int64]Manga{
	1: Manga{ID: 1, Title: "Berserk", Author: "Kentarou Miura"},
	2: Manga{ID: 2, Title: "Full Metal Alchemist", Author: "Hiromu Arakawa"},
}

// ToJSON to be used for marshalling of Manga type
func (m Manga) ToJSON() []byte {
	ToJSON, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

// FromJSON to be used for unmarshalling of Manga type
func FromJSON(data []byte) Manga {
	manga := Manga{}
	err := json.Unmarshal(data, &manga)
	if err != nil {
		panic(err)
	}
	return manga
}

// AllManga returns a slice of all books
func AllManga() []Manga {
	values := make([]Manga, len(ListOfManga))
	idx := 0
	for _, manga := range ListOfManga {
		values[idx] = manga
		idx++
	}
	return values
}

// ListOfMangaHandleFunc to be used as http.HandleFunc for Manga API
func ListOfMangaHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		listOfManga := AllManga()
		writeJSON(w, listOfManga)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		manga := FromJSON(body)
		id, created := CreateManga(manga)
		if created {
			w.Header().Add("Location", "/api/manga/"+strconv.FormatInt(id, 10))
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

// MangaHandleFunc to be used as http.HandleFunc for Manga API
func MangaHandleFunc(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(r.URL.Path[len("/api/manga/"):], 10, 64)

	switch method := r.Method; method {
	case http.MethodGet:
		manga, found := GetManga(id)
		if found {
			writeJSON(w, manga)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		manga := FromJSON(body)
		exists := UpdateManga(id, manga)
		if exists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		DeleteManga(id)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

func writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

// GetManga returns the Manga for a given ID
func GetManga(id int64) (Manga, bool) {
	manga, found := ListOfManga[id]
	return manga, found
}

// CreateManga creates a new Manga if it does not exist
func CreateManga(manga Manga) (int64, bool) {
	_, exists := ListOfManga[manga.ID]
	if exists {
		return 0, false
	}
	ListOfManga[manga.ID] = manga
	return manga.ID, true
}

// UpdateManga updates an existing manga
func UpdateManga(id int64, manga Manga) bool {
	_, exists := ListOfManga[id]
	if exists {
		ListOfManga[id] = manga
	}
	return exists
}

// DeleteManga removes a book from the map by ID key
func DeleteManga(id int64) {
	delete(ListOfManga, id)
}
