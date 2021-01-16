package main

import (
	"encoding/json"
	"log"
	"net/http"
	"workoutwidget.fit/sensehatrest/model"
)

func main() {

	server := http.Server{
		Addr: "127/0.0.1:8080",
	}

	http.HandleFunc("/record/", handleRequest)
	server.ListenAndServe()

}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {

	healthCheck := model.HealthCheck{
		Service: "Sense Hat REST API",
		Status: "Active",
	}

	output, err := json.MarshalIndent(&healthCheck, "", "\t\t")
	if err != nil {
		log.Printf("Main.handleGet(...) -> Could not marshal struct!")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

	return
}

func handlPost(w http.ResponseWriter, r *http.Request) (err error) {

	requestLength := r.ContentLength
	bodyBytes := make([]byte, requestLength)

	r.Body.Read(bodyBytes)

	var gyroRecord model.GyroRecord
	json.Unmarshal(bodyBytes, &gyroRecord)

	err = writeRecordToDb(gyroRecord)
	if err != nil {
		w.WriteHeader(500)
	}

	w.WriteHeader(201)

	return
}

func writeRecordToDb(record model.GyroRecord) (err error){

	// TODO write to DB

	return
}