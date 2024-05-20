package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", view)

	log.Println("server start listening on port 8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func view(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetView(w, r)
	case http.MethodPost:
		PostView(w, r)
	default:
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}

func GetView(w http.ResponseWriter, r *http.Request) {

}

func PostView(w http.ResponseWriter, r *http.Request) {

}
