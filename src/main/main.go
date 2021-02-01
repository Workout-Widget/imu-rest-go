package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
	"workoutwidget.fit/sensehatrest/controller"
	"workoutwidget.fit/sensehatrest/service"
)

func main() {

	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	log.Println("Sense HAT REST API loading...")
	log.Println("Establishing database connection")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/test"))
	if err != nil {
		log.Fatalf("Could not connect to database! %s\n", err.Error())
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	log.Println("Instantiating Motion Controller")
	motionController := controller.MotionController{
		MotionRepo: &service.MotionService{
			Client: client,
		},
	}

	log.Println("Instantiating the Info Controller")
	infoController := controller.InfoController{}

	log.Println("Assigning handler functions...")
	http.HandleFunc("/motion/", motionController.HandleMotionRequest)
	http.HandleFunc("/experiment/", motionController.HandleExperimentRequest)
	http.HandleFunc("/info/", infoController.HandleInfoRequests)

	log.Println("Starting server...")
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
