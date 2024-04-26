package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Release struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Type string `json:"type"`
	Imdb_id string `json:"imdb_id"`
	Tmdb_id int `json:"tmdb_id"`
	Tmdb_type string `json:"tmdb_type"`
	Season_number int `json:"season_number"`
	Poster_url string `json:"poster_url"`
	Source_release_date string `json:"source_release_date"`
	Source_id int `json:"source_id"`
	Source_name string `json:"source_name"`
	Is_original int `json:"is_original"`
}

type Releases  struct {
	Releases []Release `json:"releases"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func getReleases() {
	// query api
	response, err := http.Get("https://api.watchmode.com/v1/releases/?limit=20&apiKey=rzgc7dNFABuAZVsgILSGYQ8y7ahvvYHZjoraxf6O")

	// handle error
	if err != nil {
		log.Fatal(err)
	}

	// parse json
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))
	var responseObject Releases
	json.Unmarshal(responseData, &responseObject)

}	

func main() {
	getReleases()
	handleRequests();
}

