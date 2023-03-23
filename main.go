//API Design

//Imagine you have been tasked with designing the backend for a music streaming website like Spotify.
//In this website,
//(1) Users will be able to see a list of artists and albums
//(2) Admins will be able to add, remove or update information for albums and artists.

//Route Handlers /Endpoints
//router.HandleFunc("/api/albums", getAlbums).Methods("GET")
//router.HandleFunc("/api/albums/{id}", getAlbum).Methods("GET")
//router.HandleFunc("/api/albums", createAlbum).Methods("POST")
//router.HandleFunc("/api/albums/{id}", updateAlbum).Methods("PUT")
//router.HandleFunc("/api/albums/{id}", deleteAlbum).Methods("DELETE")
//
//router.HandleFunc("/api/artists", getArtists).Methods("GET")
//router.HandleFunc("/api/artists/{id}", getArtist).Methods("GET")

//todo
//"Improvements"
//Use map in database for efficiency (better time complexity)

package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sort"
	"strconv"
	//"fmt"
	//"time"
	//
	//"github.com/go-chi/chi/v5"
	//"github.com/go-chi/chi/v5/middleware"
	//"github.com/go-chi/jwtauth/v5"
	//"github.com/lestrrat-go/jwx/jwa"
	//"github.com/lestrrat-go/jwx/jwt"
)

// Album Struct
type Album struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Artist   Artist `json:"artist"`
	Language string `json:"language"`
}

// Artist Struct
type Artist struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// For new ID generation
var maxAlbumID int

// Init a albums var as a slice of Album struct
var albums []Album

func initMaxID() {
	maxAlbumID = 0
}

func initAlbumDB() {
	var newAlbum Album
	maxAlbumID++
	newAlbum = Album{
		ID:    maxAlbumID,
		Title: "Amar mon mojaiya re",
		Artist: Artist{
			FirstName: "MonMojaiya",
			LastName:  "artist",
		},
		Language: "Bengali",
	}
	albums = append(albums, newAlbum)

	maxAlbumID++
	newAlbum = Album{
		ID:    maxAlbumID,
		Title: "Ekdin Matir Vitore Hobe Ghor",
		Artist: Artist{
			FirstName: "EkdinMatir",
			LastName:  "artist",
		},
		Language: "Bengali",
	}
	albums = append(albums, newAlbum)

	maxAlbumID++
	newAlbum = Album{
		ID:    maxAlbumID,
		Title: "Arijit Singh er Ekti Gaaaan",
		Artist: Artist{
			FirstName: "Arijit",
			LastName:  "Singh",
		},
		Language: "Hindi",
	}
	albums = append(albums, newAlbum)

}

//{
//"title": "Dupli Tumi Bondhu Kala Pakhi Gaaaan",
//"artist": {
//"firstname": "KalaPakhi",
//"lastname": "Artist"
//},
//"language": "Bangla"
//}

func initDB() {
	initMaxID()
	initAlbumDB()
}

// Get All Albums
func getAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sort.SliceStable(albums, func(i, j int) bool {
		return albums[i].ID < albums[j].ID
	})
	json.NewEncoder(w).Encode(albums)
}

// Get Single Album
func getAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get the params

	paramsID, _ := strconv.Atoi(params["id"])
	//Loop through albums and find that id
	for _, curAlbum := range albums {
		if curAlbum.ID == paramsID {
			json.NewEncoder(w).Encode(curAlbum)
			return
		}
	}

	json.NewEncoder(w).Encode("No Album data available for the given id")

}

// Create a new Album
func createAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newAlbum Album
	err := json.NewDecoder(r.Body).Decode(&newAlbum)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	maxAlbumID++
	newAlbum.ID = maxAlbumID
	albums = append(albums, newAlbum)
	json.NewEncoder(w).Encode(newAlbum)
}

func updateAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get the params
	paramsID, _ := strconv.Atoi(params["id"])

	var newAlbum Album
	err := json.NewDecoder(r.Body).Decode(&newAlbum)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	//Loop through albums and find that id
	for index, curAlbum := range albums {
		if curAlbum.ID == paramsID {
			//time complexity can be improved by using map
			albums = append(albums[:index], albums[index+1:]...)
			newAlbum.ID = paramsID

			albums = append(albums, newAlbum)

			break
		}
	}
	sort.SliceStable(albums, func(i, j int) bool {
		return albums[i].ID < albums[j].ID
	})
	json.NewEncoder(w).Encode(albums)
}
func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get the params
	paramsID, _ := strconv.Atoi(params["id"])

	//Loop through albums and find that id
	for index, curAlbum := range albums {
		if curAlbum.ID == paramsID {
			//time complexity can be improved by using map
			albums = append(albums[:index], albums[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(albums)
}

// Get All Artist
func getArtists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var uniqueArtist = make(map[Artist]bool)
	var uniqueArtistList []Artist
	for _, curAlbum := range albums {
		curArtist := curAlbum.Artist
		_, keyExists := uniqueArtist[curArtist]
		if keyExists == false {
			uniqueArtist[curArtist] = true
			uniqueArtistList = append(uniqueArtistList, curArtist)
		}
	}
	json.NewEncoder(w).Encode(uniqueArtistList)
}

// Get Top X Artist by album total number of album count
func getTopArtists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get the params

	numberOfArtistToShow, _ := strconv.Atoi(params["x"])

	var artistCount = make(map[Artist]int)

	for _, curAlbum := range albums {
		artistCount[curAlbum.Artist]++
	}

	type ArtistsAlbumCount struct {
		Artist   Artist
		AlbumCnt int
	}
	var artistList []ArtistsAlbumCount

	for artist, x := range artistCount {
		artistList = append(artistList, ArtistsAlbumCount{
			artist,
			x,
		})
	}

	sort.SliceStable(artistList, func(i, j int) bool {
		return artistList[i].AlbumCnt > artistList[j].AlbumCnt
	})

	maxLen := len(artistList)
	if numberOfArtistToShow < maxLen {
		maxLen = numberOfArtistToShow
	}

	artistList = artistList[:maxLen]
	json.NewEncoder(w).Encode(artistList)
}

func main() {
	//Init Router
	router := mux.NewRouter()

	//Giving Initial Data
	initDB()

	//Route Handlers /Endpoints
	router.HandleFunc("/api/albums", getAlbums).Methods("GET")
	router.HandleFunc("/api/albums/{id}", getAlbum).Methods("GET")
	router.HandleFunc("/api/albums", createAlbum).Methods("POST")
	router.HandleFunc("/api/albums/{id}", updateAlbum).Methods("PUT")
	router.HandleFunc("/api/albums/{id}", deleteAlbum).Methods("DELETE")

	router.HandleFunc("/api/artists", getArtists).Methods("GET")
	router.HandleFunc("/api/artists/{x}", getTopArtists).Methods("GET")

	log.Fatal(http.ListenAndServe(":5050", router))

}
