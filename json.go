package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}

	type errResponse struct {
		Error string `json:"error"` // we can specify a key for marshalling our message into JSON object like below
	}
	/*
	{
		"error": <our error msg>
	}
	*/

	respondWithJson(w, code, errResponse{msg})
}

func respondWithJson[T any](w http.ResponseWriter, code int, payload T) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v\n", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
