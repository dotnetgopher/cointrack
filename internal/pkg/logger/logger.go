package log

import (
	"os"

	"github.com/rs/zerolog"
)

func New() *zerolog.Logger {
	l := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &l
}
