package main

import (
	"encoding/json"
	"fmt"
	// "hello/go/pkg/mod/github.com/gorilla/mux@v1.8.0"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"Id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}


var movies [] Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
	for index, item:= range movies{
		if item.Id==params["Id"] {
			movies=append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:=mux.Vars(r)
	for _, item:=range movies{
		if item.Id==params["Id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies,movie)

	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params:=mux.Vars(r)

	for index, item:= range movies{
        if item.Id==params["Id"]{
			movies=append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id=params["Id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main(){
  
	r:= mux.NewRouter()


	//for testing
	movies = append(movies, Movie{
		Id: "1",
		Isbn: "4228",
	    Title: "firstmovie",
		Director: &Director{
			Firstname: "Chris",
			Lastname: "Nolan",
		},
	})

	movies = append(movies, Movie{
		Id: "2",
		Isbn: "5347",
	    Title: "secondmovie",
		Director: &Director{
			Firstname: "Chris",
			Lastname: "Nolan",
		},
	})

	//api definitions

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{Id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{Id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{Id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}