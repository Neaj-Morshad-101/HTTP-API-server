package data

import (
	"github.com/go-chi/jwtauth"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

var jwtkey = []byte("Neaj's Secret Key, He will not share it")
var TokenAuth *jwtauth.JWTAuth
var tokenString string
var token jwt.Token

type Credentials struct {
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

// MaxAlbumID For new ID generation
var MaxAlbumID int

// Albums Init a Albums var as a slice of Album struct
var Albums []Album
var Credslist map[string]string

func InitCreds() {
	TokenAuth = jwtauth.New(string(jwa.HS256), jwtkey, nil)
	Credslist = make(map[string]string)
	creds := Credentials{
		"Neaj Morshad",
		"1234",
	}
	Credslist[creds.Username] = creds.Password
}

func InitMaxID() {
	MaxAlbumID = 0
}

func InitAlbumDB() {
	var newAlbum Album
	MaxAlbumID++
	newAlbum = Album{
		ID:    MaxAlbumID,
		Title: "Amar mon mojaiya re",
		Artist: Artist{
			FirstName: "MonMojaiya",
			LastName:  "artist",
		},
		Language: "Bengali",
	}
	Albums = append(Albums, newAlbum)

	MaxAlbumID++
	newAlbum = Album{
		ID:    MaxAlbumID,
		Title: "Ekdin Matir Vitore Hobe Ghor",
		Artist: Artist{
			FirstName: "EkdinMatir",
			LastName:  "artist",
		},
		Language: "Bengali",
	}
	Albums = append(Albums, newAlbum)

	MaxAlbumID++
	newAlbum = Album{
		ID:    MaxAlbumID,
		Title: "Arijit Singh er Ekti Gaaaan",
		Artist: Artist{
			FirstName: "Arijit",
			LastName:  "Singh",
		},
		Language: "Hindi",
	}
	Albums = append(Albums, newAlbum)

}
