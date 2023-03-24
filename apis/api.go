package apis

import (
	"encoding/json"
	"fmt"
	"github.com/Neaj-Morshad-101/HTTP-API-server/data"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
)

func initAll() {
	data.InitCreds()
	data.InitMaxID()
	data.InitAlbumDB()
}

// GetAlbums Get All Albums
func GetAlbums(w http.ResponseWriter, r *http.Request) {
	log.Println("Called: getAlbums()")
	w.Header().Set("Content-Type", "application/json")
	sort.SliceStable(data.Albums, func(i, j int) bool {
		return data.Albums[i].ID < data.Albums[j].ID
	})
	json.NewEncoder(w).Encode(data.Albums)
}

// GetAlbum Get Single Album
func GetAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := chi.URLParam(r, "id")
	paramsID, _ := strconv.Atoi(param)
	//
	//params := mux.Vars(r) // Get the params
	//
	//paramsID, _ := strconv.Atoi(params["id"])

	//Loop through data.Albums and find that id
	for _, curAlbum := range data.Albums {
		if curAlbum.ID == paramsID {
			json.NewEncoder(w).Encode(curAlbum)
			return
		}
	}

	json.NewEncoder(w).Encode("No Album data available for the given id")

}

// CreateAlbum Create a new Album
func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newAlbum data.Album
	err := json.NewDecoder(r.Body).Decode(&newAlbum)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	data.MaxAlbumID++
	newAlbum.ID = data.MaxAlbumID
	data.Albums = append(data.Albums, newAlbum)
	json.NewEncoder(w).Encode(newAlbum)
}

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := chi.URLParam(r, "id")
	paramsID, _ := strconv.Atoi(param)
	//params := mux.Vars(r) // Get the params
	//paramsID, _ := strconv.Atoi(params["id"])

	var newAlbum data.Album
	err := json.NewDecoder(r.Body).Decode(&newAlbum)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	//Loop through data.Albums and find that id
	for index, curAlbum := range data.Albums {
		if curAlbum.ID == paramsID {
			//time complexity can be improved by using map
			data.Albums = append(data.Albums[:index], data.Albums[index+1:]...)
			newAlbum.ID = paramsID

			data.Albums = append(data.Albums, newAlbum)
			sort.SliceStable(data.Albums, func(i, j int) bool {
				return data.Albums[i].ID < data.Albums[j].ID
			})
			json.NewEncoder(w).Encode(data.Albums)

			return
		}
	}

	json.NewEncoder(w).Encode("No album data available for the given id")

}
func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r) // Get the params

	param := chi.URLParam(r, "id")
	paramsID, _ := strconv.Atoi(param)
	//fmt.Println("------------------>paramid", paramsID)

	//Loop through data.Albums and find that id
	for index, curAlbum := range data.Albums {
		if curAlbum.ID == paramsID {
			//fmt.Println("----------------->hi")
			//time complexity can be improved by using map
			data.Albums = append(data.Albums[:index], data.Albums[index+1:]...)
			break
		}
	}
	fmt.Println(data.Albums)

	json.NewEncoder(w).Encode(data.Albums)
}

// GetArtists Get All Artist
func GetArtists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var uniqueArtist = make(map[data.Artist]bool)
	var uniqueArtistList []data.Artist
	for _, curAlbum := range data.Albums {
		curArtist := curAlbum.Artist
		_, keyExists := uniqueArtist[curArtist]
		if keyExists == false {
			uniqueArtist[curArtist] = true
			uniqueArtistList = append(uniqueArtistList, curArtist)
		}
	}
	json.NewEncoder(w).Encode(uniqueArtistList)
}

// GetTopArtists Get Top X Artist by album total number of album count
func GetTopArtists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := chi.URLParam(r, "cnt")
	//paramsID, _ := strconv.Atoi(param)
	//params := mux.Vars(r) // Get the params

	numberOfArtistToShow, _ := strconv.Atoi(param)

	var artistCount = make(map[data.Artist]int)

	for _, curAlbum := range data.Albums {
		artistCount[curAlbum.Artist]++
	}

	type ArtistsAlbumCount struct {
		Artist   data.Artist
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

	var creds data.Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)

	fmt.Println(creds)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	correctPassword, ok := data.Credslist[creds.Username]

	if !ok || creds.Password != correctPassword {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//tokenAuth = jwtauth.New(string(jwa.HS256), jwtkey, nil)

	expiretime := time.Now().Add(10 * time.Minute)

	_, tokenString, err := data.TokenAuth.Encode(map[string]interface{}{
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

func StartServer(port int) {
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
		r.Use(jwtauth.Verifier(data.TokenAuth))
		r.Use(jwtauth.Authenticator)

		//	r.Use(middleware.BasicAuth("user", data.Credslist))

		r.Route("/albums", func(r chi.Router) {
			r.Get("/", GetAlbums)
			r.Get("/{id}", GetAlbum)
			r.Post("/", CreateAlbum)
			r.Delete("/{id}", DeleteAlbum)
			r.Put("/{id}", UpdateAlbum)
		})

		r.Route("/artists", func(r chi.Router) {
			r.Get("/", GetArtists)
			r.Get("/{cnt}", GetTopArtists)
		})
		r.Post("/logout", Logout)

	})
	Server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: r,
	}
	fmt.Println(Server.ListenAndServe())

	//http.ListenAndServe(":5050", r)

	//previous code using max
	//router := max.NewRouter()
	////Route Handlers /Endpoints
	//router.HandleFunc("/albums", GetAlbums).Methods("GET")
	//router.HandleFunc("/albums/{id}", GetAlbum).Methods("GET")
	//router.HandleFunc("/albums", createAlbum).Methods("POST")
	//router.HandleFunc("/albums/{id}", UpdateAlbum).Methods("PUT")
	//router.HandleFunc("/albums/{id}", DeleteAlbum).Methods("DELETE")
	//
	//router.HandleFunc("/artists", GetArtists).Methods("GET")
	//router.HandleFunc("/artists/{x}", getTopArtists).Methods("GET")

}
