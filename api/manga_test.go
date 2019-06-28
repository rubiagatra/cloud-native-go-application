package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMangaToJSON(t *testing.T) {
	manga := Manga{ID: 1, Title: "Berserk", Author: "Kentarou Miura"}
	json := manga.ToJSON()

	assert.Equal(t, `{"id":1,"title":"Berserk","author":"Kentarou Miura"}`, string(json), "Manga JSON marshalling wrong.")
}

func TestMangaFromJSON(t *testing.T) {
	json := []byte(`{"id":1,"title":"Berserk","author":"Kentarou Miura"}`)
	manga := FromJSON(json)
	assert.Equal(t, Manga{ID: 1, Title: "Berserk", Author: "Kentarou Miura"}, manga, "Manga JSON unmarshalling wrong.")
}

func TestAllManga(t *testing.T) {
	listOfManga := AllManga()
	assert.Len(t, listOfManga, 2, "Wrong number of mangas.")
}

func TestCreateNewManga(t *testing.T) {
	manga := Manga{ID: 3, Title: "One Piece", Author: "Eiichiro Oda"}
	id, created := CreateManga(manga)
	assert.True(t, created, "Manga was not created.")
	assert.EqualValues(t, 3, id, "Wrong ID.")
}

func TestDoNotCreateExistingManga(t *testing.T) {
	manga := Manga{ID: 2}
	_, created := CreateManga(manga)
	assert.False(t, created, "Manga was created.")
}

func TestUpdateExistingManga(t *testing.T) {
	manga := Manga{Title: "Test Update", Author: "Me Again", ID: 1}
	updated := UpdateManga(1, manga)
	assert.True(t, updated, "Manga not updated.")

	manga, _ = GetManga(1)
	assert.Equal(t, "Test Update", manga.Title, "Title not updated.")
	assert.Equal(t, "Me Again", manga.Author, "Author not updated.")
}

func TestDeleteManga(t *testing.T) {
	DeleteManga(1)
	assert.Len(t, AllManga(), 2, "Wrong number of mangas after delete.")
}
