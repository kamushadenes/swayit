package common

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	olog "log"
	"os"
	"time"
)

var Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()

func init() {
	zerolog.TimeFieldFormat = time.RFC3339

	olog.SetFlags(0)
	olog.SetOutput(ioutil.Discard)
}
