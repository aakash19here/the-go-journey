package app

import (
	"log"
	"os"
)

const PORT = ":8080"

type Application struct {
	Logger *log.Logger
	Port   string
}

// app constructor
func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	return &Application{
		Logger: logger,
		Port:   PORT,
	}, nil
}
