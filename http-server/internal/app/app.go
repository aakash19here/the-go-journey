package app

import (
	"log"
	"os"
)

const PORT int = 8080

type Application struct {
	Logger *log.Logger
	Port   int
}

// app constructor
func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	return &Application{
		Logger: logger,
		Port:   PORT,
	}, nil
}
