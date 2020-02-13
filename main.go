package main

import (
	"flag"

	"github.com/furkanmavili/discord/handler"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	err := handler.Connect(Token)
	if err != nil {
		panic(err)
	}
}
