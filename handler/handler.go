package handler

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/furkanmavili/discord/commands"
)

// Connect  is connection to discord with given token
func Connect(token string) error {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return err
	}
	dg.AddHandler(manager)
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return err
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	// Cleanly close down the Discord session.
	dg.Close()
	return nil
}

func manager(s *discordgo.Session, m *discordgo.MessageCreate) {

	checkprefix := strings.HasPrefix(m.Content, "!")
	// ignore all messages bot itself and not starting with prefix
	if m.Author.ID == s.State.User.ID || !checkprefix {
		return
	}

	seperateMessage := strings.Split(m.Content, " ")
	var c string = seperateMessage[0]
	c = strings.ToLower(c)
	args := strings.Join(seperateMessage[1:], " ")
	fmt.Println(seperateMessage[0], args)

	result := checkCommand(c)
	if !result {
		return
	}

	switch c {
	case "!help":
		commands.Help(s, m)
	// case "!delete":
	// 	messageDelete(s, m)
	case "!answer":
		commands.Ball(s, m)
	case "!invite":
		err := commands.Inviter(s, m)
		if err != nil {
			panic(err)
		}
	case "!quote":
		err := commands.Quote(s, m)
		if err != nil {
			panic(err)
		}
	case "!dice":
		err := commands.Dice(s, m)
		if err != nil {
			panic(err)
		}
	}
}

func checkCommand(c string) bool {
	var commandList = []string{"!help", "!answer", "!invite", "!quote", "!dice"}
	for _, v := range commandList {
		if v == c {
			return true
		}
	}
	return false
}
