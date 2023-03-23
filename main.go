package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

//(1) Users will be able to see a list of artists and albums
//(2) Admins will be able to add, remove or update information for albums and artists.
//Table: albums
//id (integer, primary key)
//title (string)
//artist_id (integer, foreign key referencing artists(id))
//genre (string)
//release_date (date)

// Album Struct
type Album struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Artist   *Artist `json:"artist"`
	Language string  `json:"language"`
}

// Artist Struct
type Artist struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Init a albums var as a slice of Album struct
var albums []Album
var artists []Artist

func initArtistDB() {
	var newArtist Artist
	newArtist = Artist{
		ID:        1,
		FirstName: "Mon",
		LastName:  "Mojaiya",
	}

	artists = append(artists, newArtist)

	newArtist = Artist{
		ID:        2,
		FirstName: "Ekdin",
		LastName:  "Matir",
	}

	artists = append(artists, newArtist)

	newArtist = Artist{
		ID:        3,
		FirstName: "Arijit",
		LastName:  "Singh",
	}

	artists = append(artists, newArtist)

}

func initAlbumDB() {
	var newAlbum Album
	newAlbum = Album{
		ID:       1,
		Title:    "Amar mon mojaiya re",
		Artist:   &artists[0],
		Language: "Bengali",
	}
	albums = append(albums, newAlbum)

	newAlbum = Album{
		ID:       2,
		Title:    "Ekdin Matir Vitore Hobe Ghor",
		Artist:   &artists[1],
		Language: "Bengali",
	}
	albums = append(albums, newAlbum)

	newAlbum = Album{
		ID:       3,
		Title:    "Arijit Singh er Gaaaan",
		Artist:   &artists[2],
		Language: "Hindi",
	}
	albums = append(albums, newAlbum)
}

// Get All Albums
func getAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(albums)
}

// Get Single Album
func getAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get the params
	//Loop through albums and find that id
	paramsID, _ := strconv.Atoi(params["id"])
	for _, curAlbum := range albums {
		if curAlbum.ID == paramsID {
			json.NewEncoder(w).Encode(curAlbum)
		}
	}
	json.NewEncoder(w).Encode(Album{})

}

// Create a new Album
func createAlbum(w http.ResponseWriter, r *http.Request) {

}

func updateAlbum(w http.ResponseWriter, r *http.Request) {

}
func deleteAlbum(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//Init Router
	router := mux.NewRouter()

	//Giving Initial Data
	initArtistDB()
	initAlbumDB()

	//Route Handlers /Endpoints
	router.HandleFunc("/api/albums", getAlbums).Methods("GET")
	router.HandleFunc("/api/albums/{id}", getAlbum).Methods("GET")
	router.HandleFunc("/api/albums", createAlbum).Methods("POST")
	router.HandleFunc("/api/albums/{id}", updateAlbum).Methods("PUT")
	router.HandleFunc("/api/albums/{id}", deleteAlbum).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5050", router))

}
