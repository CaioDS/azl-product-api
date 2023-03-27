package main

import (
	handlers "azl-vote-api/internal/domain/handlers"
	repository "azl-vote-api/internal/domain/repository"
	diHandlers "azl-vote-api/internal/infrastructure/configuration/dependency_injections/handlers"
	diRepository "azl-vote-api/internal/infrastructure/configuration/dependency_injections/repository"
	diUsecase "azl-vote-api/internal/infrastructure/configuration/dependency_injections/usecase"
	routerConfiguration "azl-vote-api/internal/infrastructure/configuration/router"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-chi/chi/v5"
	"go.uber.org/dig"
)

func main() {
	container := dig.New()

	dbConnection, error := sql.Open("sqlserver", "server=x,1433;user id=sa;password=x;database=x;")
	if error != nil {
		panic(error)
	}
	defer dbConnection.Close()

	error = container.Provide(func() repository.ContestRepository {
		return diRepository.ProvideContestRepository(dbConnection)
	})
	if error != nil {
		panic(error)
	}

	error = container.Provide(diUsecase.ProvideListContestsUsecase)
	if error != nil {
		panic(error)
	}

	error = container.Provide(diUsecase.ProvideCreateContestUsecase)
	if error != nil {
		panic(error)
	}

	error = container.Provide(diHandlers.ProvideContestsHandlers)
	if error != nil {
		panic(error)
	}

	error = container.Provide(chi.NewRouter)
	if error != nil {
		panic(error)
	}

	error = container.Invoke(func(router *chi.Mux, contestHandler handlers.ContestsHandlers) {
		routerConfiguration.SetContestsRoutes(router, contestHandler)

		fmt.Println("Estamos ouvindo")
		http.ListenAndServe(":8000", router)

	})
	if error != nil {
		panic(error)
	}
}
