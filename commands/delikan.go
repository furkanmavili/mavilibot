package commands

import (
	"github.com/bwmarrin/discordgo"
)

//Delikan func
func Delikan(s *discordgo.Session, m *discordgo.MessageCreate) error {
	s.ChannelMessageSend(m.ChannelID, "https://www.youtube.com/watch?v=O-iFJkvHSH4")
	return nil
}
