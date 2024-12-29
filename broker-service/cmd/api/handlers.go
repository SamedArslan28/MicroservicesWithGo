package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Message string `json:"message"`
	Error   bool   `json:"error"`
	Data    any    `json:"data"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Message: "Broker API",
		Error:   false,
		Data:    nil,
	}
	out, _ := json.MarshalIndent(payload, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_, err := w.Write(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
