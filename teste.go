package main

import "net/http"

func main() {

	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Olá Mundo"))
}