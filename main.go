package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func aliveAndKicking(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "alive and kicking")
}

func listMusics(w http.ResponseWriter, req *http.Request) {
	dir, err := os.ReadDir(songDir)
	if err != nil {
		log.Fatal(err)
	}

	songs := []string{}
	for _, d := range dir {
		songs = append(songs, d.Name())
	}
	jsonB, err := json.Marshal(map[string][]string{"songs": songs})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, string(jsonB))
}

var musicHandler = http.FileServer(http.Dir(songDir))

const songDir = "songs"
const port = "1337"

func main() {
	http.HandleFunc("/is-alive", aliveAndKicking)
	http.HandleFunc("/musics", listMusics)

	http.Handle("/play/", http.StripPrefix("/play/", musicHandler))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}
