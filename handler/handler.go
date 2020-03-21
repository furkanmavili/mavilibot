package handler

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/furkanmavili/mavili-bot/commands"
)

// Connect  is connection to discord with given token
func Connect(token string) error {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return err
	}
	dg.AddHandler(manager)
	channels, _ := dg.GuildChannels("675106841436356628")

	for _, v := range channels {
		fmt.Printf("Channel id: %s  Channel name: %s\n", v.ID, v.Name)
	}
	go func() {
		for range time.NewTicker(time.Minute).C {
			_, err := dg.ChannelMessageSend("675109890204762143", "dont forget washing ur hands!")
			if err != nil {
				log.Println("couldn't send ticker message", err)
			}
		}
	}()

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

	komutlar := Register()
	checkprefix := strings.HasPrefix(m.Content, "!")
	// ignore all messages bot itself and not starting with prefix
	if m.Author.ID == s.State.User.ID || !checkprefix {
		return
	}

	seperateMessage := strings.Split(m.Content, " ")
	var c string = strings.ToLower(seperateMessage[0])
	args := strings.Join(seperateMessage[1:], " ")
	fmt.Println(seperateMessage[0], args)

	err := route(c, komutlar, s, m)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Böyle bir komut yok!")
	}
}

func route(komutAdı string, komutlar map[string]commands.Command, s *discordgo.Session, m *discordgo.MessageCreate) error {
	val, exists := komutlar[komutAdı]
	if exists {
		val.Run(s, m)
		return nil
	}
	return fmt.Errorf("komut bulunamadı")
}

// Register function
func Register() map[string]commands.Command {

	quote := commands.Command{
		Name:        "quote",
		Description: "prints some quotes",
		Run:         commands.Quote,
	}
	answer := commands.Command{
		Name:        "answer",
		Description: "gives answer",
		Run:         commands.Ball,
	}
	dice := commands.Command{
		Name:        "dice",
		Description: "roll dice",
		Run:         commands.Dice,
	}
	help := commands.Command{
		Name:        "help",
		Description: "bot help",
		Run:         commands.Help,
	}
	invite := commands.Command{
		Name:        "invite",
		Description: "creates guild invite",
		Run:         commands.Invite,
	}
	komutlar := map[string]commands.Command{
		"!quote":  quote,
		"!answer": answer,
		"!dice":   dice,
		"!help":   help,
		"!invite": invite,
	}
	return komutlar
}
