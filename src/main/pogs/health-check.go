package pogs

type HealthCheck struct {
	Service string `json:"service"`
	Status  string `json:"status"`
}
