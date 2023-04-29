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
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-chi/chi/v5"
	"go.uber.org/dig"
)

func main() {
	container := dig.New()

	//envVariablesConfig.SetupEnvironmentVariables()

	dbConnection, error := sql.Open("sqlserver", os.Getenv("DB_CONNECTION"))
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

	error = container.Provide(diHandlers.ProvideHealthCheckHandlers)
	if error != nil {
		panic(error)
	}

	error = container.Provide(chi.NewRouter)
	if error != nil {
		panic(error)
	}

	error = container.Invoke(func(router *chi.Mux, contestHandler handlers.ContestsHandlers, healthCheckHandler handlers.HealthCheckHandlers) {
		routerConfiguration.SetContestsRoutes(router, contestHandler)
		routerConfiguration.SetHealthCheckRoutes(router, healthCheckHandler)

		fmt.Println("Running Azl-Vote-API on 80 Port")
		http.ListenAndServe(":80", router)

	})
	if error != nil {
		panic(error)
	}
}
