package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	MovieID      string    `json:"movie_id"`
	MovieName    string    `json:"movie_name"`
	MovieRatings string    `json:"movie_ratings"`
	MovieIsbn    string    `json:"movie_isbn"`
	Director     *Director `json:"director"`
}
type Director struct {
	DirectorName string `json:"director_name"`
	DirectorAge  string `json:"director_age"`
}

var movies []Movie

func main() {

	movies = append(movies, Movie{
		MovieID:      "101",
		MovieName:    "See",
		MovieRatings: "9.8",
		MovieIsbn:    "1343",
		Director: &Director{
			"Peter Blay",
			"45",
		},
	})
	movies = append(movies, Movie{
		MovieID:      "102",
		MovieName:    "Ghajini",
		MovieRatings: "9.8",
		MovieIsbn:    "1343",
		Director: &Director{
			"Stigger Blay",
			"45",
		},
	})
	movies = append(movies, Movie{
		MovieID:      "103",
		MovieName:    "Bahubali",
		MovieRatings: "9.8",
		MovieIsbn:    "1343",
		Director: &Director{
			"Aboagye",
			"45",
		},
	})

	r := mux.NewRouter()
	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/get-movie/{id}", getSingleMovie).Methods("GET")
	r.HandleFunc("/update-movie/{id}", updateSingleMovie).Methods("PUT")
	r.HandleFunc("/delete-movie/{id}", deleteSingleMovie).Methods("DELETE")
	r.HandleFunc("/add-movie", addMovie).Methods("POST")

	fmt.Printf("Firing up the server at 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		return
	}
}

func deleteSingleMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.MovieID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getSingleMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.MovieID == params["id"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				return
			}
			return
		}
	}
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.MovieID = strconv.Itoa(rand.Intn(1000000000000))
	movies = append(movies, movie)
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		return
	}
}

func updateSingleMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var movie Movie
	for index, item := range movies {
		if item.MovieID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			err := json.NewDecoder(r.Body).Decode(&movie)
			if err != nil {
				return
			}
			movies = append(movies, movie)
			err = json.NewEncoder(w).Encode(movies)
			movie.MovieID = strconv.Itoa(rand.Intn(1000000000000))
			if err != nil {
				return
			}
		}
	}

}
