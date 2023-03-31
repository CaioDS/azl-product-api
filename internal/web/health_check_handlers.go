package web

import (
	"net/http"
)

type HealthCheckHandlers struct{}

func NewHealthCheckHandlers() *HealthCheckHandlers {
	return &HealthCheckHandlers{}
}

func (healthCheckHandlers *HealthCheckHandlers) CheckAPIHealth(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
}
