package main

import (
	//"encoding/json"
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"  
	//	"math/rand"
	//"strconv"
	"github.com/gorilla/mux"
)

//Book strut is basically model , kind of like class, OOP

type Book struct {
	ID string `json: "id"`
	Isbn string `json: "isbn"`
	Title string `json: "title"`
	Author *Author `json: "author"`
}

//Author struct

type Author struct {
	Firstname string `json: "json: firstname"`
	Lastname string `json: "json: lastname"`
}

//init books variable as a slice book struct (slice is variable length array)

var books []Book


//get all books

func getBooks(w http.ResponseWriter, r *http.Request) { //response and request types
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//get single book
func getBook(w http.ResponseWriter, r *http.Request) { //response and request types
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //will get anny params
	//loop through books and find with id
	for _, item := range books {  //range used to loop through other data structues
		if item .ID == params["id"] {
			json.NewEncoder(w).Encode(item) //outputs the item of single book
			return 
		}
	}

	json.NewEncoder(w).Encode(&Book{}) 
}


func createBook(w http.ResponseWriter, r *http.Request) { //response and request types
	w.Header().Set("Content-type", "application/json")
	var book Book  //set to book struct
	_= json.NewDecoder(r.Body).Decode(&book)
	book.ID= strconv.Itoa(rand.Intn(1000000))//rand get a random int between 1-100000 ...not safe for production
	books = append(books, book) //appends book to our books array
	json.NewEncoder(w).Encode(book) 


}

func updateBook(w http.ResponseWriter, r *http.Request) { //response and request types
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
		books = append(books[:index], books[index+1:]...) //like a slice in js, gets rid of book
	}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) { //response and request types
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
		books = append(books[:index], books[index+1:]...) //like a slice in js, gets rid of book
		var book Book
		_= json.NewDecoder(r.Body).Decode(&book)
		book.ID = params["id"]
		books = append(books, book)
		json.NewEncoder(w).Encode(book)
		return
	}
	}
	json.NewEncoder(w).Encode(books)
}



func main(){
	//Init router
	r := mux.NewRouter();

	//mock data

	books =append(books, Book{ID: "1", Isbn:"3435345", Title: "Book one", Author : &Author{Firstname: "John", Lastname: "Do"}});
	books =append(books, Book{ID: "2", Isbn:"3434345345", Title: "Book two", Author : &Author{Firstname: "Steve", Lastname: "Do"}});

	//route handlers / endpoints

	r.HandleFunc("/api/books", getBooks).Methods("GET")  //needs method to know what type of http method is needed
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")  //needs method to know what type of http method is needed
	r.HandleFunc("/api/books", createBook).Methods("POST")  //needs method to know what type of http method is needed
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")  //needs method to know what type of http method is needed
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE") 
	
	log.Fatal(http.ListenAndServe(":8000", r)) //r is the router and ListenAndServe is similar to nodes listen,log fatal gets an error if it fails
}



