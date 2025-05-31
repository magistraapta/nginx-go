package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"nginx-go/config"
	"os"
)

func main() {
	config.LoadEnv()

	router := http.NewServeMux()

	router.HandleFunc("/", GetMessage)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	fmt.Println("Server is running on port", os.Getenv("PORT"))
}

type Message struct {
	Message string `json:"message"`
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	message := Message{
		Message: "Hello from docker",
	}

	json.NewEncoder(w).Encode(message)
}
