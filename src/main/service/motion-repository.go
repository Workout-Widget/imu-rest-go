package service

import (
	"go.mongodb.org/mongo-driver/mongo"
	"workoutwidget.fit/sensehatrest/model"
)

type MotionCrudRepository interface {
	GetAllByExperimentId(id string) ([]model.MotionRecord, error)
	GetById(id string) (model.MotionRecord, error)
	Save(record *model.MotionRecord) (*mongo.InsertOneResult, error)
	DeleteById(id string) (error)
}