package environmentvariablesconfig

import (
	"log"

	"github.com/joho/godotenv"
)

func SetupEnvironmentVariables() {
	error := godotenv.Load(".env")
	if error != nil {
		log.Fatalf("Ocorreu um erro ao carregar o log %v", error)
		panic(error)
	}
}
