package dependencyinjections

import (
	handler "azl-vote-api/internal/domain/handlers"
	usecase "azl-vote-api/internal/domain/usecase/contests"
	handlerImpl "azl-vote-api/internal/web"
)

func ProvideContestsHandlers(listContestsUsecase usecase.ListContestsUsecase, createContestUsecase usecase.CreateContestUsecase) handler.ContestsHandlers {
	return handlerImpl.NewContestsHandlers(listContestsUsecase, createContestUsecase)
}
