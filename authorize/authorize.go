package authorize

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var authKey = []byte(os.Getenv("SESSION_KEY"))
var encryptKey = []byte(os.Getenv("SESSION_KEY"))

//var authKey = []byte("SESSION_KEY")

var store = sessions.NewCookieStore(authKey)

func InitSession(r *http.Request) *sessions.Session {
	session, err := store.Get(r, "authorized")
	if err != nil {
		log.Panic(err)
	}
	return session
}
