package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"workoutwidget.fit/sensehatrest/model"
	"workoutwidget.fit/sensehatrest/service"
)

type MotionController struct {
	MotionRepo service.MotionCrudRepository
}

// HandleMotionRequest - handle requests for the motion objects
func (mc *MotionController) HandleMotionRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		mc.handleGet(w, r)
		break
	case "POST":
		mc.handlePost(w, r)
		break
	case "PUT":
		http.Error(w, "PUT Not a supported method", http.StatusMethodNotAllowed)
		break
	case "DELETE":
		http.Error(w, "DELETE Not a supported method", http.StatusMethodNotAllowed)
		break
	}
}

func (mc *MotionController) handleGet(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query()) > 0 {
		mc.handleGetWithParams(w, r)
	} else {
		mc.handleGetWithoutParams(w, r)
	}

}

func (mc *MotionController) handleGetWithParams(w http.ResponseWriter, r *http.Request) {

	if id := r.URL.Query().Get("id"); id == "" {

		log.Println("Error: Record ID not present.")
		http.Error(w, "Error: Record ID not present.", http.StatusBadRequest)
		return

	} else {

		record, err := mc.MotionRepo.GetById(id)
		if err != nil {
			log.Printf("Cannot retrieve record with ID: %s, Error: %s\n", id, err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		output, err := json.MarshalIndent(record, "", "\t\t")
		w.WriteHeader(201)
		w.Header().Set("Content-Type", "appliation/json")
		_, _ = w.Write(output)
	}
}

func (mc *MotionController) handleGetWithoutParams(w http.ResponseWriter, r *http.Request) {

	http.Error(w, "NOT IMPLEMENTED!", http.StatusInternalServerError)

}

func (mc *MotionController) handlePost(w http.ResponseWriter, r *http.Request) {

	//	log.Println("Saving motion record")
	requestLength := r.ContentLength
	bodyBytes := make([]byte, requestLength)

	r.Body.Read(bodyBytes)

	// create struct reference to pass hold unmarshalled input
	var motionRecord model.MotionRecord
	if err := json.Unmarshal(bodyBytes, &motionRecord); err != nil {
		log.Printf("MotionController.handlePost(...) -> ERROR: Could not unmarshall json. Error: %s\n", err.Error())
		http.Error(w, "An unexpected error occurred.", http.StatusInternalServerError)
		return
	}

	// save the record
	result, err := mc.MotionRepo.Save(&motionRecord)
	if err != nil {
		log.Fatalln(err)
	}

	// get the ID of the last inserted record
	lastInsertId := result.InsertedID.(primitive.ObjectID)

	/* TODO fix this
	// get the last inserted record
	record, err := mc.MotionRepo.GetById(lastInsertId.String())
	if err != nil {
		log.Println(err)
		return
	}
	// marshal the struct for JSON response
	output, err := json.MarshalIndent(&record, "", "\t\t")
	if err != nil {
		log.Println(err)
		return
	}
	*/

	//log.Printf("Saved record with ID: %s\n", lastInsertId)
	value := fmt.Sprintf(`{"objectId": "%s"}`, lastInsertId)
	//output, err := json.MarshalIndent(value, "", "\t\t")
	//if err != nil {
	//	log.Printf("MotionController.handlePost(...) -> Could not marshal InsertedID:%s\n", err.Error())
	//	http.Error(w, "An unexpected error occurred.", http.StatusInternalServerError)
	//	return
	//}

	w.WriteHeader(201)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(value))

	return
}
