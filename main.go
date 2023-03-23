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
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwa"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
)

var jwtkey = []byte("Neaj's Secret Key, He will not share it")
var tokenAuth *jwtauth.JWTAuth
var tokenString string
var token jwt.Token

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

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
var credslist map[string]string

func initCreds() {
	tokenAuth = jwtauth.New(string(jwa.HS256), jwtkey, nil)
	credslist = make(map[string]string)
	creds := credentials{
		"Neaj Morshad",
		"1234",
	}
	credslist[creds.Username] = creds.Password
}

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

func initAll() {
	initCreds()
	initMaxID()
	initAlbumDB()
}

// Get All Albums
func getAlbums(w http.ResponseWriter, r *http.Request) {
	log.Println("Called: getAlbums()")
	w.Header().Set("Content-Type", "application/json")
	sort.SliceStable(albums, func(i, j int) bool {
		return albums[i].ID < albums[j].ID
	})
	json.NewEncoder(w).Encode(albums)
}

// Get Single Album
func getAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := chi.URLParam(r, "id")
	paramsID, _ := strconv.Atoi(param)
	//
	//params := mux.Vars(r) // Get the params
	//
	//paramsID, _ := strconv.Atoi(params["id"])

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
	param := chi.URLParam(r, "id")
	paramsID, _ := strconv.Atoi(param)
	//params := mux.Vars(r) // Get the params
	//paramsID, _ := strconv.Atoi(params["id"])

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
			sort.SliceStable(albums, func(i, j int) bool {
				return albums[i].ID < albums[j].ID
			})
			json.NewEncoder(w).Encode(albums)

			return
		}
	}

	json.NewEncoder(w).Encode("No album data available for the given id")

}
func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r) // Get the params

	param := chi.URLParam(r, "id")
	paramsID, _ := strconv.Atoi(param)
	fmt.Println("------------------>paramid", paramsID)

	//Loop through albums and find that id
	for index, curAlbum := range albums {
		if curAlbum.ID == paramsID {
			fmt.Println("----------------->hi")
			//time complexity can be improved by using map
			albums = append(albums[:index], albums[index+1:]...)
			break
		}
	}
	fmt.Println(albums)

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
	param := chi.URLParam(r, "cnt")
	//paramsID, _ := strconv.Atoi(param)
	//params := mux.Vars(r) // Get the params

	numberOfArtistToShow, _ := strconv.Atoi(param)

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

func Login(w http.ResponseWriter, r *http.Request) {

	var creds credentials

	err := json.NewDecoder(r.Body).Decode(&creds)

	fmt.Println(creds)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	correctPassword, ok := credslist[creds.Username]

	if !ok || creds.Password != correctPassword {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//tokenAuth = jwtauth.New(string(jwa.HS256), jwtkey, nil)

	expiretime := time.Now().Add(2 * time.Minute)

	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{
		"aud": "Neaj Morshad",
		"exp": expiretime.Unix(),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Value:   tokenString,
		Expires: expiretime,
	})
	w.WriteHeader(http.StatusOK)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Expires: time.Now(),
	})
	w.WriteHeader(http.StatusOK)
}

func main() {
	fmt.Println("started")

	//Giving Initial Data
	initAll()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Post("/login", Login)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		//	r.Use(middleware.BasicAuth("user", credslist))

		r.Route("/albums", func(r chi.Router) {
			r.Get("/", getAlbums)
			r.Get("/{id}", getAlbum)
			r.Post("/", createAlbum)
			r.Delete("/{id}", deleteAlbum)
			r.Put("/{id}", updateAlbum)
		})

		r.Route("/artists", func(r chi.Router) {
			r.Get("/", getArtists)
			r.Get("/{cnt}", getTopArtists)
		})
		r.Post("/logout", Logout)

	})
	http.ListenAndServe(":5050", r)

	//previous code using max
	//router := max.NewRouter()
	////Route Handlers /Endpoints
	//router.HandleFunc("/albums", getAlbums).Methods("GET")
	//router.HandleFunc("/albums/{id}", getAlbum).Methods("GET")
	//router.HandleFunc("/albums", createAlbum).Methods("POST")
	//router.HandleFunc("/albums/{id}", updateAlbum).Methods("PUT")
	//router.HandleFunc("/albums/{id}", deleteAlbum).Methods("DELETE")
	//
	//router.HandleFunc("/artists", getArtists).Methods("GET")
	//router.HandleFunc("/artists/{x}", getTopArtists).Methods("GET")

}
