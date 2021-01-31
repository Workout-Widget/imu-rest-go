package model

import "time"

type MotionRecord struct {
	ID             string    `json:"id"`
	ExperimentID   string    `json:"experimentId"`
	Subject        string    `json:"subject"`
	SensorLocation string    `json:"sensorLocation"`
	Type           string    `json:"type"`
	Raw            bool      `json:"raw"`
	XRoll          float64   `json:"xRoll"`
	YPitch         float64   `json:"yPitch"`
	ZYaw           float64   `json:"zYaw"`
	Timestamp      time.Time `json:"timestamp"`
}
