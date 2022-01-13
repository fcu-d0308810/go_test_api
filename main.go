package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func AllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented !")
	fmt.Printf("%T\n", w)
}

func FindMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented !")
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented !")
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented !")
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented !")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", AllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", FindMovie).Methods("GET")
	r.HandleFunc("/movies", CreateMovie).Methods("POST")
	r.HandleFunc("/movies", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovie).Methods("DELETE")
	fmt.Println("go api start....")
	http.ListenAndServe(":5000", r)
}
