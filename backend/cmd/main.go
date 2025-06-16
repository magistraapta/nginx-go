package main

import (
	"encoding/json"
	"log"
	"net/http"
	"nginx-go/config"
	"os"
	"text/template"
)

func main() {
	config.LoadEnv()

	router := http.NewServeMux()

	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/load", GetMessage)

	log.Println("Server is running on port", os.Getenv("PORT"))

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

type Message struct {
	Message string `json:"message"`
	Server  string `json:"server"`
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	message := Message{
		Message: "Hello from docker",
		Server:  os.Getenv("PORT"),
	}

	json.NewEncoder(w).Encode(message)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{
		Message: "Hello world",
		Server:  os.Getenv("PORT"),
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
