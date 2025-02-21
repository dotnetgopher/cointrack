package db

import (
	"time"

	"github.com/rs/zerolog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

func New(log *zerolog.Logger) (*gorm.DB, error) {
	Logger := logger.New(
		log,
		logger.Config{
			SlowThreshold:             300 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Error,           // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,                  // Disable color
		},
	)

	var err error = nil
	db, err = gorm.Open(sqlite.Open("cointrack.db"), &gorm.Config{Logger: Logger})

	if err != nil {
		return nil, err
	}

	return db, err
}
