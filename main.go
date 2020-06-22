package main

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

var (
	addr = flag.String("addr", "localhost:8080", "address to host the server on")
)

var (
	strava_config StravaAPIKeys
	conf          *oauth2.Config
)

func main() {
	strava_config = loadSecrets()

	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/login", loginHandler)

	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		log.Fatal(err)
	}

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello!")
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "TODO")
}

type StravaAPIKeys struct {
	ClientSecret string `json:"client_secret"`
	ClientID     string `json:"client_id"`
}

func loadSecrets() StravaAPIKeys {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	jsonBlob, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var secrets StravaAPIKeys

	json.Unmarshal(jsonBlob, &secrets)

	return secrets
}
