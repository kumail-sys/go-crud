package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			fmt.Println("deleted Id:", item.Id)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMoviesbyId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application-json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
	//json.NewEncoder(w).Encode("created movies")

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			fmt.Println("deleted Id:", item.Id)
			break
		}
	}

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = params["id"]
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{Id: "203", Isbn: "3567", Title: "Thor", Director: &Director{Firstname: "Mark", Lastname: "manson"}})
	movies = append(movies, Movie{Id: "205", Isbn: "3567", Title: "Thor", Director: &Director{Firstname: "Mark", Lastname: "manson"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMoviesbyId).Methods("GET")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies", createMovie).Methods("POST")

	fmt.Println("server starting at the port 8000:")
	log.Fatal(http.ListenAndServe(":8000", r))
}
