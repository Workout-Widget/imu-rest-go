package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MotionRecord struct {
	ID             primitive.ObjectID `bson:"_id"`
	ExperimentID   string             `bson:"experimentId"`
	Subject        string             `bson:"subject"`
	SensorLocation string             `bson:"sensorLocation"`
	Type           string             `bson:"type"`
	Raw            bool               `bson:"raw"`
	XRoll          float64            `bson:"xRoll"`
	YPitch         float64            `bson:"yPitch"`
	ZYaw           float64            `bson:"zYaw"`
	Timestamp      time.Time          `bson:"timestamp"`
}
