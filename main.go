package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/gorilla/mux"
	"github.com/pandagrrl/rxb-project/internal/models"
)

func main() {
	err := models.InitDB("postgres://postgres:postgres@localhost:5555/dvdrental?sslmode=disable")
	if err != nil {
		log.Fatal(errors.Wrap(err, "unable to connect to database"))
	}

	r := mux.NewRouter()
	r.HandleFunc("/categories", GetCategories).Methods("GET")
	r.HandleFunc("/films", SearchFilms).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func SearchFilms(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	title := v.Get("title")
	films, err := models.SearchFilms(title)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	filmsJson, err := json.Marshal(films)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	fmt.Fprintf(w, "%s", filmsJson)
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := models.AllFilmCategories()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	categoriesJson, err := json.Marshal(categories)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	fmt.Fprintf(w, "%s", categoriesJson)
}
