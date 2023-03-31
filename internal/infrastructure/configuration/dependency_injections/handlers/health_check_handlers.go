package dependencyinjections

import (
	handler "azl-vote-api/internal/domain/handlers"
	handlerImpl "azl-vote-api/internal/web"
)

func ProvideHealthCheckHandlers() handler.HealthCheckHandlers {
	return handlerImpl.NewHealthCheckHandlers()
}
