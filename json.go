package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWIthError(w http.ResponseWriter, code int, msg string) {
	if code>499 {
		log.Println("Responding with 5XX error ", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}){
	data, err := json.Marshal(payload)

	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to marshall json response %v", payload)
		return
	}

	w.Header().Add("Content-Type", "Application/json")
	w.WriteHeader(code)
	w.Write(data)
}