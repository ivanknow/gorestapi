package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
*/
//Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func getBooks(writer http.ResponseWriter, reader *http.Request) {

}

func getBookById(writer http.ResponseWriter, reader *http.Request) {

}

func updateBook(writer http.ResponseWriter, reader *http.Request) {

}

func deleteBook(writer http.ResponseWriter, reader *http.Request) {

}

func createBook(writer http.ResponseWriter, reader *http.Request) {

}

func main() {
	fmt.Printf("hello, world\n")
	//Create router
	router := mux.NewRouter()

	router.HandleFunc("api/books", getBooks).Methods("GET")
	router.HandleFunc("api/books", createBook).Methods("POST")
	router.HandleFunc("api/books{id}", deleteBook).Methods("DELETE")
	router.HandleFunc("api/books{id}", updateBook).Methods("PUT")
	router.HandleFunc("api/books{id}", getBookById).Methods("GET")

	log.Fatal(http.ListenAndServe(":8001", router))

}
