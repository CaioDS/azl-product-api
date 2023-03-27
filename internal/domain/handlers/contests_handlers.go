package web

import (
	"net/http"
)

type ContestsHandlers interface {
	ListContestHandler(res http.ResponseWriter, req *http.Request)
	CreateContestHandler(res http.ResponseWriter, req *http.Request)
}
