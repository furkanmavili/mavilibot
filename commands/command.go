package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Command struct for all commands
type Command struct {
	Name        string
	Description string
	Run         func(s *discordgo.Session, m *discordgo.MessageCreate)
}

// Register function
func Register() map[string]Command {

	quote := Command{
		Name:        "quote",
		Description: "prints some quotes",
		Run:         Quote,
	}
	answer := Command{
		Name:        "answer",
		Description: "gives answer",
		Run:         Ball,
	}
	dice := Command{
		Name:        "dice",
		Description: "roll dice",
		Run:         Dice,
	}
	help := Command{
		Name:        "help",
		Description: "bot help",
		Run:         Help,
	}
	invite := Command{
		Name:        "invite",
		Description: "creates guild invite",
		Run:         Invite,
	}
	komutlar := map[string]Command{
		"!quote":  quote,
		"!answer": answer,
		"!dice":   dice,
		"!help":   help,
		"!invite": invite,
	}
	return komutlar
}

func Route(komutAdı string, komutlar map[string]Command, s *discordgo.Session, m *discordgo.MessageCreate) error {
	val, exists := komutlar[komutAdı]
	if exists {
		val.Run(s, m)
		return nil
	}
	return fmt.Errorf("komut bulunamadı")
}
