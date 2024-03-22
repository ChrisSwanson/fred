package fred

import (
	"io"
	"log"
	"os"

	"github.com/rs/zerolog"
)

func ConfigureLogger(logFile string, logLevel string, stdOut bool) (logger *zerolog.Logger) {

	logLvl, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		// just default to warn level
		logLvl = zerolog.WarnLevel
	}

	var logOutput []io.Writer

	if logFile != "" {
		// if the logFile has been passed, validate it exists
		_, err := os.Stat(logFile)
		if os.IsNotExist(err) {
			// make the file.
			file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
			if err != nil {
				log.Fatalf("error creating file: %v", err)
			}
			file.Close()
		}

		logFileHandle, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}

		logOutput = append(logOutput, logFileHandle)
	}

	if stdOut {
		logOutput = append(logOutput, os.Stdout)
	}

	// create the logger with any of the above (os.Stdout, and or the logfile)
	log := zerolog.New(io.MultiWriter(logOutput...)).With().Timestamp().Logger().Level(logLvl)
	return &log
}
