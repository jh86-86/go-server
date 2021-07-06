package main

import "net/http" //contains all things we need

func main(){
	http.HandleFunc("/hello-world", func (w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello world")) //casted into a byte array
	})
	http.ListenAndServe(":8080", nil) //nil not important for now
}