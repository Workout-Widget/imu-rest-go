package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"workoutwidget.fit/sensehatrest/model"
)

const (
	DB_NAME string = "test"
	MOTION_RECORD_COLLECTION = "motion_record"
)

type MotionService struct {
	Client *mongo.Client
}

func (ms *MotionService) GetAllByExperimentId(id string) (records []model.MotionRecord, err error) {
	log.Printf("MotionService.GetAllByExperimentId(...) -> Getting all records for the experiment: %s\n", id)
	collection := ms.Client.Database(DB_NAME).Collection(MOTION_RECORD_COLLECTION)

	filter := bson.D{
		{ "experimentId", id},
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Printf("MotionService.GetAllByExperimentId(...) -> %s\n", err.Error())
		return
	}
	defer cursor.Close(context.TODO())

	records = make([]model.MotionRecord, cursor.RemainingBatchLength())
	for cursor.Next(context.TODO()) {
		record := model.MotionRecord{}
		err := cursor.Decode(&record)
		if err != nil {
			log.Printf("MotionService.GetAllByExperimentId(...) -> Failed to decode MotionRecord")
		} else {
			records = append(records, record)
		}
	}

	return
}

func (ms *MotionService) GetById(id string) (record model.MotionRecord, err error){
	log.Printf("Getting record with the ID: %s\n", id)
	collection := ms.Client.Database(DB_NAME).Collection(MOTION_RECORD_COLLECTION)
	log.Println("Getting ObjectID")
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Printf("MotionService.GetById(...) -> Error creating ObjectID: %s", err.Error())
		return
	}

	filter := bson.D{
		{"_id", objId},
	}

	log.Println("Finding record with ID: ", id)

	err = collection.FindOne(context.TODO(), filter).Decode(&record)

	return
}

func (ms *MotionService) Save(record *model.MotionRecord) (result *mongo.InsertOneResult, err error) {

	collection := ms.Client.Database(DB_NAME).Collection(MOTION_RECORD_COLLECTION)

	filter := bson.D{
		{"experimentId",record.ExperimentID},
		{"subject",record.Subject},
		{"sensorLocation",record.SensorLocation},
		{"type",record.Type},
		{"raw",record.Raw},
		{"xRoll",record.XRoll},
		{"yPitch",record.YPitch},
		{"zYaw",record.ZYaw},
		{"timestamp",record.Timestamp},
	}

	result, err = collection.InsertOne(context.TODO(), filter)
	if err != nil {
		log.Printf("Error during save: %s\n", err.Error())
		return
	}

	log.Printf("x: %f y: %f z: %f\n",
		record.XRoll,
		record.YPitch,
		record.ZYaw,
	)

	return
}

func (ms *MotionService) DeleteById(id string) (err error) {

	collection := ms.Client.Database(DB_NAME).Collection(MOTION_RECORD_COLLECTION)
	filter := bson.D{
		{"_id", id},
	}
	_, err = collection.DeleteOne(context.TODO(), filter)
	return

}