package model

type GyroRecord struct {
	ID    string  `json:"id"`
	Roll  float64 `json:"roll"`
	Pitch float64 `json:"pitch"`
	Yaw   float64 `json:"yaw"`
}
