package structlog


import (
	"io"
	"log"
	"os"
)

var EventLogger *log.Logger = nil

func GetLogger() *log.Logger {
	if (EventLogger == nil) {
		panic("Event logger is nil.")
	}
	return EventLogger
}

func InitLogging(logFileName string) *log.Logger {


	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", logFileName, ":", err)
	}

	multi := io.MultiWriter(file, os.Stdout)
	MyLog := log.New(multi, "", log.LUTC)

	// remove default date/time flags from log output
	MyLog.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	EventLogger = MyLog

	return MyLog;
}

