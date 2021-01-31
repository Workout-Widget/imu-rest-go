package service

import "go.mongodb.org/mongo-driver/mongo"

type ExperimentRepository struct {
	Client *mongo.Client
}
