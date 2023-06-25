package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("hello handler", r.Header)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Message: "Hello World"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorldHandler)
	log.Println("Server is listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
