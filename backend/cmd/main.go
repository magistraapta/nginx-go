package main

import (
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

	log.Println("Server is running on port", os.Getenv("PORT"))

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

type Message struct {
	Message string `json:"message"`
	Server  string `json:"server"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{
		Message: "Hello from docker",
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
