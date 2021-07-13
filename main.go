package main

import (
	"github.com/game-backend/config"
	"github.com/game-backend/logger"
)

func main() {
	config.SetupEnvironment()
	logger.SetUpLogging()
}
