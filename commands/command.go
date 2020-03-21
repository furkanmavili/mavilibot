package commands

import "github.com/bwmarrin/discordgo"

// Command struct for all commands
type Command struct {
	Name        string
	Description string
	Run         func(s *discordgo.Session, m *discordgo.MessageCreate)
}
