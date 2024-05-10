package utility

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var resourceManager = ResourceManager{}
var timeStampFormat, _ = resourceManager.GetProperty("timeStampFormat")
var logFile = openLogFile()

// Logrus logger configuration (can be used globally in this application)
var Log = &logrus.Logger{
	Out:          io.MultiWriter(logFile, os.Stdout),
	Formatter:    &logrus.TextFormatter{TimestampFormat: timeStampFormat},
	Hooks:        make(logrus.LevelHooks),
	Level:        logrus.TraceLevel,
	ReportCaller: true,
}

func openLogFile() *os.File {
	project := resourceManager.GetProjectLocation()
	logFileDir := filepath.Join(project, "log")

	if _, err := os.Stat(logFileDir); os.IsNotExist(err) {
		os.Mkdir(logFileDir, 0755)
	}

	logfilePath := filepath.Join(logFileDir, "errorlog.log")
	logFile, err := os.OpenFile(logfilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Could not configure logger")
	}
	return logFile
}
