package main

import (
	"encoding/json"
	"fmt"
	"hello/go/pkg/mod/github.com/gorilla/mux@v1.8.0"
	"log"
	"math/random"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
  id string `json: "id"`
  isbn string `json: "isbn"`
  title string `json: "title"`
  director *Director `json: "director`
}

type Director struct{
   firstname string `json: "firstname"`
   lastname string `json: "lastname`
}

var movies [] Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	
}

func main(){
  
	r:= mux.NewRouter()


	//for testing
	movies = append(movies, Movie{
		id: "1",
		isbn: "4228",
	    title: "firstmovie",
		director: &Director{
			firstname: "Chris",
			lastname: "Nolan",
		},
	})

	movies = append(movies, Movie{
		id: "2",
		isbn: "5347",
	    title: "secondmovie",
		director: &Director{
			firstname: "Chris",
			lastname: "Nolan",
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/(id)", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("movies/(id)", updateMovie).Methods("PUT")
	r.HandleFunc("movies/(id)", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}