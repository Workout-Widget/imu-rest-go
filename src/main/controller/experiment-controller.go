package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

func (mc *MotionController) HandleExperimentRequest(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		mc.getAllByExperimentId(w, r)
		break
	case "POST":
		http.Error(w, "POST not implemented!", http.StatusMethodNotAllowed)
	case "PUT":
		http.Error(w, "PUT not implemented!", http.StatusMethodNotAllowed)
		break
	case "DELETE":
		http.Error(w, "DELETE not implemented!", http.StatusMethodNotAllowed)
	}
}

func (mc *MotionController) getAllByExperimentId(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	log.Printf("ExperimentController.getAllByExperimentId(...) -> Getting all records for experiment %s\n", id)

	records, err := mc.MotionRepo.GetAllByExperimentId(id)
	if err != nil {
		log.Printf("MotionController.getAllByExperimentId(...) -> ERROR: %v", err)
	}
	log.Println("")
	output, err := json.MarshalIndent(records, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, _ = w.Write(output)
	return
}
