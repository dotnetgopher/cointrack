package main

import (
	"cointrack/internal/pkg/db"

	logger "cointrack/internal/pkg/logger"
)

func main() {
	loggerInstance := logger.New()
	db, err := db.New(loggerInstance)
	if err != nil {
		loggerInstance.Fatal().Err(err).Msg("database init failed")
	}

	
}
