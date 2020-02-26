package main

import (
	"log"
	"os"

	_ "github.com/bwmarrin/discordgo"
	"github.com/furkanmavili/discord/handler"
)

// Variables used for command line parameters


func main() {

	value, exists := os.LookupEnv("DISCORD_TOKEN")
	if !exists {
		log.Println("Token bulunamadı.")
		return
	}
	err := handler.Connect(value);
	if err != nil {
		panic(err)
	}
}
