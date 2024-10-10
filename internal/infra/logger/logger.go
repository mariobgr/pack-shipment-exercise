package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

type Logger struct {
	*log.Logger
}

// NewLogger creates a new Logger that writes in logs/default.log
func NewLogger() *Logger {
	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		panic("unable to create directory for logger")
	}

	logFile, err := os.OpenFile("logs/default.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(logFile)

	return &Logger{
		log.Default(),
	}
}

func (l *Logger) Error(v ...any) {
	l.Logger.SetPrefix(fmt.Sprintf("[ERROR] %s ", getCaller()))
	l.Logger.Println(v...)
}

func (l *Logger) Info(v ...any) {
	l.Logger.SetPrefix(fmt.Sprintf("[INFO] %s ", getCaller()))
	l.Logger.Println(v...)
}

func (l *Logger) Debug(v ...any) {
	l.Logger.SetPrefix(fmt.Sprintf("[DEBUG] %s ", getCaller()))
	l.Logger.Println(v...)
}

func (l *Logger) Warn(v ...any) {
	l.Logger.SetPrefix(fmt.Sprintf("[WARN] %s ", getCaller()))
	l.Logger.Println(v...)
}

// getCaller returns the file and line where the log method was called
func getCaller() string {
	_, filename, line, _ := runtime.Caller(2)
	dir, _ := os.Getwd()
	return fmt.Sprintf("%s:%d", strings.Replace(filename, dir+"/", "", 1), line)
}
