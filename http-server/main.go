package main

import (
	"net/http"
	"time"

	"github.com/aakash19here/the-go-journey/internal/app"
)

func main() {
	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	app.Logger.Println("We are starting the App")

	server := &http.Server{
		Addr:         app.Port,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()

	if err != nil {
		app.Logger.Fatal(err)
	}
}
