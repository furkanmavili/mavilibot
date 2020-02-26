package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/bwmarrin/discordgo"
	"github.com/furkanmavili/discord/handler"
)

// Variables used for command line parameters

func main() {

	go herokuListen()

	value, exists := os.LookupEnv("DISCORD_TOKEN")
	if !exists {
		log.Println("Token doesn't exist.")
		return
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
