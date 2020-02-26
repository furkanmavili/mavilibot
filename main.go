package main

import (
	"net/http"
	"os"

	"github.com/furkanmavili/discord/handler"
)

func main() {

	go herokuListen()

	value, exists := os.LookupEnv("DISCORD_TOKEN")
	if !exists {
		panic("DISCORD_TOKEN doesn't exist.")
	}
	err := handler.Connect(value)
	if err != nil {
		panic(err)
	}
}

func herokuListen() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT not set")
	}
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
