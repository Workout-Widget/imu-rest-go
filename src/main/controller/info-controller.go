package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"workoutwidget.fit/sensehatrest/pogs"
)

type InfoController struct { }

func (ic *InfoController) HandleInfoRequests(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetInfo(w)
	case "POST":
		log.Println("POST Not implemented.")
	case "PUT":
		log.Println("PUT not implemented")
	case "DELETE":
		log.Println("DELETE not implemented")
	}
}

func handleGetInfo(w http.ResponseWriter) {

	healthCheck := pogs.HealthCheck{
		Service: "Sense Hat REST API",
		Status: "Active",
	}

	output, err := json.MarshalIndent(&healthCheck, "", "\t\t")
	if err != nil {
		log.Printf("InfoController.handleGet(...) -> Could not marshal struct!")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)

	return
}
