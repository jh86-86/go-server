package main

import(
	//"encoding/json"
	"log"
	"net/http"
//	"math/rand"
	//"strconv"
	"github.com/gorilla/mux"

)

//Book strut is basically model , kind of like class, OOP

type Book struct {
	ID string `json: "id"`
	Isbn string `json: "isbn"`
	title string `json: "title"`
	Author *Author `json: "author"`
}

//Author struct

type Author struct {
	Firstname string `json: "json: firstname"`
	Lastname string `json: "json: lastname"`
	Nickname string `json: "json: nickname:`
}


//get all books

func getBooks(w http.ResponseWriter, r *http.Request) { //response and request types
	
}

//get single book
func getBook(w http.ResponseWriter, r *http.Request) { //response and request types

}


func createBook(w http.ResponseWriter, r *http.Request) { //response and request types

}

func updateBook(w http.ResponseWriter, r *http.Request) { //response and request types

}

func deleteBook(w http.ResponseWriter, r *http.Request) { //response and request types

}



func main(){
	//Init router
	r := mux.NewRouter();

	//route handlers / endpoints

	r.HandleFunc("/api/books", getBooks).Methods("GET")  //needs method to know what type of http method is needed
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")  //needs method to know what type of http method is needed
	r.HandleFunc("/api/books", createBook).Methods("Post")  //needs method to know what type of http method is needed
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")  //needs method to know what type of http method is needed
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE") 
	
	log.Fatal(http.ListenAndServe(":8000", r)) //r is the router and ListenAndServe is similar to nodes listen,log fatal gets an error if it fails
}



