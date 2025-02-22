package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{Text: "Hello from Go!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
