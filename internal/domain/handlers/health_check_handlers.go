package web

import "net/http"

type HealthCheckHandlers interface {
	CheckAPIHealth(res http.ResponseWriter, req *http.Request)
}
