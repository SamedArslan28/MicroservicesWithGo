package main

import (
	"log"
	"log-service/data"
	"net/http"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(writer http.ResponseWriter, request *http.Request) {
	var payload JSONPayload
	err := app.readJSON(writer, request, &payload)
	if err != nil {
		return
	}
	entry := data.LogEntry{
		Name: payload.Name,
		Data: payload.Data,
	}
	err = app.Models.LogEntry.Insert(entry)
	if err != nil {
		log.Println("Error writing log entry:", err)
		return
	}

	response := jsonResponse{
		Message: "logged",
		Data:    nil,
	}
	app.writeJSON(writer, http.StatusAccepted, response)

}
