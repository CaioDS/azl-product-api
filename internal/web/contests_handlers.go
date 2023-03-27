package web

import (
	dtos "azl-vote-api/internal/domain/dtos/contests"
	usecase "azl-vote-api/internal/domain/usecase/contests"
	"encoding/json"
	"log"
	"net/http"
)

type ContestsHandlers struct {
	ListContestsUsecase  usecase.ListContestsUsecase
	CreateContestUsecase usecase.CreateContestUsecase
}

func NewContestsHandlers(listContestsUsecase usecase.ListContestsUsecase, createContestUsecase usecase.CreateContestUsecase) *ContestsHandlers {
	return &ContestsHandlers{
		ListContestsUsecase:  listContestsUsecase,
		CreateContestUsecase: createContestUsecase,
	}
}

func (contestHandlers *ContestsHandlers) ListContestHandler(res http.ResponseWriter, req *http.Request) {
	output, error := contestHandlers.ListContestsUsecase.Execute()

	if error != nil {
		res.WriteHeader(http.StatusInternalServerError)
		log.Println("error", error)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(output)
}

func (contestsHandlers *ContestsHandlers) CreateContestHandler(res http.ResponseWriter, req *http.Request) {
	var input dtos.CreateContestInputDto

	error := json.NewDecoder(req.Body).Decode(&input)
	if error != nil {
		res.WriteHeader(http.StatusBadRequest)
		log.Println("error", error)
		return
	}

	output, error := contestsHandlers.CreateContestUsecase.Execute(input)
	if error != nil {
		res.WriteHeader(http.StatusInternalServerError)
		log.Println("error", error)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(output)
}
