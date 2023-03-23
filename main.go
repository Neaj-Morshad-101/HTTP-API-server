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
//Use map in database for efficiency in time complexity

package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sort"
	"strconv"
)

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

// ArtistInfo Struct
type ArtistInfo struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// For new ID generation
var maxArtistID int
var maxAlbumID int

// Init a albums var as a slice of Album struct
var albums []Album
var artists []Artist

func initMaxID() {
	maxArtistID = 0
	maxAlbumID = 0
}

func initArtistDB() {
	var newArtist Artist
	maxArtistID++
	newArtist = Artist{
		ID:        maxArtistID,
		FirstName: "MonMojaiya",
		LastName:  "artist",
	}
	artists = append(artists, newArtist)

	maxArtistID++
	newArtist = Artist{
		ID:        maxArtistID,
		FirstName: "EkdinMatir",
		LastName:  "artist",
	}
	artists = append(artists, newArtist)

	maxArtistID++
	newArtist = Artist{
		ID:        maxArtistID,
		FirstName: "Arijit",
		LastName:  "Singh",
	}
	artists = append(artists, newArtist)
}

func initAlbumDB() {
	var newAlbum Album
	maxAlbumID++
	newAlbum = Album{
		ID:       maxAlbumID,
		Title:    "Amar mon mojaiya re",
		Artist:   &artists[0],
		Language: "Bengali",
	}
	albums = append(albums, newAlbum)

	maxAlbumID++
	newAlbum = Album{
		ID:       maxAlbumID,
		Title:    "Ekdin Matir Vitore Hobe Ghor",
		Artist:   &artists[1],
		Language: "Bengali",
	}
	albums = append(albums, newAlbum)

	maxAlbumID++
	newAlbum = Album{
		ID:       maxAlbumID,
		Title:    "Arijit Singh er Ekti Gaaaan",
		Artist:   &artists[2],
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
	initArtistDB()
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
	maxArtistID++
	newAlbum.Artist.ID = maxArtistID
	albums = append(albums, newAlbum)
	artists = append(artists, *newAlbum.Artist)
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

			for idx, curArtist := range artists {
				if curArtist.ID == curAlbum.Artist.ID {
					artists = append(artists[:idx], artists[idx+1:]...)
					break
				}
			}

			maxArtistID++
			newAlbum.Artist.ID = maxArtistID
			albums = append(albums, newAlbum)
			artists = append(artists, *newAlbum.Artist)
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
			for idx, curArtist := range artists {
				if curArtist.ID == curAlbum.Artist.ID {
					artists = append(artists[:idx], artists[idx+1:]...)
					break
				}
			}
			break
		}
	}

	json.NewEncoder(w).Encode(albums)
}

// Get All Artist with duplicates
//func getArtistsAll(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(artists)
//}

func getArtists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var uniqueArtist = make(map[ArtistInfo]int)
	var uniqueArtistList []Artist
	for _, curArtist := range artists {
		var newArtistInfo ArtistInfo
		newArtistInfo.FirstName = curArtist.FirstName
		newArtistInfo.LastName = curArtist.LastName

		_, keyExists := uniqueArtist[newArtistInfo]
		if keyExists == false {
			uniqueArtist[newArtistInfo] = curArtist.ID
			uniqueArtistList = append(uniqueArtistList, curArtist)
		}
	}
	// Not necessary, because it will remain sorted everytime, but recheck it
	//sort.SliceStable(uniqueArtistList, func(i, j int) bool {
	//	return uniqueArtistList[i].ID < uniqueArtistList[j].ID
	//})
	json.NewEncoder(w).Encode(uniqueArtistList)
}

// Get Single Artist
func getArtist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // Get the params

	paramsID, _ := strconv.Atoi(params["id"])
	//Loop through artist and find that id
	for _, curArtist := range artists {
		if curArtist.ID == paramsID {
			json.NewEncoder(w).Encode(curArtist)
			return
		}
	}

	json.NewEncoder(w).Encode("No Artist data available for the given id")
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
	router.HandleFunc("/api/artists/{id}", getArtist).Methods("GET")

	log.Fatal(http.ListenAndServe(":5050", router))

}
